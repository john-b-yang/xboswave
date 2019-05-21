// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iot.proto

//This is designed to be included by the main xbos proto file and includes the
//definitions for the XBOS IoT devices
//
//Maintainer: Gabe Fierro
//Version 1.0

package xbospb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FanMode int32

const (
	FanMode_FanAuto FanMode = 0
	FanMode_FanOn   FanMode = 1
	FanMode_FanOff  FanMode = 2
)

var FanMode_name = map[int32]string{
	0: "FanAuto",
	1: "FanOn",
	2: "FanOff",
}

var FanMode_value = map[string]int32{
	"FanAuto": 0,
	"FanOn":   1,
	"FanOff":  2,
}

func (x FanMode) String() string {
	return proto.EnumName(FanMode_name, int32(x))
}

func (FanMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{0}
}

type HVACMode int32

const (
	HVACMode_HVACModeOff      HVACMode = 0
	HVACMode_HVACModeHeatOnly HVACMode = 1
	HVACMode_HVACModeCoolOnly HVACMode = 2
	HVACMode_HVACModeAuto     HVACMode = 3
)

var HVACMode_name = map[int32]string{
	0: "HVACModeOff",
	1: "HVACModeHeatOnly",
	2: "HVACModeCoolOnly",
	3: "HVACModeAuto",
}

var HVACMode_value = map[string]int32{
	"HVACModeOff":      0,
	"HVACModeHeatOnly": 1,
	"HVACModeCoolOnly": 2,
	"HVACModeAuto":     3,
}

func (x HVACMode) String() string {
	return proto.EnumName(HVACMode_name, int32(x))
}

func (HVACMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{1}
}

type HVACState int32

const (
	HVACState_HVACStateOff        HVACState = 0
	HVACState_HVACStateHeatStage1 HVACState = 1
	HVACState_HVACStateCoolStage1 HVACState = 2
	HVACState_HVACStateHeatStage2 HVACState = 3
	HVACState_HVACStateCoolStage2 HVACState = 4
)

var HVACState_name = map[int32]string{
	0: "HVACStateOff",
	1: "HVACStateHeatStage1",
	2: "HVACStateCoolStage1",
	3: "HVACStateHeatStage2",
	4: "HVACStateCoolStage2",
}

var HVACState_value = map[string]int32{
	"HVACStateOff":        0,
	"HVACStateHeatStage1": 1,
	"HVACStateCoolStage1": 2,
	"HVACStateHeatStage2": 3,
	"HVACStateCoolStage2": 4,
}

func (x HVACState) String() string {
	return proto.EnumName(HVACState_name, int32(x))
}

func (HVACState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{2}
}

type URI struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *URI) Reset()         { *m = URI{} }
func (m *URI) String() string { return proto.CompactTextString(m) }
func (*URI) ProtoMessage()    {}
func (*URI) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{0}
}

func (m *URI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URI.Unmarshal(m, b)
}
func (m *URI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URI.Marshal(b, m, deterministic)
}
func (m *URI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URI.Merge(m, src)
}
func (m *URI) XXX_Size() int {
	return xxx_messageInfo_URI.Size(m)
}
func (m *URI) XXX_DiscardUnknown() {
	xxx_messageInfo_URI.DiscardUnknown(m)
}

var xxx_messageInfo_URI proto.InternalMessageInfo

func (m *URI) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *URI) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Triple struct {
	Subject              *URI     `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Predicate            *URI     `protobuf:"bytes,2,opt,name=predicate,proto3" json:"predicate,omitempty"`
	Object               *URI     `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Triple) Reset()         { *m = Triple{} }
func (m *Triple) String() string { return proto.CompactTextString(m) }
func (*Triple) ProtoMessage()    {}
func (*Triple) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{1}
}

func (m *Triple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Triple.Unmarshal(m, b)
}
func (m *Triple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Triple.Marshal(b, m, deterministic)
}
func (m *Triple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Triple.Merge(m, src)
}
func (m *Triple) XXX_Size() int {
	return xxx_messageInfo_Triple.Size(m)
}
func (m *Triple) XXX_DiscardUnknown() {
	xxx_messageInfo_Triple.DiscardUnknown(m)
}

var xxx_messageInfo_Triple proto.InternalMessageInfo

func (m *Triple) GetSubject() *URI {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *Triple) GetPredicate() *URI {
	if m != nil {
		return m.Predicate
	}
	return nil
}

func (m *Triple) GetObject() *URI {
	if m != nil {
		return m.Object
	}
	return nil
}

