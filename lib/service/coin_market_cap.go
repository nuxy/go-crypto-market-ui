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
	instance     common.EndpointInterface
	endpointName string
}

//
// Create a private service instance.
//
func newCoinMarketCap(endpointName string) *CoinMarketCap {
	service := &CoinMarketCap{}
	service.endpointName = endpointName
	service.init()
	return service
}

//
// Assign select endpoint interface.
//
func (service *CoinMarketCap) init() {
	switch service.endpointName {
	case "Quotes":
		service.instance = (endpoint.Quotes{})
		break
	}
}

//
// URL returns an unprocessed location.
//
func (CoinMarketCap) URL(endpointName string) string {
	service := newCoinMarketCap(endpointName)

	return "https://pro-api.coinmarketcap.com/v1/" + service.instance.URI() + "&CMC_PRO_API_KEY=%s"
}

//
// Schema returns response data types a given endpoint.
//
func (CoinMarketCap) Schema(endpointName string) interface{} {
	service := newCoinMarketCap(endpointName)

	return service.instance.Schema()
}
