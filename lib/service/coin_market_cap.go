//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package service

import (
	"github.com/nuxy/go-crypto-market-ui/lib/common"
	endpoint "github.com/nuxy/go-crypto-market-ui/lib/service/coin_market_cap"
)

//
// CoinMarketCap declared data types.
//
type CoinMarketCap struct{}

//
// URL returns an unprocessed location.
//
func (service CoinMarketCap) URL(endpointName string) string {
	return "https://pro-api.coinmarketcap.com/v1/" + service.endpointInterface(endpointName).URI()
}

//
// Params returns query parameters for a given endpoint.
//
func (service CoinMarketCap) Params(endpointName string) []string {
	return service.endpointInterface(endpointName).Params()
}

//
// Parse returns API response body data for a given endpoint.
//
func (service CoinMarketCap) Parse(endpointName string, body []byte) interface{} {
	return service.endpointInterface(endpointName).Parse(body)
}

//
// Assign runtime selected interface for a given endpoint.
//
func (service CoinMarketCap) endpointInterface(endpointName string) common.EndpointInterface {
	var instance common.EndpointInterface

	switch endpointName {
	case "Metadata":
		instance = (endpoint.Metadata{})
		break

	case "Quotes":
		instance = (endpoint.Quotes{})
		break
	}

	return instance
}
