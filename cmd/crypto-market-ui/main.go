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
	ui "github.com/nuxy/go-crypto-market-ui"
	"github.com/nuxy/go-crypto-market-ui/lib"
	"github.com/nuxy/go-crypto-market-ui/lib/widgets"
)

// Version Makefile $VERSION
var Version string

func init() {
	widgets.Version = Version
}

//
// Let's get this party started.
//
func main() {

	// TODO: Support multiple services.
	config := lib.NewConfig("CoinMarketCap")

	ui.NewTerminal(config)
}
