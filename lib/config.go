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
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

//
// ConfigHolding declared data types.
//
type ConfigHolding struct {
	Units float64 `yaml:"units"`
	Price float64 `yaml:"price"`
}

type symbol []ConfigHolding

//
// ConfigSymbol declared data types.
//
type ConfigSymbol map[string]symbol

//
// ConfigOpts declared data types.
//
type ConfigOpts struct {
	ServiceName string       `yaml:"service"`
	APIKey      string       `yaml:"apiKey"`
	Currency    string       `yaml:"currency"`
	Language    string       `yaml:"language"`
	RefreshRate int64        `yaml:"refreshRate"`
	Symbols     ConfigSymbol `yaml:"symbols"`
}

//
// Config declared data types.
//
type Config struct {
	options ConfigOpts
	file    string
}

//
// NewConfig create a new config instance.
//
func NewConfig(serviceName string) *Config {
	config := &Config{}
	config.file = filePath(".crypto-market-ui.yml")
	config.init(serviceName)
	return config
}

//
// Options returns ConfigOpts name/value(s).
//
func (config *Config) Options() ConfigOpts {
	return config.options
}

//
// ServiceName sets/returns the option value.
//
func (config *Config) ServiceName(v ...string) string {
	if len(v) == 1 {
		config.options.ServiceName = v[0]
	}

	return config.options.ServiceName
}

//
// APIKey sets/returns the option value.
//
func (config *Config) APIKey(v ...string) string {
	if len(v) == 1 {
		config.options.APIKey = v[0]
	}

	return config.options.APIKey
}

//
// Currency sets/returns the option value.
//
func (config *Config) Currency(v ...string) string {
	if len(v) == 1 {
		config.options.Currency = v[0]
	}

	return config.options.Currency
}

//
// Language sets/returns the option value.
//
func (config *Config) Language(v ...string) string {
	if len(v) == 1 {
		config.options.Language = v[0]
	}

	return config.options.Language
}

//
// RefreshRate sets/returns the option value.
//
func (config *Config) RefreshRate(v ...string) int64 {
	if len(v) == 1 {
		i, err := strconv.ParseInt(v[0], 10, 64)

		if err != nil {
			i = 15
		}

		config.options.RefreshRate = i
	}

	return config.options.RefreshRate
}

//
// Symbols sets/returns the option value.
//
func (config *Config) Symbols(v ...ConfigSymbol) ConfigSymbol {
	if len(v) == 1 {
		config.options.Symbols = v[0]
	}

	return config.options.Symbols
}

//
// Save the config field values.
//
func (config *Config) Save() {
	config.write(config.options)
}

//
// IsValid checks config values exist.
//
func (config *Config) IsValid() bool {
	var valid bool

	if len(config.ServiceName()) > 0 && len(config.APIKey()) > 0 && len(config.Currency()) > 0 && len(config.Language()) > 0 && len(config.Symbols()) > 0 && config.RefreshRate() > 0 {
		valid = true
	}

	return valid
}

//
// Initialize a new config, if none exists.
//
func (config *Config) init(serviceName string) {
	_, err := os.Stat(config.file)

	if os.IsNotExist(err) {
		config.create()
	}

	config.read(serviceName)
}

//
// Create a new config with default values.
//
func (config *Config) create() {
	config.write(
		ConfigOpts{
			ServiceName: "CoinMarketCap",
			Currency:    "USD",
			Language:    "en",
			Symbols:     ConfigSymbol{"BTC": make([]ConfigHolding, 1)},
			RefreshRate: 15,
		},
	)
}

//
// Read selected config YAML key/value pairs.
//
func (config *Config) read(serviceName string) {
	data, err := ioutil.ReadFile(config.file)

	if err != nil {
		log.Fatal("Cannot read ", err)
	}

	options := []ConfigOpts{}

	err = yaml.Unmarshal(data, &options)

	if err != nil {
		log.Fatal("Cannot decode ", err)
	}

	var option ConfigOpts

	for i := 0; i < len(options); i++ {
		option = options[i]

		if serviceName == option.ServiceName {
			config.options = option
			break
		}
	}
}

//
// Write the locally hosted config file.
//
func (config *Config) write(file ConfigOpts) {
	apis := []ConfigOpts{}

	// TODO: Support multiple services.
	apis = append(apis, file)

	data, err := yaml.Marshal(&apis)

	if err != nil {
		log.Fatal("Cannot encode ", err)
	}

	ioutil.WriteFile(config.file, data, 0644)
}

//
// Returns the file relative path to $HOME
//
func filePath(name string) string {
	usr, err := user.Current()

	if err != nil {
		log.Fatal("Cannot open path ", err)
	}

	return path.Join(usr.HomeDir, name)
}
