//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"
	"path"
)

//
// Config declared data types.
//
type Config struct {
	APIConfig
	file string
}

//
// NewConfig create a new config instance.
//
func NewConfig(serviceName string) *Config {
	config := &Config{}
	config.Name = serviceName
	config.file = config.filePath(".crypto-market-ui.json")
	return config
}

//
// Load config file, or create when it doesn't exist.
//
func (config *Config) Load() APIConfig {
	data, err := ioutil.ReadFile(config.file)

	if err != nil {
		config.create()
	}

	return config.read(data)
}

//
// Create the locally hosted configuration file.
//
func (config *Config) create() {
	apis := []APIConfig{}

	data, err := json.Marshal(apis)

	if err != nil {
		log.Fatal("Cannot encode ", err)
	}

	ioutil.WriteFile(config.file, data, 0644)
}

//
// Parse the configuration JSON key/value pairs.
///
func (config *Config) read(data []byte) APIConfig {
	results := []APIConfig{}

	err := json.Unmarshal(data, &results)

	if err != nil {
		log.Fatal("Cannot decode ", err)
	}

	var api APIConfig

	for i := 0; i < len(results); i++ {
		api = results[i]

		if api.Name == config.Name {
			break
		}
	}

	// Return struct reference.
	return api
}

//
// Return file including relative path to $HOME
//
func (config *Config) filePath(name string) string {
	usr, err := user.Current()

	if err != nil {
		log.Fatal("Cannot open path ", err)
	}

	return path.Join(usr.HomeDir, name)
}
