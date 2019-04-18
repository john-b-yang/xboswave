package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/john-b-yang/xboswave/golang/drivers/pelican/storage"
	"github.com/john-b-yang/xboswave/golang/drivers/pelican/types"
	pb "github.com/john-b-yang/xboswave/golang/drivers/protos"
	pb2 "github.com/john-b-yang/xboswave/proto"
	"google.golang.org/grpc"
)

type setpointsMsg struct {
	HeatingSetpoint *float64 `msgpack:"heating_setpoint"`
	CoolingSetpoint *float64 `msgpack:"cooling_setpoint"`
}

type stateMsg struct {
	HeatingSetpoint *float64 `msgpack:"heating_setpoint"`
	CoolingSetpoint *float64 `msgpack:"cooling_setpoint"`
	Override        *bool    `msgpack:"override"`
	Mode            *int     `msgpack:"mode"`
	Fan             *bool    `msgpack:"fan"`
}

type stageMsg struct {
	HeatingStages *int32 `msgpack:"enabled_heat_stages"`
	CoolingStages *int32 `msgpack:"enabled_cool_stages"`
}

type occupancyMsg struct {
	Occupancy bool  `msgpack:"occupancy"`
	Time      int64 `msgpack:"time"`
}

// Entity File (WAVE 3)
const EntityFile = "entity.ent"

// SiteRouter: Location of WAVEMQ site router to Publish and Subscribe to
const SiteRouter = "127.0.0.1:4516"

