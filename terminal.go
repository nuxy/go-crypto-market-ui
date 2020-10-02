//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package cryptomarketui

import (
	"log"

	ui "github.com/gizak/termui/v3"

	"github.com/nuxy/go-crypto-market-ui/lib"
	"github.com/nuxy/go-crypto-market-ui/lib/widgets"
)

//
// Init creates a new terminal instance.
//
func Init(config lib.APIConfig) {
	if err := ui.Init(); err != nil {
		log.Fatal("Failed to initialize terminal ", err)
	}

	defer ui.Close()

	// Fetch the data
	api := lib.NewAPI(config, "Quotes")

	request := lib.NewRequest(api)

	widgets.NewWidget(request.Get())
}
