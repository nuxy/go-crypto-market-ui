//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package lib

//
// ServiceConfig declared data types.
//
type ServiceConfig struct {
	Name    string   `json:"name"`
	APIKey  string   `json:"apiKey"`
	URL     string   `json:"url"`
	Symbols []string `json:"symbols"`
}

//
// Service declared data types.
//
type Service struct {
	Config ServiceConfig
}

//
// NewService creates a new service instance.
//
func NewService(config ServiceConfig) *Service {
	service := &Service{}
	service.Config = config
	return service
}

//
// URL returns a valid resource identifier.
//
func (service *Service) URL() string {
	return service.Config.URL
}
