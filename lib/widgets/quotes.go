//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package widgets

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/nuxy/go-crypto-market-ui/lib/results"
)

// Widget properties.
const propTitle  string = "Quotes"
const propLeft   int    = 1
const propRight  int    = 145
const propTop    int    = 1
const propBottom int    = 8

//
// Widget declared data types.
//
type Widget struct {
	Data interface{}
}

//
// NewWidget creates a new widget instance.
//
func NewWidget(data interface{}) *Widget {
	widget := &Widget{}
	widget.Data = data
	widget.render()
	return widget
}

//
// Render the terminal widget.
//
func (widget *Widget) render() {
	obj := widgets.NewList()
	obj.Title         = propTitle
	obj.Rows          = widget.build()
	obj.TextStyle     = ui.NewStyle(ui.ColorYellow)
	obj.PaddingLeft   = 1
	obj.PaddingTop    = 1
	obj.PaddingRight  = 1
	obj.PaddingBottom = 1

	obj.SetRect(
		propLeft,
		propTop,
		propRight,
		propBottom
	)

	ui.Render(obj)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "<Escape>":
			return
		}

		ui.Render(obj)
	}
}

//
// Builds result rows list.
//
func (widget Widget) build() []string {
	header := counter("#") + symbol("Ticker") + name("Name") + price("Price") + marketCap("Market Cap") + volume24h("24h Volume") + totalSupply("Total Supply") + percentChange("% Change (1h)") + percentChange("% Change (24h)") + percentChange("% Change (1d)")

	rows := []string{header}

	for i, v := range widget.Data.([]results.Quotes) {
		row := counter(i + 1) + symbol(v.Symbol) + name(v.Name) + price(v.Price) + marketCap(v.MarketCap) + volume24h(v.Volume24h) + totalSupply(v.TotalSupply) + percentChange(v.PercentChange1h) + percentChange(v.PercentChange24h) + percentChange(v.PercentChange7d)

		rows = append(rows, row)
	}

	return rows
}

//
// Returns formatted column values.
//
func counter(v interface{}) string {
	return fmt.Sprintf("%-5v", v)
}

func symbol(v interface{}) string {
	return fmt.Sprintf("%-10v", v)
}

func name(v interface{}) string {
	return fmt.Sprintf("%-12v", v)
}

func price(v interface{}) string {
	return fmt.Sprintf("%-13v", v)
}

func marketCap(v interface{}) string {
	return fmt.Sprintf("%-16v", v)
}

func volume24h(v interface{}) string {
	return fmt.Sprintf("%-16v", v)
}

func percentChange(v interface{}) string {
	return fmt.Sprintf("%-16v", v)
}

func totalSupply(v interface{}) string {
	return fmt.Sprintf("%-16v", v)
}
