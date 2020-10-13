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
	Config *lib.Config
}

//
// NewTerminal creates a new terminal instance.
//
func NewTerminal(config *lib.Config) *Terminal {
	terminal := &Terminal{}
	terminal.Config = config
	terminal.init()
	return terminal
}

//
// Initialize terminal widgets.
//
func (terminal *Terminal) init() {
	if err := ui.Init(); err != nil {
		log.Fatal("Failed to initialize terminal ", err)
	}

	defer ui.Close()

	ticker := time.NewTicker(terminal.Config.RefreshRate() * time.Second).C

	w1 := terminal.renderQuotes()

	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents:
			w1.Events(e)

			switch e.ID {

			// Close the terminal.
			case "<Escape>":
				return

			// Resize the screen.
			case "<Resize>":
				w1.Render()
			}

		case <-ticker:
			w1.Render()
		}

		w1.Render()
	}
}

//
// Requests Quotes and renders widget.
//
func (terminal *Terminal) renderQuotes() *widgets.Quotes {
	api := lib.NewAPI(terminal.Config, "Quotes")

	request := lib.NewRequest(api)
	results := request.Get()

	widget := widgets.NewQuotes(results, terminal.Config.Currency)
	widget.Render()
	return widget
}