type XBOSIoTDeviceState struct {
	// current time at device/service
	//unit:ns
	Time uint64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	// unique identifier for this request; used to line up with device state requests
	Requestid int64 `protobuf:"varint,2,opt,name=requestid,proto3" json:"requestid,omitempty"`
	// any error that occured since the last device report. If requestid above is non-zero,
	// then this error corresponds to the request with the given requestid
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// XBOS IoT devices
	Thermostat               *Thermostat               `protobuf:"bytes,4,opt,name=thermostat,proto3" json:"thermostat,omitempty"`
	Meter                    *Meter                    `protobuf:"bytes,5,opt,name=meter,proto3" json:"meter,omitempty"`
	Light                    *Light                    `protobuf:"bytes,6,opt,name=light,proto3" json:"light,omitempty"`
	Evse                     *EVSE                     `protobuf:"bytes,7,opt,name=evse,proto3" json:"evse,omitempty"`
	WeatherStation           *WeatherStation           `protobuf:"bytes,8,opt,name=weather_station,json=weatherStation,proto3" json:"weather_station,omitempty"`
	WeatherStationPrediction *WeatherStationPrediction `protobuf:"bytes,9,opt,name=weather_station_prediction,json=weatherStationPrediction,proto3" json:"weather_station_prediction,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *XBOSIoTDeviceState) Reset()         { *m = XBOSIoTDeviceState{} }
func (m *XBOSIoTDeviceState) String() string { return proto.CompactTextString(m) }
func (*XBOSIoTDeviceState) ProtoMessage()    {}
func (*XBOSIoTDeviceState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{2}
}

func (m *XBOSIoTDeviceState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XBOSIoTDeviceState.Unmarshal(m, b)
}
func (m *XBOSIoTDeviceState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XBOSIoTDeviceState.Marshal(b, m, deterministic)
}
func (m *XBOSIoTDeviceState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XBOSIoTDeviceState.Merge(m, src)
}
func (m *XBOSIoTDeviceState) XXX_Size() int {
	return xxx_messageInfo_XBOSIoTDeviceState.Size(m)
}
func (m *XBOSIoTDeviceState) XXX_DiscardUnknown() {
	xxx_messageInfo_XBOSIoTDeviceState.DiscardUnknown(m)
}

var xxx_messageInfo_XBOSIoTDeviceState proto.InternalMessageInfo

func (m *XBOSIoTDeviceState) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *XBOSIoTDeviceState) GetRequestid() int64 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *XBOSIoTDeviceState) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *XBOSIoTDeviceState) GetThermostat() *Thermostat {
	if m != nil {
		return m.Thermostat
	}
	return nil
}

func (m *XBOSIoTDeviceState) GetMeter() *Meter {
	if m != nil {
		return m.Meter
	}
	return nil
}

func (m *XBOSIoTDeviceState) GetLight() *Light {
	if m != nil {
		return m.Light
	}
	return nil
}

func (m *XBOSIoTDeviceState) GetEvse() *EVSE {
	if m != nil {
		return m.Evse
	}
	return nil
}

func (m *XBOSIoTDeviceState) GetWeatherStation() *WeatherStation {
	if m != nil {
		return m.WeatherStation
	}
	return nil
}

func (m *XBOSIoTDeviceState) GetWeatherStationPrediction() *WeatherStationPrediction {
	if m != nil {
		return m.WeatherStationPrediction
	}
	return nil
}

type XBOSIoTDeviceActuation struct {
	// current time at device/service
	//unit:ns
	Time uint64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	// unique identifier for this request; used to line up with device state responses
	Requestid int64 `protobuf:"varint,2,opt,name=requestid,proto3" json:"requestid,omitempty"`
	// XBOS IoT devices
	Thermostat           *Thermostat `protobuf:"bytes,3,opt,name=thermostat,proto3" json:"thermostat,omitempty"`
	Meter                *Meter      `protobuf:"bytes,4,opt,name=meter,proto3" json:"meter,omitempty"`
	Light                *Light      `protobuf:"bytes,5,opt,name=light,proto3" json:"light,omitempty"`
	Evse                 *EVSE       `protobuf:"bytes,6,opt,name=evse,proto3" json:"evse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *XBOSIoTDeviceActuation) Reset()         { *m = XBOSIoTDeviceActuation{} }
func (m *XBOSIoTDeviceActuation) String() string { return proto.CompactTextString(m) }
func (*XBOSIoTDeviceActuation) ProtoMessage()    {}
func (*XBOSIoTDeviceActuation) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{3}
}

func (m *XBOSIoTDeviceActuation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XBOSIoTDeviceActuation.Unmarshal(m, b)
}
func (m *XBOSIoTDeviceActuation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XBOSIoTDeviceActuation.Marshal(b, m, deterministic)
}
func (m *XBOSIoTDeviceActuation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XBOSIoTDeviceActuation.Merge(m, src)
}
func (m *XBOSIoTDeviceActuation) XXX_Size() int {
	return xxx_messageInfo_XBOSIoTDeviceActuation.Size(m)
}
func (m *XBOSIoTDeviceActuation) XXX_DiscardUnknown() {
	xxx_messageInfo_XBOSIoTDeviceActuation.DiscardUnknown(m)
}

var xxx_messageInfo_XBOSIoTDeviceActuation proto.InternalMessageInfo

func (m *XBOSIoTDeviceActuation) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *XBOSIoTDeviceActuation) GetRequestid() int64 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *XBOSIoTDeviceActuation) GetThermostat() *Thermostat {
	if m != nil {
		return m.Thermostat
	}
	return nil
}

func (m *XBOSIoTDeviceActuation) GetMeter() *Meter {
	if m != nil {
		return m.Meter
	}
	return nil
}

func (m *XBOSIoTDeviceActuation) GetLight() *Light {
	if m != nil {
		return m.Light
	}
	return nil
}

func (m *XBOSIoTDeviceActuation) GetEvse() *EVSE {
	if m != nil {
		return m.Evse
	}
	return nil
}

