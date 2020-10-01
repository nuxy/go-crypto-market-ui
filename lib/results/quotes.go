//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package results

//
// Quotes declared data types.
//
type Quotes struct {
	Name             string
	Symbol           string
	TotalSupply      int64
    Price            float64
	Volume24h        float64
	PercentChange1h  float64
	PercentChange24h float64
	PercentChange7d  float64
	MarketCap        float64
}
