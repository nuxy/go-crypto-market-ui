//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package coinmarketcap

import (
	"github.com/nuxy/go-crypto-market-ui/lib/service/coin_market_cap/metadata"
)

//
// Metadata declared data types.
//
type Metadata struct{}

//
// URI returns an unprocessed path.
//
func (Metadata) URI() string {
	return "cryptocurrency/info"
}

//
// Params returns query parameters.
//
func (Metadata) Params() []string {
	return []string{"CMC_PRO_API_KEY={APIKey}", "symbol={Symbols}"}
}

//
// Parse returns API response body data.
//
func (Metadata) Parse(body []byte) interface{} {
	return metadata.Parse(body)
}