type XBOSIoTContext struct {
	// current time at device/service
	//unit:ns
	Time uint64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	// any triples this device wants to add about itself
	// these triples will be assumed to be generated by the entity
	// who has created/signed this message
	Context              []*Triple `protobuf:"bytes,2,rep,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *XBOSIoTContext) Reset()         { *m = XBOSIoTContext{} }
func (m *XBOSIoTContext) String() string { return proto.CompactTextString(m) }
func (*XBOSIoTContext) ProtoMessage()    {}
func (*XBOSIoTContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{4}
}

func (m *XBOSIoTContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XBOSIoTContext.Unmarshal(m, b)
}
func (m *XBOSIoTContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XBOSIoTContext.Marshal(b, m, deterministic)
}
func (m *XBOSIoTContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XBOSIoTContext.Merge(m, src)
}
func (m *XBOSIoTContext) XXX_Size() int {
	return xxx_messageInfo_XBOSIoTContext.Size(m)
}
func (m *XBOSIoTContext) XXX_DiscardUnknown() {
	xxx_messageInfo_XBOSIoTContext.DiscardUnknown(m)
}

var xxx_messageInfo_XBOSIoTContext proto.InternalMessageInfo

func (m *XBOSIoTContext) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *XBOSIoTContext) GetContext() []*Triple {
	if m != nil {
		return m.Context
	}
	return nil
}

// Thermostat
type Thermostat struct {
	//Current temperature recorded by thermostat
	//unit:celsius
	Temperature *Double `protobuf:"bytes,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	//unit:% rh
	RelativeHumidity *Double `protobuf:"bytes,2,opt,name=relative_humidity,json=relativeHumidity,proto3" json:"relative_humidity,omitempty"`
	//If true, then the thermostat is in override mode
	//unit: t/f
	Override *Bool `protobuf:"bytes,3,opt,name=override,proto3" json:"override,omitempty"`
	//If true, the fan is on; else it is off
	//unit: t/f
	FanState *Bool `protobuf:"bytes,4,opt,name=fan_state,json=fanState,proto3" json:"fan_state,omitempty"`
	//Current operating mode of the fan
	//unit: xbos/iot/FanMode
	FanMode FanMode `protobuf:"varint,5,opt,name=fan_mode,json=fanMode,proto3,enum=xbospb.FanMode" json:"fan_mode,omitempty"`
	//Current operating mode of the HVAC
	//unit: xbos/iot/HVACMode
	Mode HVACMode `protobuf:"varint,6,opt,name=mode,proto3,enum=xbospb.HVACMode" json:"mode,omitempty"`
	//Current HVAC state
	//unit: xbos/iot/HVACState
	State HVACState `protobuf:"varint,7,opt,name=state,proto3,enum=xbospb.HVACState" json:"state,omitempty"`
	//number of heat stages enabled
	EnabledHeatStages *Int32 `protobuf:"bytes,8,opt,name=enabled_heat_stages,json=enabledHeatStages,proto3" json:"enabled_heat_stages,omitempty"`
	//number of cool stages enabled
	EnabledCoolStages *Int32 `protobuf:"bytes,9,opt,name=enabled_cool_stages,json=enabledCoolStages,proto3" json:"enabled_cool_stages,omitempty"`
	//heating setpoint
	//unit: celsius
	HeatingSetpoint *Double `protobuf:"bytes,10,opt,name=heating_setpoint,json=heatingSetpoint,proto3" json:"heating_setpoint,omitempty"`
	//cooling setpoint
	//unit: celsius
	CoolingSetpoint      *Double  `protobuf:"bytes,11,opt,name=cooling_setpoint,json=coolingSetpoint,proto3" json:"cooling_setpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Thermostat) Reset()         { *m = Thermostat{} }
func (m *Thermostat) String() string { return proto.CompactTextString(m) }
func (*Thermostat) ProtoMessage()    {}
func (*Thermostat) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{5}
}

func (m *Thermostat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Thermostat.Unmarshal(m, b)
}
func (m *Thermostat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Thermostat.Marshal(b, m, deterministic)
}
func (m *Thermostat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Thermostat.Merge(m, src)
}
func (m *Thermostat) XXX_Size() int {
	return xxx_messageInfo_Thermostat.Size(m)
}
func (m *Thermostat) XXX_DiscardUnknown() {
	xxx_messageInfo_Thermostat.DiscardUnknown(m)
}

var xxx_messageInfo_Thermostat proto.InternalMessageInfo

func (m *Thermostat) GetTemperature() *Double {
	if m != nil {
		return m.Temperature
	}
	return nil
}

func (m *Thermostat) GetRelativeHumidity() *Double {
	if m != nil {
		return m.RelativeHumidity
	}
	return nil
}

func (m *Thermostat) GetOverride() *Bool {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *Thermostat) GetFanState() *Bool {
	if m != nil {
		return m.FanState
	}
	return nil
}

func (m *Thermostat) GetFanMode() FanMode {
	if m != nil {
		return m.FanMode
	}
	return FanMode_FanAuto
}

func (m *Thermostat) GetMode() HVACMode {
	if m != nil {
		return m.Mode
	}
	return HVACMode_HVACModeOff
}

func (m *Thermostat) GetState() HVACState {
	if m != nil {
		return m.State
	}
	return HVACState_HVACStateOff
}

func (m *Thermostat) GetEnabledHeatStages() *Int32 {
	if m != nil {
		return m.EnabledHeatStages
	}
	return nil
}

func (m *Thermostat) GetEnabledCoolStages() *Int32 {
	if m != nil {
		return m.EnabledCoolStages
	}
	return nil
}

func (m *Thermostat) GetHeatingSetpoint() *Double {
	if m != nil {
		return m.HeatingSetpoint
	}
	return nil
}

func (m *Thermostat) GetCoolingSetpoint() *Double {
	if m != nil {
		return m.CoolingSetpoint
	}
	return nil
}

type ThermostatSchedule struct {
	//Map day of week to daily schedule
	ScheduleMap          map[string]*ThermostatScheduleDay `protobuf:"bytes,1,rep,name=scheduleMap,proto3" json:"scheduleMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ThermostatSchedule) Reset()         { *m = ThermostatSchedule{} }
func (m *ThermostatSchedule) String() string { return proto.CompactTextString(m) }
func (*ThermostatSchedule) ProtoMessage()    {}
func (*ThermostatSchedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{6}
}

func (m *ThermostatSchedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermostatSchedule.Unmarshal(m, b)
}
func (m *ThermostatSchedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermostatSchedule.Marshal(b, m, deterministic)
}
func (m *ThermostatSchedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermostatSchedule.Merge(m, src)
}
func (m *ThermostatSchedule) XXX_Size() int {
	return xxx_messageInfo_ThermostatSchedule.Size(m)
}
func (m *ThermostatSchedule) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermostatSchedule.DiscardUnknown(m)
}

