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
	"fmt"
	"strings"

	api "github.com/nuxy/go-crypto-market-ui/lib/service"
)

//
// ServiceConfig declared data types.
//
type ServiceConfig struct {
	Name    string   `json:"name"`
	APIKey  string   `json:"apiKey"`
	Symbols []string `json:"symbols"`
}

//
// Service declared data types.
//
type Service struct {
	Config       ServiceConfig
	instance     api.Service
	endpointName string
}

//
// NewService creates a new service instance.
//
func NewService(config ServiceConfig, endpointName string) *Service {
	service := &Service{}
	service.Config       = config
	service.endpointName = endpointName
	service.init()
	return service
}

//
// Assigns selected runtime interface.
//
func (service *Service) init() {
	switch service.Config.Name {
		case "CoinMarketCap":
			service.instance = (api.CoinMarketCap{})
			break
	}
}

//
// URL returns as constructed location.
//
func (service *Service) URL() string {
	rawURL  := service.instance.URL(service.endpointName)
	symbols := strings.Join(service.Config.Symbols, ",")

	return fmt.Sprintf(rawURL, symbols, service.Config.APIKey)
}