func main() {
	// Loading Settings Configuration File
	var confBytes []byte
	confBytes, confReadErr := ioutil.ReadFile("config.json")
	if confReadErr != nil {
		fmt.Printf("Failed to read config.json file properly: %s\n", confReadErr)
	}
	var confData map[string]interface{}
	if unmarshalErr := json.Unmarshal(confBytes, &confData); unmarshalErr != nil {
		fmt.Printf("Failed to unmarshal config.json file properly: %s\n", unmarshalErr)
	}
	username := confData["username"].(string)
	password := confData["password"].(string)
	sitename := confData["sitename"].(string)
	namespace := confData["namespace"].(string)
	baseURI := confData["base_uri"].(string)
	namespaceBytes, namespaceErr := base64.URLEncoding.DecodeString(namespace)
	if namespaceErr != nil {
		fmt.Printf("Failed to convert namespace to bytes: %v", namespaceErr)
		os.Exit(1)
	}

	// Loading Polling Interval Configuration File
	var paramBytes []byte
	paramBytes, paramReadErr := ioutil.ReadFile("params.json")
	if paramReadErr != nil {
		fmt.Printf("Failed to read params.json file properly: %s\n", paramReadErr)
		os.Exit(1)
	}
	var paramData map[string]interface{}
	if unmarshalErr := json.Unmarshal(paramBytes, &paramData); unmarshalErr != nil {
		fmt.Printf("Failed to unmarshal params.json file properly %s\n", unmarshalErr)
		os.Exit(1)
	}

	pollInt, pollIntErr := time.ParseDuration(paramData["poll_interval"].(string))
	pollDr, pollDrErr := time.ParseDuration(paramData["poll_interval_dr"].(string))
	pollSched, pollSchedErr := time.ParseDuration(paramData["poll_interval_sched"].(string))
	if pollIntErr != nil {
		fmt.Printf("Failed to parse duration of poll interval properly: %v", pollIntErr)
		os.Exit(1)
	}
	if pollDrErr != nil {
		fmt.Printf("Failed to parse duration of demond response (DR) poll interval properly: %v", pollDrErr)
		os.Exit(1)
	}
	if pollSchedErr != nil {
		fmt.Printf("Failed to parse duration of schedule poll interval properly: %v", pollDrErr)
		os.Exit(1)
	}

	// Load WAVE3 Entity File to be used
	perspective, perspectiveErr := ioutil.ReadFile(EntityFile)
	if perspectiveErr != nil {
		fmt.Printf("Could not load entity %v, you might need to create one and grant it permissions\n", EntityFile)
		os.Exit(1)
	}

	// Establish a GRPC connection to the site router
	conn, err := grpc.Dial(SiteRouter, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Could not connect to the site router: %v\n", err)
		os.Exit(1)
	}
	client := pb.NewWAVEMQClient(conn)

	// Go Channel for communicating errors
	errorChan := make(chan bool)

	// Retrieving Pelicans
	pelicans, err := storage.ReadPelicans(username, password, sitename)
	if err != nil {
		fmt.Printf("Failed to read thermostat info: %v\n", err)
		os.Exit(1)
	}

	for _, pelican := range pelicans {
		pelican := pelican
		name := strings.Replace(pelican.Name, " ", "_", -1)
		name = strings.Replace(name, "&", "_and_", -1)
		name = strings.Replace(name, "'", "", -1)
		fmt.Println("Transforming", pelican.Name, "=>", name)

		// Ensure thermostat is running with correct number of stages
		if err := pelican.ModifyStages(&types.PelicanStageParams{
			HeatingStages: &pelican.HeatingStages,
			CoolingStages: &pelican.CoolingStages,
		}); err != nil {
			fmt.Printf("Failed to configure heating/cooling stages for pelican %s: %s\n", pelican.Name, err)
			errorChan <- true
		}

		subscribeParams := &pb.SubscribeParams{
			Perspective: &pb.Perspective{
				EntitySecret: &pb.EntitySecret{
					DER: perspective,
				},
			},
			Uri:        baseURI + "/" + pelican.Name,
			Namespace:  namespaceBytes,
			Identifier: "setpoints",
			Expiry:     60, // TODO(john-b-yang): Set appropriate amount of time here (currently 1 minute)
		}

		setpointsStream, setpointsErr := client.Subscribe(context.Background(), subscribeParams)
		if setpointsErr != nil {
			fmt.Printf("Failed to subscribe to setpoints slot: %v\n", setpointsErr)
			errorChan <- true
		}

		subscribeParams.Identifier = "schedule"
		scheduleStream, scheduleErr := client.Subscribe(context.Background(), subscribeParams)
		if scheduleErr != nil {
			fmt.Printf("Failed to subscribe to schedule slot: %v\n", scheduleErr)
			errorChan <- true
		}

		go func() {
			for {
				msg, err := setpointsStream.Recv()
				if err != nil {
					fmt.Println("Received malformed PO on setpoints slot. Dropping, ", err)
					errorChan <- true
				}
				if msg.Error != nil {
					fmt.Println("Received malformed PO on setpoints slot. Dropping, ", msg.Error.Message)
					errorChan <- true
				}

				content := []byte{}
				for _, po := range msg.Message.Tbs.Payload {
					content = append(content, po.Content...)
				}

				// TODO(john-b-yang): bw2 "ValueInto" with conversion method for byte slice -> struct (HINT: proto.Unmarshal)

				var setpoints setpointsMsg

				params := types.PelicanSetpointParams{
					HeatingSetpoint: setpoints.HeatingSetpoint,
					CoolingSetpoint: setpoints.CoolingSetpoint,
				}
				if err := pelican.ModifySetpoints(&params); err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Set heating setpoint to %v and cooling setpoint to %v\n",
						setpoints.HeatingSetpoint, setpoints.CoolingSetpoint)
				}
			}
		}()

		go func() {
			for {
				msg, err := scheduleStream.Recv()
				if err != nil {
					fmt.Println("Received malformed PO on schedule slot. Dropping, ", err)
					errorChan <- true
				}
				if msg.Error != nil {
					fmt.Println("Received malformed PO on schedule slot. Dropping, ", msg.Error.Message)
					errorChan <- true
				}

				content := []byte{}
				for _, po := range msg.Message.Tbs.Payload {
					content = append(content, po.Content...)
				}

				scheduleMsg := &pb2.ThermostatSchedule{}
				if unmarshalErr := proto.Unmarshal(content, scheduleMsg); unmarshalErr != nil {
					fmt.Println("Failed to unmarshal schedule message into correct format properly. Dropping, ", unmarshalErr)
				}

				var schedule types.ThermostatSchedule
				if schedule.DaySchedules == nil {
					fmt.Println("Received message on stage slot with no content. Dropping.")
					return
				}

				if err := pelican.SetSchedule(&schedule); err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Set pelican schedule to: %v", schedule)
				}
			}
		}()

		go func() {
			for {
				if status, err := pelican.GetStatus(); err != nil {
					fmt.Println("Failed to retrieve Pelican status: ", err)
					errorChan <- true
				} else if status != nil {
					fmt.Printf("%s %+v\n", pelican.Name, status)
					status := &pb2.XBOSIoTDeviceState{
						Time: uint64(status.Time),
						Thermostat: &pb2.Thermostat{
							Temperature:      &pb2.Double{Value: status.Temperature},
							RelativeHumidity: &pb2.Double{Value: status.RelHumidity},
							Override:         &pb2.Bool{Value: status.Override},
							FanState:         &pb2.Bool{Value: status.Fan},
							Mode:             pb2.HVACMode(status.Mode),
							State:            pb2.HVACState(status.State),
						},
					}
					// TODO(john-b-yang): "pelicanSchedule" may not be correct schema
					// TODO(john-b-yang): context.Background() default may not be correct
					if publishErr := publishWrapper(status, perspective, namespaceBytes, client, "pelicanStatus"); publishErr != nil {
						fmt.Println(publishErr)
						errorChan <- true
					}
				}
				time.Sleep(pollInt)
			}
		}()

		go func() {
			for {
				if schedStatus, schedErr := pelican.GetSchedule(); schedErr != nil {
					fmt.Println("Failed to retrieve Pelican's Schedule: ", schedErr)
					errorChan <- true
				} else {
					fmt.Printf("%s Schedule: %+v\n", pelican.Name, schedStatus)

					// Convert Go Struct to Proto Message
					schedule := &pb2.ThermostatSchedule{}
					for day, blockSchedules := range schedStatus.DaySchedules {
						var blockList []*pb2.ThermostatScheduleBlock
						for _, block := range blockSchedules {
							blockMsg := &pb2.ThermostatScheduleBlock{
								HeatingSetpoint: &pb2.Double{Value: block.HeatSetting},
								CoolingSetpoint: &pb2.Double{Value: block.CoolSetting},
								Time:            block.Time,
							}
							blockList = append(blockList, blockMsg)
						}
						schedule.ScheduleMap[day] = &pb2.ThermostatScheduleDay{
							Blocks: blockList,
						}
					}
					// TODO(john-b-yang): "pelicanSchedule" may not be correct schema
					// TODO(john-b-yang): context.Background() default may not be correct
					if publishErr := publishWrapper(schedule, perspective, namespaceBytes, client, "pelicanSchedule"); publishErr != nil {
						fmt.Println(publishErr)
						errorChan <- true
					}
				}
				time.Sleep(pollSched)
			}
		}()

		// TODO(john-b-yang): No Corresponding Proto Message for DR, Occupancy
		go func() {
			for {
				if drStatus, drErr := pelican.TrackDREvent(); drErr != nil {
					fmt.Println("Failed to retrieve Pelican's DR status: ", drErr)
					errorChan <- true
				} else if drStatus != nil {
					fmt.Printf("%s DR Status: %+v\n", pelican.Name, drStatus)
					// TODO(john-b-yang): Implement DR Status Publishing
				}
				time.Sleep(pollDr)
			}
		}()

		occupancy, err := pelican.GetOccupancy()
		if err != nil {
			fmt.Println("Failed to retrieve initial occupancy reading: ", err)
			errorChan <- true
		}

		// Start occupancy tracking loop only if thermostat has the necessary sensor
		if occupancy != types.OCCUPANCY_UNKNOWN {
			go func() {
				for {
					occupancy, err := pelican.GetOccupancy()
					if err != nil {
						fmt.Println("Failed to read thermostat occupancy: ", err)
						errorChan <- true
					} else {
						occupancyMsg := occupancyMsg{
							Occupancy: (occupancy == types.OCCUPANCY_OCCUPIED),
							Time:      time.Now().UnixNano(),
						}
						fmt.Printf("%s Occupancy Status: %+v\n", pelican.Name, occupancyMsg)
						// TODO(john-b-yang): Implement Occupancy Msg Publishing
					}
					time.Sleep(pollInt)
				}
			}()
		}
	}
	<-errorChan
}