var xxx_messageInfo_ThermostatSchedule proto.InternalMessageInfo

func (m *ThermostatSchedule) GetScheduleMap() map[string]*ThermostatScheduleDay {
	if m != nil {
		return m.ScheduleMap
	}
	return nil
}

type ThermostatScheduleDay struct {
	//Daily schedule is Multiple Blocks
	Blocks               []*ThermostatScheduleBlock `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ThermostatScheduleDay) Reset()         { *m = ThermostatScheduleDay{} }
func (m *ThermostatScheduleDay) String() string { return proto.CompactTextString(m) }
func (*ThermostatScheduleDay) ProtoMessage()    {}
func (*ThermostatScheduleDay) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{7}
}

func (m *ThermostatScheduleDay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermostatScheduleDay.Unmarshal(m, b)
}
func (m *ThermostatScheduleDay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermostatScheduleDay.Marshal(b, m, deterministic)
}
func (m *ThermostatScheduleDay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermostatScheduleDay.Merge(m, src)
}
func (m *ThermostatScheduleDay) XXX_Size() int {
	return xxx_messageInfo_ThermostatScheduleDay.Size(m)
}
func (m *ThermostatScheduleDay) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermostatScheduleDay.DiscardUnknown(m)
}

var xxx_messageInfo_ThermostatScheduleDay proto.InternalMessageInfo

func (m *ThermostatScheduleDay) GetBlocks() []*ThermostatScheduleBlock {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type ThermostatScheduleBlock struct {
	//heating setpoint
	//unit: celsius
	HeatingSetpoint *Double `protobuf:"bytes,1,opt,name=heating_setpoint,json=heatingSetpoint,proto3" json:"heating_setpoint,omitempty"`
	//cooling setpoint
	//unit: celsius
	CoolingSetpoint *Double `protobuf:"bytes,2,opt,name=cooling_setpoint,json=coolingSetpoint,proto3" json:"cooling_setpoint,omitempty"`
	//Current system mode of thermostat
	//unit: xbos/iot/HVACMode
	Mode HVACMode `protobuf:"varint,3,opt,name=mode,proto3,enum=xbospb.HVACMode" json:"mode,omitempty"`
	//Time when schedule block takes effect
	//format: RRule
	Time                 string   `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThermostatScheduleBlock) Reset()         { *m = ThermostatScheduleBlock{} }
func (m *ThermostatScheduleBlock) String() string { return proto.CompactTextString(m) }
func (*ThermostatScheduleBlock) ProtoMessage()    {}
func (*ThermostatScheduleBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{8}
}

