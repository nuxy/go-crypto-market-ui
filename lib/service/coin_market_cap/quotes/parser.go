//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package quotes

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/nuxy/go-crypto-market-ui/lib/results"
)

//
// Parse returns API response body as results type.
//
func Parse(body []byte) interface{} {
	var schema Response

	err := json.Unmarshal(body, &schema)

	if err != nil {
		log.Fatal("Parser error: ", err)
	}

	status := schema.Status

	if len(status.ErrorMessage) > 0 {
		panic(status.ErrorMessage)
	}

	data  := schema.Data
	count := reflect.ValueOf(data).Len()

	if count == 0 {
		log.Fatal("Malformed response.")
	}

	items := make([]results.Quotes, count)

	i := 0

	// Store applicable items.
	for _, v1 := range data {
		items[i].Name        = v1.Name
		items[i].Symbol      = v1.Symbol
		items[i].TotalSupply = int64(v1.TotalSupply)

		for _, v2 := range v1.Quote {
			items[i].Price            = v2.Price
			items[i].Volume24h        = v2.Volume24h
			items[i].PercentChange1h  = v2.PercentChange1h
			items[i].PercentChange24h = v2.PercentChange24h
			items[i].PercentChange7d  = v2.PercentChange7d
			items[i].MarketCap        = v2.MarketCap
		}

		i++
	}

	return items
}
