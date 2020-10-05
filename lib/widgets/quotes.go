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

	"github.com/nuxy/go-crypto-market-ui/lib/common"
	"github.com/nuxy/go-crypto-market-ui/lib/results"
)

// Widget properties.
const propTitle  string   = "Quotes"
const propLeft   int      = 1
const propRight  int      = 145
const propTop    int      = 1
const propBottom int      = 8
const propText   ui.Color = ui.ColorYellow

// List item padding.
const padCounter          int = 4
const padSymbol           int = 9
const padName             int = 14
const padPrice            int = 12
const padMarketCap        int = 19
const padVolume24h        int = 18
const padTotalSupply      int = 18
const padPercentChange1h  int = 16
const padPercentChange24h int = 16
const padPercentChange7d  int = 10

//
// Quotes declared data types.
//
type Quotes struct {
	Currency *common.Currency
	data     interface{}
}

//
// NewQuotes creates a new widget instance.
//
func NewQuotes(data interface{}, currency string) *Quotes {
	widget := &Quotes{}
	widget.Currency = common.NewCurrency(currency)
	widget.data     = data
	return widget
}

//
// Render the widget.
//
func (widget *Quotes) render() {
	obj := widgets.NewList()
	obj.Title         = propTitle
	obj.Rows          = widget.build()
	obj.TextStyle     = ui.NewStyle(propText)
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
		select {
		case e := <-uiEvents:
			switch e.ID {

			// Scroll item down.
			case "<Down>":
				obj.ScrollDown()

			// Scroll item up.
			case "<Up>":
				obj.ScrollUp()
			}
		}
	}
}

//
// Builds result rows list.
//
func (widget *Quotes) build() []string {

	// TODO: Header text should be localized.
	header := common.PadRgt("#", padCounter) + common.PadRgt("Ticker", padSymbol) + common.PadRgt("Name", padName) + common.PadRgt("Price", padPrice) + common.PadRgt("Market Cap", padMarketCap) + common.PadRgt("24h Volume", padVolume24h) + common.PadRgt("Total Supply", padTotalSupply) + common.PadRgt("% Change (1h)", padPercentChange1h) + common.PadRgt("% Change (24h)", padPercentChange24h) + common.PadRgt("% Change (7d)", padPercentChange7d)

	rows := []string{header}

	for i, v := range widget.data.([]results.Quotes) {
		row := common.PadRgt(i + 1, padCounter) + common.PadRgt(v.Symbol, padSymbol) + common.PadRgt(v.Name, padName) + common.PadRgt(widget.price(v.Price), padPrice) + common.PadRgt(widget.marketCap(v.MarketCap), padMarketCap) + common.PadRgt(widget.volume24h(v.Volume24h), padVolume24h) + common.PadRgt(widget.totalSupply(v.TotalSupply), padTotalSupply) + common.PadRgt(widget.percentChange(v.PercentChange1h), padPercentChange1h) + common.PadRgt(widget.percentChange(v.PercentChange24h), padPercentChange24h) + common.PadRgt(widget.percentChange(v.PercentChange7d), padPercentChange7d)

		rows = append(rows, row)
	}

	return rows
}

//
// Returns termui style instance.
//
func (Quotes) style(color ui.Color) ui.Style {
	return ui.NewStyle(color)
}

//
// Returns post-processed column values.
//
func (widget Quotes) marketCap(v float64) string {
	return widget.Currency.Format(v, 0)
}

func (widget Quotes) price(v float64) string {
	return widget.Currency.Format(v, 2)
}

func (widget Quotes) volume24h(v float64) string {
	return widget.Currency.Format(v, 0)
}

func (Quotes) percentChange(v float64) string {
	return fmt.Sprintf("%.2f", v)
}

func (Quotes) totalSupply(v int64) string {
	return common.FormatCommas(v)
}
