package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gtfierro/xboswave/ingester/types"
	"github.com/immesys/wavemq/mqpb"
	_ "github.com/mattn/go-sqlite3"
	logrus "github.com/sirupsen/logrus"
)

// TODO: these go into sqlite table
type ArchiveRequest struct {
	Schema string
	Plugin string
	URI    types.SubscriptionURI

	// reported
	Inserted       time.Time
	LastError      string
	ErrorTimestamp time.Time
}

type subscription struct {
	S    mqpb.WAVEMQ_SubscribeClient
	stop chan struct{}
	uri  types.SubscriptionURI
}

type ConfigManager struct {
	db *sql.DB
}

func NewCfgManager(cfg *Config) (*ConfigManager, error) {
	var err error
	cfgmgr := &ConfigManager{}
	cfgmgr.db, err = sql.Open("sqlite3", cfg.Store.Path)
	if err != nil {
		return nil, err
	}

	// set up tables
	_, err = cfgmgr.db.Exec(`CREATE TABLE IF NOT EXISTS requests (
        schema  TEXT NOT NULL,
        plugin  TEXT NOT NULL,
        namespace TEXT NOT NULL,
        resource TEXT NOT NULL,
        inserted DATETIME DEFAULT CURRENT_TIMESTAMP,
        lastError TEXT,
        errorTimestamp DATETIME DEFAULT NULL
    );`)

	return cfgmgr, err
}

type RequestFilter struct {
	Schema    *string
	Plugin    *string
	Namespace *string
	Resource  *string
	HasError  *bool
	//TODO: inserted time
}

func (cfgmgr *ConfigManager) List(filter *RequestFilter) ([]ArchiveRequest, error) {
	stmt := "SELECT DISTINCT schema, plugin, namespace, resource FROM requests"
	var results []ArchiveRequest

	if filter == nil {
		stmt += ";"
	} else {
		var filters []string
		if filter.Schema != nil {
			filters = append(filters, fmt.Sprintf("schema='%s' ", *filter.Schema))
		}
		if filter.Plugin != nil {
			filters = append(filters, fmt.Sprintf("plugin='%s' ", *filter.Plugin))
		}
		if filter.Namespace != nil {
			filters = append(filters, fmt.Sprintf("namespace='%s' ", *filter.Namespace))
		}
		if filter.Resource != nil {
			filters = append(filters, fmt.Sprintf("resource='%s' ", *filter.Resource))
		}
		if filter.HasError != nil {
			if *filter.HasError {
				filters = append(filters, fmt.Sprint("lastError != '' "))
			} else {
				filters = append(filters, fmt.Sprint("lastError = '' "))
			}
		}
		stmt = fmt.Sprintf("%s WHERE %s;", stmt, strings.Join(filters, " AND "))
	}

	rows, err := cfgmgr.db.Query(stmt)
	if err != nil {
		return results, err
	}
	defer rows.Close()
	for rows.Next() {
		req := &ArchiveRequest{
			URI: types.SubscriptionURI{},
		}
		if err := rows.Scan(&req.Schema, &req.Plugin, &req.URI.Namespace, &req.URI.Resource); err != nil {
			logrus.Fatal(err)
		}
		results = append(results, *req)
	}

	return results, nil
}

func (cfgmgr *ConfigManager) Add(req ArchiveRequest) error {
	filter := &RequestFilter{
		Schema:    &req.Schema,
		Plugin:    &req.Plugin,
		Namespace: &req.URI.Namespace,
		Resource:  &req.URI.Resource,
	}
	results, err := cfgmgr.List(filter)
	if err != nil {
		return err
	} else if len(results) == 0 {
		stmt := "INSERT INTO requests(schema, plugin, namespace, resource) VALUES (?, ?, ?, ?);"
		_, err := cfgmgr.db.Exec(stmt, req.Schema, req.Plugin, req.URI.Namespace, req.URI.Resource)
		return err
	}

	return nil
}

// deletes the request from the config manager; returns true if anyone else uses the same subscription URI
func (cfgmgr *ConfigManager) Delete(req ArchiveRequest) (bool, error) {

	stmt := "DELETE FROM requests WHERE schema=? AND plugin=? AND namespace=? AND resource=?;"
	_, err := cfgmgr.db.Exec(stmt, req.Schema, req.Plugin, req.URI.Namespace, req.URI.Resource)
	if err != nil {
		return true, err
	}

	filter := &RequestFilter{
		Namespace: &req.URI.Namespace,
		Resource:  &req.URI.Resource,
	}
	existingSubs, err := cfgmgr.List(filter)
	return len(existingSubs) > 0, err
}

func (cfgmgr *ConfigManager) MarkErrorURI(uri types.SubscriptionURI, subErr string) error {
	stmt := "UPDATE requests SET lastError = ?, errorTimestamp = ? WHERE namespace = ? AND resource = ?"
	_, err := cfgmgr.db.Exec(stmt, subErr, time.Now(), uri.Namespace, uri.Resource)
	return err
}

func (cfgmgr *ConfigManager) ClearErrorURI(uri types.SubscriptionURI) error {
	stmt := "UPDATE requests SET lastError = '', errorTimestamp = '' WHERE namespace = ? AND resource = ?"
	_, err := cfgmgr.db.Exec(stmt, uri.Namespace, uri.Resource)
	return err
}

func (cfgmgr *ConfigManager) Status() ([]ArchiveRequest, error) {
	stmt := "SELECT schema, plugin, namespace, resource, inserted, coalesce(lastError, ''), errorTimestamp FROM requests;"
	var results []ArchiveRequest
	rows, err := cfgmgr.db.Query(stmt)
	if err != nil {
		return results, err
	}
	defer rows.Close()
	for rows.Next() {
		row := &ArchiveRequest{
			URI: types.SubscriptionURI{},
		}
		var et interface{}
		if err := rows.Scan(&row.Schema, &row.Plugin, &row.URI.Namespace, &row.URI.Resource, &row.Inserted, &row.LastError, &et); err != nil {
			return results, err
		}
		if et != nil {
			row.ErrorTimestamp = et.(time.Time)
		} else {
			row.ErrorTimestamp = time.Unix(0, 0)
		}

		results = append(results, *row)
	}
	return results, nil
}
