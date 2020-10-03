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
// Terminal declared data types.
//
type Terminal struct {
	Config lib.APIConfig
}

//
// NewTerminal creates a new terminal instance.
//
func NewTerminal(config lib.APIConfig) *Terminal {
	terminal := &Terminal{}
	terminal.Config = config
	terminal.initWidgets()
	terminal.initEvents()
	return terminal
}

//
// Initialize terminal widgets.
//
func (terminal *Terminal) initWidgets() {
	if err := ui.Init(); err != nil {
		log.Fatal("Failed to initialize terminal ", err)
	}

	terminal.renderQuotes()
}

//
// Initialize terminal events.
//
func (terminal *Terminal) initEvents() {
	defer ui.Close()

	ticker := time.NewTicker(terminal.Config.RefreshRate * time.Second).C

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
			terminal.renderQuotes()
		}
	}
}

//
// Requests Quotes and renders widget.
//
func (terminal *Terminal) renderQuotes() {
	api := lib.NewAPI(terminal.Config, "Quotes")

	request := lib.NewRequest(api)
	results := request.Get()

	widget := widgets.NewQuotes(results, terminal.Config.Currency)
	widget.Render()
}
