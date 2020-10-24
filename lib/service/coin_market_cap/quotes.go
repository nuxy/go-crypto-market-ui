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
	"github.com/nuxy/go-crypto-market-ui/lib/service/coin_market_cap/quotes"
)

//
// Quotes declared data types.
//
type Quotes struct{}

//
// URI returns an unprocessed path.
//
func (Quotes) URI() string {
	return "cryptocurrency/quotes/latest"
}

//
// Params returns query parameters.
//
func (Quotes) Params() []string {
	return []string{"CMC_PRO_API_KEY={APIKey}", "convert={Currency}", "symbol={Symbols}"}
}

//
// Parse returns API response body data.
//
func (Quotes) Parse(body []byte) interface{} {
	return quotes.Parse(body)
}
