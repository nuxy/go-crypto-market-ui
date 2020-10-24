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
	"strings"

	"github.com/nuxy/go-crypto-market-ui/lib/common"
	"github.com/nuxy/go-crypto-market-ui/lib/service"
)

//
// API declared data types.
//
type API struct {
	Config       *Config
	endpointName string
}

//
// NewAPI creates a new service instance.
//
func NewAPI(config *Config, endpointName string) *API {
	api := &API{}
	api.Config       = config
	api.endpointName = endpointName
	return api
}

//
// URL returns as constructed location.
//
func (api *API) URL() string {
	return api.serviceInterface().URL(api.endpointName) + "?" + api.params()
}

//
// Returns the endpoint query parameters.
//
func (api *API) params() string {
	values := api.serviceInterface().Params(api.endpointName)
	params := make([]string, len(values))

	for i, v := range values {
		field := strings.Split(v, "=")

		switch field[1] {
		case "{APIKey}":
			params[i] = field[0] + "=" + api.Config.APIKey()
		case "{Currency}":
			params[i] = field[0] + "=" + api.Config.Currency()
		case "{Language}":
			params[i] = field[0] + "=" + api.Config.Language()
		case "{Symbols}":
			params[i] = field[0] + "=" + api.symbols()
		default:
			params[i] = field[0] + "=" + field[1]
		}
	}

	return strings.Join(params, "&")
}

//
// Parse returns API response body data.
//
func (api *API) Parse(body []byte) interface{} {
	return api.serviceInterface().Parse(api.endpointName, body)
}

//
// Returns config defined Symbols.
//
func (api *API) symbols() string {
	items := api.Config.Symbols()

	values := make([]string, 0, len(items))

	for k := range items {
		values = append(values, k)
	}

	return strings.Join(values, ",")
}

//
// Assigns runtime selected interface.
//
func (api *API) serviceInterface() common.ServiceInterface {
	var instance common.ServiceInterface

	// TODO: Support multiple services.
	switch api.Config.ServiceName() {
	case "CoinMarketCap":
		instance = (service.CoinMarketCap{})
		break
	}

	return instance
}
