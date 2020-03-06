//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

// ConfigurationModel represent the configuration model
type ConfigurationModel struct {
	RunMode        string `json:"run_mode"`
	Port           string `json:"port"`
	DialectMaster  string `json:"dialect_master"`
	DsnMaster      string `json:"dsn_master"`
	IdleConnMaster int    `json:"idle_conn_master"`
	MaxConnMaster  int    `json:"max_conn_master"`
	DialectSlave   string `json:"dialect_slave"`
	DsnSlave       string `json:"dsn_slave"`
	IdleConnSlave  int    `json:"idle_conn_slave"`
	MaxConnSlave   int    `json:"max_conn_slave"`
}

var (
	// Configuration represent the variable of configuration model
	Configuration = &ConfigurationModel{}
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	basepath = strings.Replace(basepath, "config", "", -1)
	file := basepath + "config.json"
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("Failed to load auth configuration file: %s", err.Error()))
	}

	err = json.Unmarshal(raw, Configuration)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse auth configuration file: %s", err.Error()))
	}
}
