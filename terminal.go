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
	"os"
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
	Config    *lib.Config
	Currency  *common.Currency
	Language  *common.Language
	useTicker bool
}

//
// NewTerminal creates a new terminal instance.
//
func NewTerminal(config *lib.Config) *Terminal {
	terminal := &Terminal{}
	terminal.Config   = config
	terminal.Currency = common.NewCurrency(config.Currency())
	terminal.Language = common.NewLanguage(config.Language())
	terminal.initTermui()
	return terminal
}

//
// Initializes termui dashboard.
//
func (terminal *Terminal) initTermui() {
	if err := ui.Init(); err != nil {
		log.Fatal("Failed to initialize termui ", err)
	}

	defer ui.Close()

	// Render terminal widgets.
	if terminal.Config.IsValid() {
		terminal.renderDashboard()
	} else {
		terminal.renderSetup()
	}
}

//
// Initializes keyboard/mouse event handling.
//
func (terminal *Terminal) initEvents(actions common.WidgetAction, events common.WidgetEvent) {
	uiEvents := ui.PollEvents()

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case e := <-uiEvents:
			events(e)

			switch e.ID {

			// Show help menu.
			case "?":
				terminal.renderHelp()

			// Exit the terminal.
			case "<End>":
				terminal.exitTerminal()

			// Reset the terminal.
			case "<Escape>":
				terminal.resetTermui()

			// Show Setup screen.
			case "<Home>":
				terminal.renderSetup()

			// Resize the screen.
			case "<Resize>":
				actions()
			}

		case <-ticker.C:
			if terminal.useTicker {
				actions()
			}
		}
	}
}

//
// Renders dashboard widgets.
//
func (terminal *Terminal) renderDashboard() {
	terminal.useTicker = true

	widget1 := terminal.initClock()
	widget2 := terminal.initQuotes()

	actions := func() {
		widget1.Render()
		widget2.Render()
	}

	events := func(e ui.Event) {
		widget2.Events(e)
	}

	actions()

	terminal.initEvents(actions, events)
}

//
// Renders Help widget.
//
func (terminal *Terminal) renderHelp() {
	terminal.useTicker = false

	terminal.initHelp().Render()
}

//
// Renders Setup widget.
//
func (terminal *Terminal) renderSetup() {
	terminal.useTicker = false

	widget := terminal.initSetup()

	actions := func() {
		widget.Render()
	}

	events := func(e ui.Event) {
		widget.Events(e)
	}

	actions()

	terminal.initEvents(actions, events)
}

//
// Returns an instance of the Clock widget.
//
func (terminal *Terminal) initClock() *widgets.Clock {
	return widgets.NewClock(terminal.Language)
}

//
// Returns an instance of the Help widget.
//
func (terminal *Terminal) initHelp() *widgets.Help {
	return widgets.NewHelp(terminal.Language)
}

//
// Returns an instance of the Quotes widget.
//
func (terminal *Terminal) initQuotes() *widgets.Quotes {
	return widgets.NewQuotes(terminal.Config, terminal.Currency, terminal.Language)
}

//
// Returns an instance of the Setup widget.
//
func (terminal *Terminal) initSetup() *widgets.Setup {
	return widgets.NewSetup(terminal.Config, terminal.Language)
}

//
// Exits the terminal application.
//
func (terminal *Terminal) exitTerminal() {
	ui.Close()
	os.Exit(0)
}

//
// Resets the termui dashboard.
//
func (terminal *Terminal) resetTermui() {
	ui.Close()

	terminal.initTermui()
}