func publishWrapper(message proto.Message, perspective []byte, namespaceBytes []byte, client pb.WAVEMQClient, schema string) error {
	scheduleBytes, scheduleErr := proto.Marshal(message)
	if scheduleErr != nil {
		return fmt.Errorf("Failed to serialized Pelican message: %v", scheduleErr)
	}
	payload := &pb.PayloadObject{
		Schema:  schema,
		Content: scheduleBytes,
	}
	publishParams := &pb.PublishParams{
		Perspective: &pb.Perspective{
			EntitySecret: &pb.EntitySecret{
				DER: perspective,
			},
		},
		Content:   []*pb.PayloadObject{payload},
		Namespace: namespaceBytes,
	}

	client.Publish(context.Background(), publishParams)
	return nil
}

/*
	for i, pelican := range pelicans {

		tstatIfaces[i].SubscribeSlot("state", func(msg *bw2.SimpleMessage) {
			po := msg.GetOnePODF(TSTAT_PO_DF)
			if po == nil {
				fmt.Println("Received message on state slot without required PO. Dropping.")
				return
			}

			var state stateMsg
			if err := po.(bw2.MsgPackPayloadObject).ValueInto(&state); err != nil {
				fmt.Println("Received malformed PO on state slot. Dropping.", err)
				return
			}

			params := types.PelicanStateParams{
				HeatingSetpoint: state.HeatingSetpoint,
				CoolingSetpoint: state.CoolingSetpoint,
			}
			fmt.Printf("%+v", state)
			if state.Mode != nil {
				m := float64(*state.Mode)
				params.Mode = &m
			}

			if state.Override != nil && *state.Override {
				f := float64(1)
				params.Override = &f
			} else {
				f := float64(0)
				params.Override = &f
			}

			if state.Fan != nil && *state.Fan {
				f := float64(1)
				params.Fan = &f
			} else {
				f := float64(0)
				params.Fan = &f
			}

			if err := pelican.ModifyState(&params); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Set Pelican state to: %+v\n", params)
			}
		})

		tstatIfaces[i].SubscribeSlot("stages", func(msg *bw2.SimpleMessage) {
			po := msg.GetOnePODF(TSTAT_PO_DF)
			if po == nil {
				fmt.Println("Received message on stage slot without required PO. Dropping.")
				return
			}

			var stages stageMsg
			if err := po.(bw2.MsgPackPayloadObject).ValueInto(&stages); err != nil {
				fmt.Println("Received malformed PO on stage slot. Dropping.", err)
				return
			}
			if stages.HeatingStages == nil && stages.CoolingStages == nil {
				fmt.Println("Received message on stage slot with no content. Dropping.")
				return
			}

			params := types.PelicanStageParams{
				HeatingStages: stages.HeatingStages,
				CoolingStages: stages.CoolingStages,
			}
			if err := pelican.ModifyStages(&params); err != nil {
				fmt.Println(err)
			} else {
				if stages.HeatingStages != nil {
					fmt.Printf("Set pelican heating stages to: %d\n", *stages.HeatingStages)
				}
				if stages.CoolingStages != nil {
					fmt.Printf("Set pelican cooling stages to: %d\n", *stages.CoolingStages)
				}
			}
		})
	}
*/
