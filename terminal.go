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
	"github.com/nuxy/go-crypto-market-ui/lib/common"
	"github.com/nuxy/go-crypto-market-ui/lib/widgets"
)

//
// Terminal declared data types.
//
type Terminal struct {
	Config   *lib.Config
	Currency *common.Currency
	Language *common.Language
}

//
// NewTerminal creates a new terminal instance.
//
func NewTerminal(config *lib.Config) *Terminal {
	terminal := &Terminal{}
	terminal.Config = config
	terminal.Currency = common.NewCurrency(config.Currency())
	terminal.Language = common.NewLanguage(config.Language())
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

	ticker := time.NewTicker(time.Second).C

	widget := terminal.renderQuotes()

	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents:
			widget.Events(e)

			switch e.ID {

			// Close the terminal.
			case "<Escape>":
				return

			// Resize the screen.
			case "<Resize>":
				widget.Render()
			}

		case <-ticker:
			widget.Render()
		}
	}
}

//
// Renders Quotes widget.
//
func (terminal *Terminal) renderQuotes() *widgets.Quotes {
	widget := widgets.NewQuotes(terminal.Config, terminal.Currency, terminal.Language)
	widget.Render()
	return widget
}
