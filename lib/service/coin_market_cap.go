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
	endpoint "github.com/nuxy/go-crypto-market-ui/lib/service/coin_market_cap"
)

//
// Service provides endpoint methods.
//
type Service interface {
	URL(endpointName string) string
}

//
// CoinMarketCap declared data types.
//
type CoinMarketCap struct {}

//
// URL returns an unprocessed location.
//
func (CoinMarketCap) URL(endpointName string) string {
	var rawURI string

	switch endpointName {
		case "Quotes":
			rawURI = (endpoint.Quotes{}).URI()
			break
	}

	return "https://pro-api.coinmarketcap.com/v1/" + rawURI + "&CMC_PRO_API_KEY=%s"
}
