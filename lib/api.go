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

	"github.com/nuxy/go-crypto-market-ui/lib/common"
	"github.com/nuxy/go-crypto-market-ui/lib/service"
)

//
// APIConfig declared data types.
//
type APIConfig struct {
	Name    string   `json:"name"`
	APIKey  string   `json:"apiKey"`
	Symbols []string `json:"symbols"`
}

//
// API declared data types.
//
type API struct {
	Config       APIConfig
	instance     common.ServiceInterface
	endpointName string
}

//
// NewAPI creates a new service instance.
//
func NewAPI(config APIConfig, endpointName string) *API {
	api := &API{}
	api.Config       = config
	api.endpointName = endpointName
	api.assignInterface()
	return api
}

//
// Assigns selected runtime interface.
//
func (api *API) assignInterface() {
	switch api.Config.Name {
	case "CoinMarketCap":
		api.instance = (service.CoinMarketCap{})
		break
	}
}

//
// URL returns as constructed location.
//
func (api *API) URL() string {
	return fmt.Sprintf(api.rawURL(), api.symbols(), api.Config.APIKey)
}

//
// Schema returns response data types.
//
func (api *API) Schema() interface{} {
	return api.instance.Schema(api.endpointName)
}

//
// Returns the endpoint defined raw URL.
//
func (api *API) rawURL() string {
	return api.instance.URL(api.endpointName)
}

//
// Returns configuration defined Symbols.
//
func (api *API) symbols() string {
	return strings.Join(api.Config.Symbols, ",")
}
