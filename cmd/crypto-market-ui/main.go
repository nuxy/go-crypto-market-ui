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

	// TODO: Support multiple services.
	config := lib.NewConfig("CoinMarketCap")

	api := lib.NewAPI(config.Load(), "Quotes")

	request := lib.NewRequest(api)
	results := request.Get()

	fmt.Println(results)
}
