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
type CoinMarketCap struct {
	instance common.EndpointInterface
}

//
// Assign interface for a given endpoint.
//
func assignInterface(endpointName string) *CoinMarketCap {
	service := &CoinMarketCap{}

	switch endpointName {
	case "Metadata":
		service.instance = (endpoint.Metadata{})
		break

	case "Quotes":
		service.instance = (endpoint.Quotes{})
		break
	}

	return service
}

//
// URL returns an unprocessed location.
//
func (CoinMarketCap) URL(endpointName string) string {
	service := assignInterface(endpointName)

	return "https://pro-api.coinmarketcap.com/v1/" + service.instance.URI() + "&CMC_PRO_API_KEY=%s"
}

//
// Parse returns API response body data a given endpoint.
//
func (CoinMarketCap) Parse(endpointName string, body []byte) interface{} {
	service := assignInterface(endpointName)

	return service.instance.Parse(body)
}