func (m *ThermostatScheduleBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermostatScheduleBlock.Unmarshal(m, b)
}
func (m *ThermostatScheduleBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermostatScheduleBlock.Marshal(b, m, deterministic)
}
func (m *ThermostatScheduleBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermostatScheduleBlock.Merge(m, src)
}
func (m *ThermostatScheduleBlock) XXX_Size() int {
	return xxx_messageInfo_ThermostatScheduleBlock.Size(m)
}
func (m *ThermostatScheduleBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermostatScheduleBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ThermostatScheduleBlock proto.InternalMessageInfo

func (m *ThermostatScheduleBlock) GetHeatingSetpoint() *Double {
	if m != nil {
		return m.HeatingSetpoint
	}
	return nil
}

func (m *ThermostatScheduleBlock) GetCoolingSetpoint() *Double {
	if m != nil {
		return m.CoolingSetpoint
	}
	return nil
}

func (m *ThermostatScheduleBlock) GetMode() HVACMode {
	if m != nil {
		return m.Mode
	}
	return HVACMode_HVACModeOff
}

func (m *ThermostatScheduleBlock) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

type Meter struct {
	//unit: kW
	Power *Double `protobuf:"bytes,1,opt,name=power,proto3" json:"power,omitempty"`
	//unit: V
	Voltage *Double `protobuf:"bytes,2,opt,name=voltage,proto3" json:"voltage,omitempty"`
	//unit: kVA
	ApparentPower *Double `protobuf:"bytes,3,opt,name=apparent_power,json=apparentPower,proto3" json:"apparent_power,omitempty"`
	//unit: KWh
	Energy               *Double  `protobuf:"bytes,4,opt,name=energy,proto3" json:"energy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Meter) Reset()         { *m = Meter{} }
func (m *Meter) String() string { return proto.CompactTextString(m) }
func (*Meter) ProtoMessage()    {}
func (*Meter) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{9}
}

func (m *Meter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meter.Unmarshal(m, b)
}
func (m *Meter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meter.Marshal(b, m, deterministic)
}
func (m *Meter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meter.Merge(m, src)
}
func (m *Meter) XXX_Size() int {
	return xxx_messageInfo_Meter.Size(m)
}
func (m *Meter) XXX_DiscardUnknown() {
	xxx_messageInfo_Meter.DiscardUnknown(m)
}

var xxx_messageInfo_Meter proto.InternalMessageInfo

func (m *Meter) GetPower() *Double {
	if m != nil {
		return m.Power
	}
	return nil
}

func (m *Meter) GetVoltage() *Double {
	if m != nil {
		return m.Voltage
	}
	return nil
}

func (m *Meter) GetApparentPower() *Double {
	if m != nil {
		return m.ApparentPower
	}
	return nil
}

func (m *Meter) GetEnergy() *Double {
	if m != nil {
		return m.Energy
	}
	return nil
}

type Light struct {
	// True if the light is on
	//unit: on/off
	State *Bool `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	// 100 is maximum brightness
	Brightness           *Int64   `protobuf:"bytes,2,opt,name=brightness,proto3" json:"brightness,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Light) Reset()         { *m = Light{} }
func (m *Light) String() string { return proto.CompactTextString(m) }
func (*Light) ProtoMessage()    {}
func (*Light) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{10}
}

func (m *Light) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Light.Unmarshal(m, b)
}
func (m *Light) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Light.Marshal(b, m, deterministic)
}
func (m *Light) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Light.Merge(m, src)
}
func (m *Light) XXX_Size() int {
	return xxx_messageInfo_Light.Size(m)
}
func (m *Light) XXX_DiscardUnknown() {
	xxx_messageInfo_Light.DiscardUnknown(m)
}

var xxx_messageInfo_Light proto.InternalMessageInfo

func (m *Light) GetState() *Bool {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Light) GetBrightness() *Int64 {
	if m != nil {
		return m.Brightness
	}
	return nil
}

type EVSE struct {
	// maximum allowed current for charging
	//unit: A
	CurrentLimit *Double `protobuf:"bytes,1,opt,name=current_limit,json=currentLimit,proto3" json:"current_limit,omitempty"`
	// active charge current
	//unit: A
	Current *Double `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	// active charge voltage
	//unit: V
	Voltage *Double `protobuf:"bytes,3,opt,name=voltage,proto3" json:"voltage,omitempty"`
	// seconds left until car is charged
	//unit: seconds
	ChargingTimeLeft *Int32 `protobuf:"bytes,4,opt,name=charging_time_left,json=chargingTimeLeft,proto3" json:"charging_time_left,omitempty"`
	// charge state of the EVSE
	//unit: on/off
	State                *Bool    `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EVSE) Reset()         { *m = EVSE{} }
func (m *EVSE) String() string { return proto.CompactTextString(m) }
func (*EVSE) ProtoMessage()    {}
func (*EVSE) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{11}
}

func (m *EVSE) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EVSE.Unmarshal(m, b)
}
func (m *EVSE) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EVSE.Marshal(b, m, deterministic)
}
func (m *EVSE) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EVSE.Merge(m, src)
}
func (m *EVSE) XXX_Size() int {
	return xxx_messageInfo_EVSE.Size(m)
}
func (m *EVSE) XXX_DiscardUnknown() {
	xxx_messageInfo_EVSE.DiscardUnknown(m)
}

var xxx_messageInfo_EVSE proto.InternalMessageInfo

func (m *EVSE) GetCurrentLimit() *Double {
	if m != nil {
		return m.CurrentLimit
	}
	return nil
}

func (m *EVSE) GetCurrent() *Double {
	if m != nil {
		return m.Current
	}
	return nil
}

func (m *EVSE) GetVoltage() *Double {
	if m != nil {
		return m.Voltage
	}
	return nil
}

func (m *EVSE) GetChargingTimeLeft() *Int32 {
	if m != nil {
		return m.ChargingTimeLeft
	}
	return nil
}

func (m *EVSE) GetState() *Bool {
	if m != nil {
		return m.State
	}
	return nil
}

func init() {
	proto.RegisterEnum("xbospb.FanMode", FanMode_name, FanMode_value)
	proto.RegisterEnum("xbospb.HVACMode", HVACMode_name, HVACMode_value)
	proto.RegisterEnum("xbospb.HVACState", HVACState_name, HVACState_value)
	proto.RegisterType((*URI)(nil), "xbospb.URI")
	proto.RegisterType((*Triple)(nil), "xbospb.Triple")
	proto.RegisterType((*XBOSIoTDeviceState)(nil), "xbospb.XBOSIoTDeviceState")
	proto.RegisterType((*XBOSIoTDeviceActuation)(nil), "xbospb.XBOSIoTDeviceActuation")
	proto.RegisterType((*XBOSIoTContext)(nil), "xbospb.XBOSIoTContext")
	proto.RegisterType((*Thermostat)(nil), "xbospb.Thermostat")
	proto.RegisterType((*ThermostatSchedule)(nil), "xbospb.ThermostatSchedule")
	proto.RegisterMapType((map[string]*ThermostatScheduleDay)(nil), "xbospb.ThermostatSchedule.ScheduleMapEntry")
	proto.RegisterType((*ThermostatScheduleDay)(nil), "xbospb.ThermostatScheduleDay")
	proto.RegisterType((*ThermostatScheduleBlock)(nil), "xbospb.ThermostatScheduleBlock")
	proto.RegisterType((*Meter)(nil), "xbospb.Meter")
	proto.RegisterType((*Light)(nil), "xbospb.Light")
	proto.RegisterType((*EVSE)(nil), "xbospb.EVSE")
}

func init() { proto.RegisterFile("iot.proto", fileDescriptor_1804728d01c3c43d) }

var fileDescriptor_1804728d01c3c43d = []byte{
	// 1077 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0xc7, 0x43, 0x51, 0x87, 0x68, 0x14, 0xdb, 0xf4, 0x3a, 0x07, 0xc1, 0xf8, 0x3e, 0xd4, 0x60,
	0xd2, 0xd6, 0x75, 0x11, 0xa3, 0x91, 0x7b, 0x4a, 0x83, 0xa2, 0xf0, 0x29, 0xb0, 0x01, 0xbb, 0x36,
	0x56, 0x4e, 0x7a, 0x00, 0x5a, 0x81, 0xa2, 0x46, 0x16, 0x1b, 0x8a, 0xcb, 0x92, 0x4b, 0x39, 0xba,
	0x2a, 0xfa, 0x20, 0x7d, 0x8d, 0xde, 0x16, 0xe8, 0x5d, 0x5f, 0xa5, 0x57, 0x7d, 0x84, 0x62, 0x4f,
	0x14, 0xa5, 0x48, 0x86, 0xd3, 0xbb, 0xe5, 0xfe, 0x7f, 0xb3, 0xb3, 0x33, 0xb3, 0x33, 0x12, 0xd4,
	0x03, 0xc6, 0xb7, 0xe3, 0x84, 0x71, 0x46, 0xaa, 0xaf, 0xbb, 0x2c, 0x8d, 0xbb, 0xeb, 0x6b, 0x51,
	0x16, 0x86, 0x5e, 0x37, 0x44, 0x3e, 0x8e, 0x31, 0x55, 0xe2, 0xfa, 0xbd, 0x2b, 0xf4, 0xf8, 0x00,
	0x93, 0x4e, 0xca, 0x3d, 0x1e, 0xb0, 0x48, 0x6d, 0xbb, 0x4f, 0xc1, 0x7e, 0x41, 0x8f, 0xc9, 0xff,
	0xa0, 0x1e, 0x79, 0x43, 0x4c, 0x63, 0xcf, 0xc7, 0xa6, 0xb5, 0x61, 0x6d, 0xd6, 0xe9, 0x64, 0x83,
	0xdc, 0x85, 0xca, 0xc8, 0x0b, 0x33, 0x6c, 0x96, 0xa4, 0xa2, 0x3e, 0xdc, 0x5f, 0xa0, 0x7a, 0x91,
	0x04, 0x71, 0x88, 0xe4, 0x5d, 0xa8, 0xa5, 0x59, 0xf7, 0x27, 0xf4, 0xb9, 0xb4, 0x6d, 0xb4, 0x1a,
	0xdb, 0xea, 0x2a, 0xdb, 0x2f, 0xe8, 0x31, 0x35, 0x1a, 0xf9, 0x00, 0xea, 0x71, 0x82, 0xbd, 0xc0,
	0xf7, 0xb8, 0x3a, 0x6a, 0x06, 0x9c, 0xa8, 0xe4, 0x21, 0x54, 0x99, 0x3a, 0xd0, 0x7e, 0x93, 0xd3,
	0x92, 0xfb, 0x9b, 0x0d, 0xe4, 0xdb, 0xbd, 0xb3, 0xf6, 0x31, 0xbb, 0x38, 0xc0, 0x51, 0xe0, 0x63,
	0x9b, 0x0b, 0x5b, 0x02, 0x65, 0x1e, 0x0c, 0x55, 0x18, 0x65, 0x2a, 0xd7, 0x22, 0xbe, 0x04, 0x7f,
	0xce, 0x30, 0xe5, 0x41, 0x4f, 0xba, 0xb6, 0xe9, 0x64, 0x43, 0xc4, 0x87, 0x49, 0xc2, 0x12, 0xe9,
	0xac, 0x4e, 0xd5, 0x07, 0x69, 0x01, 0x88, 0x84, 0x0d, 0x99, 0xc8, 0x58, 0xb3, 0x2c, 0xef, 0x41,
	0xcc, 0x3d, 0x2e, 0x72, 0x85, 0x16, 0x28, 0xf2, 0x10, 0x2a, 0x43, 0xe4, 0x98, 0x34, 0x2b, 0x12,
	0x5f, 0x32, 0xf8, 0xa9, 0xd8, 0xa4, 0x4a, 0x13, 0x50, 0x18, 0x5c, 0x0e, 0x78, 0xb3, 0x3a, 0x0d,
	0x9d, 0x88, 0x4d, 0xaa, 0x34, 0xb2, 0x01, 0x65, 0x1c, 0xa5, 0xd8, 0xac, 0x49, 0xe6, 0x8e, 0x61,
	0x0e, 0x5f, 0xb6, 0x0f, 0xa9, 0x54, 0xc8, 0x57, 0xb0, 0x32, 0x53, 0xd3, 0xe6, 0x6d, 0x09, 0xdf,
	0x37, 0xf0, 0x37, 0x4a, 0x6e, 0x2b, 0x95, 0x2e, 0x5f, 0x4d, 0x7d, 0x93, 0x1f, 0x61, 0x7d, 0xe6,
	0x80, 0x8e, 0xaa, 0x80, 0x3c, 0xab, 0x2e, 0xcf, 0xda, 0x98, 0x7f, 0xd6, 0x79, 0xce, 0xd1, 0xe6,
	0xd5, 0x02, 0xc5, 0xfd, 0xdb, 0x82, 0xfb, 0x53, 0xf5, 0xd9, 0xf5, 0x79, 0xa6, 0x5c, 0xbf, 0x7d,
	0x8d, 0xa6, 0xab, 0x61, 0xbf, 0x5d, 0x35, 0xca, 0x37, 0xa9, 0x46, 0xe5, 0x06, 0xd5, 0xa8, 0x2e,
	0xaa, 0x86, 0xfb, 0x35, 0x2c, 0xeb, 0x58, 0xf7, 0x59, 0xc4, 0xf1, 0x35, 0x9f, 0x1b, 0xe3, 0x26,
	0xd4, 0x7c, 0x25, 0x37, 0x4b, 0x1b, 0xf6, 0x66, 0xa3, 0xb5, 0x9c, 0x87, 0x20, 0x5b, 0x89, 0x1a,
	0xd9, 0xfd, 0xa3, 0x0c, 0x30, 0x09, 0x8b, 0x7c, 0x04, 0x0d, 0x8e, 0xc3, 0x18, 0x13, 0x8f, 0x67,
	0x09, 0xea, 0x36, 0xcb, 0x8d, 0x0f, 0x58, 0xd6, 0x0d, 0x91, 0x16, 0x11, 0xf2, 0x0c, 0x56, 0x13,
	0x0c, 0x3d, 0x1e, 0x8c, 0xb0, 0x33, 0xc8, 0x86, 0x41, 0x2f, 0xe0, 0x63, 0xdd, 0x75, 0xb3, 0x76,
	0x8e, 0x01, 0x8f, 0x34, 0x47, 0x36, 0xe1, 0x36, 0x1b, 0x61, 0x92, 0x04, 0x3d, 0xd4, 0xb9, 0xce,
	0x63, 0xde, 0x63, 0x2c, 0xa4, 0xb9, 0x2a, 0x9a, 0xba, 0xef, 0x45, 0xf2, 0x01, 0xa1, 0xce, 0xf3,
	0x0c, 0xda, 0xf7, 0x22, 0xd5, 0x98, 0x5b, 0x20, 0xd6, 0x9d, 0x21, 0xeb, 0xa1, 0x4c, 0xf6, 0x72,
	0x6b, 0xc5, 0x90, 0xcf, 0xbd, 0xe8, 0x94, 0xf5, 0x90, 0xd6, 0xfa, 0x6a, 0x41, 0x1e, 0x41, 0x59,
	0x72, 0x55, 0xc9, 0x39, 0x86, 0x3b, 0x7a, 0xb9, 0xbb, 0x2f, 0x41, 0xa9, 0x92, 0xf7, 0xa1, 0xa2,
	0x1c, 0xd7, 0x24, 0xb6, 0x5a, 0xc4, 0xa4, 0x4f, 0xaa, 0x74, 0xf2, 0x25, 0xac, 0x61, 0x24, 0x46,
	0x62, 0xaf, 0x33, 0x40, 0x8f, 0x8b, 0xeb, 0x5e, 0x62, 0xaa, 0xfb, 0x25, 0x2f, 0xf9, 0x71, 0xc4,
	0x77, 0x5a, 0x74, 0x55, 0x93, 0x47, 0xe8, 0xf1, 0xb6, 0xe4, 0x8a, 0xe6, 0x3e, 0x63, 0xa1, 0x31,
	0xaf, 0x5f, 0x67, 0xbe, 0xcf, 0x58, 0xa8, 0xcd, 0x9f, 0x82, 0x23, 0xbc, 0x06, 0xd1, 0x65, 0x27,
	0x45, 0x1e, 0xb3, 0x20, 0xe2, 0x4d, 0x98, 0x5b, 0x89, 0x15, 0xcd, 0xb5, 0x35, 0x26, 0x4c, 0x85,
	0xc7, 0x29, 0xd3, 0xc6, 0x7c, 0x53, 0xcd, 0x19, 0x53, 0xf7, 0x4f, 0x0b, 0xc8, 0xe4, 0x05, 0xb5,
	0xfd, 0x01, 0xf6, 0xb2, 0x10, 0xc9, 0x29, 0x34, 0x52, 0xbd, 0x3e, 0xf5, 0xe2, 0xa6, 0x25, 0x9f,
	0xe1, 0x87, 0x6f, 0x76, 0x92, 0x31, 0xd8, 0x6e, 0x4f, 0xe8, 0xc3, 0x88, 0x27, 0x63, 0x5a, 0xb4,
	0x5f, 0xff, 0x01, 0x9c, 0x59, 0x80, 0x38, 0x60, 0xbf, 0xc2, 0xb1, 0xfe, 0x1d, 0x11, 0x4b, 0xb2,
	0x53, 0xfc, 0x05, 0x69, 0xb4, 0xfe, 0xbf, 0xd8, 0xdd, 0x81, 0x37, 0xd6, 0x3f, 0x30, 0x5f, 0x94,
	0x3e, 0xb7, 0xdc, 0x73, 0xb8, 0x37, 0x97, 0x21, 0x9f, 0x41, 0xb5, 0x1b, 0x32, 0xff, 0x55, 0xaa,
	0x23, 0x78, 0x67, 0xf1, 0x91, 0x7b, 0x82, 0xa3, 0x1a, 0x77, 0xff, 0xb2, 0xe0, 0xc1, 0x02, 0x66,
	0x6e, 0xa1, 0xac, 0xff, 0x5e, 0xa8, 0xd2, 0x8d, 0x0a, 0x95, 0xbf, 0x75, 0xfb, 0xda, 0xb7, 0x6e,
	0xc6, 0x49, 0x59, 0x66, 0x55, 0xae, 0xdd, 0xdf, 0x2d, 0xa8, 0xc8, 0x61, 0x46, 0x1e, 0x41, 0x25,
	0x66, 0x57, 0x98, 0x2c, 0xb8, 0xae, 0x12, 0xc5, 0xf8, 0x19, 0xb1, 0x50, 0x3c, 0xca, 0x05, 0x77,
	0x33, 0x32, 0xf9, 0x04, 0x96, 0xbd, 0x38, 0xf6, 0x12, 0x8c, 0x78, 0x47, 0x1d, 0x6c, 0xcf, 0x35,
	0x58, 0x32, 0xd4, 0xb9, 0x74, 0xf0, 0x1e, 0x54, 0x31, 0xc2, 0xe4, 0x72, 0xac, 0x47, 0xc1, 0x2c,
	0xae, 0x55, 0xf7, 0x7b, 0xa8, 0xc8, 0xf9, 0x4a, 0x5c, 0xd3, 0xc1, 0xd6, 0x9c, 0xd1, 0xa1, 0x9b,
	0xf7, 0x31, 0x40, 0x37, 0x11, 0x74, 0x84, 0x69, 0xaa, 0x2f, 0x5e, 0x6c, 0xba, 0x4f, 0x3f, 0xa6,
	0x05, 0xc0, 0xfd, 0xc7, 0x82, 0xb2, 0x18, 0xcc, 0x64, 0x07, 0x96, 0xfc, 0x2c, 0x91, 0x21, 0x84,
	0xc1, 0x30, 0x58, 0x54, 0xca, 0x3b, 0x1a, 0x3a, 0x11, 0x8c, 0x9c, 0xd0, 0xea, 0x7b, 0x51, 0x8a,
	0xb4, 0x5c, 0x4c, 0xa6, 0x7d, 0x7d, 0x32, 0x9f, 0x01, 0xf1, 0x07, 0x5e, 0x72, 0x29, 0x1e, 0x87,
	0xa8, 0x5b, 0x27, 0xc4, 0x3e, 0x9f, 0xfd, 0x51, 0x52, 0xd3, 0xc3, 0x31, 0xe0, 0x45, 0x30, 0xc4,
	0x13, 0xec, 0x17, 0x32, 0x54, 0x59, 0x98, 0xa1, 0xad, 0xc7, 0x50, 0xd3, 0x13, 0x94, 0x34, 0xe4,
	0x72, 0x37, 0xe3, 0xcc, 0xb9, 0x45, 0xea, 0x50, 0x79, 0xee, 0x45, 0x67, 0x91, 0x63, 0x11, 0x80,
	0xaa, 0x58, 0xf6, 0xfb, 0x4e, 0x69, 0xeb, 0x3b, 0xb8, 0x6d, 0x1e, 0x17, 0x59, 0x81, 0x86, 0x59,
	0x0b, 0xf1, 0x16, 0xb9, 0x0b, 0x8e, 0xd9, 0x10, 0x13, 0xf0, 0x2c, 0x0a, 0xc7, 0x8e, 0x55, 0xdc,
	0x15, 0x83, 0x4d, 0xee, 0x96, 0x88, 0x03, 0x77, 0xcc, 0xae, 0xf4, 0x68, 0x6f, 0xfd, 0x6a, 0x41,
	0x3d, 0x9f, 0xbe, 0x46, 0x97, 0x1f, 0xea, 0xf4, 0x07, 0xb0, 0x96, 0xef, 0xe4, 0x03, 0xf6, 0x89,
	0x63, 0x4d, 0x09, 0xf9, 0xe8, 0x7c, 0xe2, 0x94, 0xe6, 0x5b, 0xb4, 0x1c, 0x7b, 0xbe, 0x45, 0xcb,
	0x29, 0x77, 0xab, 0xf2, 0xaf, 0xed, 0xce, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xfd, 0x01,
	0x18, 0x1b, 0x0b, 0x00, 0x00,
}
