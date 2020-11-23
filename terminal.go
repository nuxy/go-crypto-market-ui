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
	Error     string
	useTicker bool
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
// Initializes terminal dependencies.
//
func (terminal *Terminal) init() {
	terminal.Currency = common.NewCurrency(terminal.Config.Currency())
	terminal.Language = common.NewLanguage(terminal.Config.Language())
	terminal.initTermui()
}

//
// Initializes termui dashboard.
//
func (terminal *Terminal) initTermui() {
	if err := ui.Init(); err != nil {
		log.Fatal("Failed to initialize termui ", err)
	}

	defer ui.Close()

	// Validate required fields.
	if !terminal.Config.IsValid() {
		terminal.Error = terminal.Language.Translate("MissingValues")

		terminal.renderSetup()
	}

	if len(terminal.Config.ServiceName()) == 0 {
		terminal.errors = append(
			terminal.errors,
			terminal.Language.Translate("InvalidService"),
		)
	}

	if !terminal.Currency.IsValid() {
		terminal.Error = terminal.Language.Translate("InvalidCurrency")

		terminal.renderSetup()
	}

	if !terminal.Language.IsValid() {
		terminal.Error = terminal.Language.Translate("InvalidLanguage")

		terminal.renderSetup()
	}

	terminal.Error = ``

	// Render termui widgets.
	terminal.renderDashboard()
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
				terminal.resetTerminal()

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

	widget1 := terminal.initQuotes()
	widget2 := terminal.initHoldings()
	widget3 := terminal.initProfile()
	widget4 := terminal.initRelease()
	widget5 := terminal.initClock()

	actions := func() {
		widget1.Render()

		selected := widget1.Selected()

		widget2.Symbol(selected).Render()
		widget3.Symbol(selected).Render()
		widget4.Render()
		widget5.Render()
	}

	events := func(e ui.Event) {
		widget1.Events(e)
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

	widget1 := terminal.initSetup()
	widget2 := terminal.initError()

	actions := func() {
		widget1.Render()

		if len(terminal.Error) > 0 {
			widget2.Message(terminal.Error).Render()
		}
	}

	events := func(e ui.Event) {
		widget1.Events(e)
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
// Returns an instance of the Error widget.
//
func (terminal *Terminal) initError() *widgets.Error {
	return widgets.NewError(terminal.Language)
}

//
// Returns an instance of the Help widget.
//
func (terminal *Terminal) initHelp() *widgets.Help {
	return widgets.NewHelp(terminal.Language)
}

//
// Returns an instance of the Holdings widget.
//
func (terminal *Terminal) initHoldings() *widgets.Holdings {
	return widgets.NewHoldings(terminal.Config, terminal.Currency, terminal.Language)
}

//
// Returns an instance of the Profile widget.
//
func (terminal *Terminal) initProfile() *widgets.Profile {
	return widgets.NewProfile(terminal.Config, terminal.Language)
}

//
// Returns an instance of the Quotes widget.
//
func (terminal *Terminal) initQuotes() *widgets.Quotes {
	return widgets.NewQuotes(terminal.Config, terminal.Currency, terminal.Language)
}

//
// Returns an instance of the Release widget.
//
func (terminal *Terminal) initRelease() *widgets.Release {
	return widgets.NewRelease()
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
// Resets the terminal application.
//
func (terminal *Terminal) resetTerminal() {
	ui.Close()

	terminal.init()
}
