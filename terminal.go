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
	"time"

	ui "github.com/gizak/termui/v3"

	"github.com/nuxy/go-crypto-market-ui/lib"
	"github.com/nuxy/go-crypto-market-ui/lib/widgets"
)

//
// Init creates a new terminal instance.
//
func Init(config lib.APIConfig) {
	ticker := time.NewTicker(config.RefreshRate * time.Second).C

	if err := ui.Init(); err != nil {
		log.Fatal("Failed to initialize terminal ", err)
	}

	defer ui.Close()

	// Initialize widgets.
	renderQuotes(config)

	// Listen for events.
	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {

			// Close the terminal.
			case "<Escape>":
				return
			}

		case <-ticker:
			renderQuotes(config)
		}
	}
}

//
// Requests Quotes and renders widget.
//
func renderQuotes(config lib.APIConfig) {
	api := lib.NewAPI(config, "Quotes")

	request := lib.NewRequest(api)

	widget := widgets.NewQuotes(request.Get())
	widget.Render()
}
