//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package main

import(
	"fmt"

	"github.com/nuxy/go-crypto-market-ui/lib"
)

//
// Let's get this party started.
//
func main() {
	config := lib.NewConfig("CoinMarketCap")

	fmt.Println(config)
}
