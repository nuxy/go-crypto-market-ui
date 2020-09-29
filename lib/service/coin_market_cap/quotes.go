//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package coinmarketcap

//
// Endpoint provides runtime methods.
//
type Endpoint interface {
	URI() string
}

//
// Quotes declared data types.
//
type Quotes struct{}

//
// URI returns an unprocessed path.
//
func (Quotes) URI() string {
	return "cryptocurrency/quotes/latest?symbol=%s"
}
