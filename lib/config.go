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
	ServiceConfig
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
func (config *Config) Load() ServiceConfig {
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

	// TODO: Add default values.
	services := []ServiceConfig{
		{
			Name:    ``,
			APIKey:  ``,
			URL:     ``,
			Symbols: make([]string, 0),
		},
	}

	data, err := json.Marshal(services)

	if err != nil {
		log.Fatal("Cannot encode ", err)
	}

	ioutil.WriteFile(config.file, data, 0644)
}

//
// Parse the configuration JSON key/value pairs.
//
func (config *Config) read(data []byte) ServiceConfig {
	results := []ServiceConfig{}

	err := json.Unmarshal(data, &results)

	if err != nil {
		log.Fatal("Cannot decode ", err)
	}

	var service ServiceConfig

	for i := 0; i < len(results); i++ {
		service = results[i]

		if service.Name == config.Name {
			break
		}
	}

	// Return struct reference.
	return service
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
