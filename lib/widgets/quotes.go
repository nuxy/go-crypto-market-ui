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

	"github.com/nuxy/go-crypto-market-ui/lib"
	"github.com/nuxy/go-crypto-market-ui/lib/common"
	"github.com/nuxy/go-crypto-market-ui/lib/results"
)

// Widget properties.
var quotesRect = common.Widget{
	Left:   1,
	Top:    8,
	Right:  145,
	Bottom: 1,
}

// List item padding.
var quotesPad = map[string]int{
	"Counter":          4,
	"Symbol":           9,
	"Name":             14,
	"Price":            12,
	"MarketCap":        19,
	"Volume24h":        18,
	"TotalSupply":      18,
	"PercentChange1h":  16,
	"PercentChange24h": 16,
	"PercentChange7d":  10,
}

//
// Quotes declared data types.
//
type Quotes struct {
	Request  *lib.Request
	Currency *common.Currency
	Language *common.Language
	instance *widgets.List
}

//
// NewQuotes creates a new widget instance.
//
func NewQuotes(config *lib.Config, currency *common.Currency, language *common.Language) *Quotes {
	widget := &Quotes{}
	widget.Request  = lib.NewRequest(lib.NewAPI(config, "Quotes"))
	widget.Currency = currency
	widget.Language = language
	return widget
}

//
// Render the widget.
//
func (widget *Quotes) Render() {
	var obj = widget.instance

	if widget.instance == nil {
		obj = widgets.NewList()
		obj.BorderStyle   = common.WidgetBorderStyle()
		obj.TextStyle     = common.WidgetActiveStyle()
		obj.TitleStyle    = common.WidgetTitleStyle()
		obj.PaddingLeft   = 1
		obj.PaddingTop    = 1
		obj.PaddingRight  = 1
		obj.PaddingBottom = 1

		obj.SetRect(
			quotesRect.Left,
			quotesRect.Top,
			quotesRect.Right,
			quotesRect.Bottom,
		)

		widget.instance = obj
	}

	obj.Title = widget.header()
	obj.Rows  = widget.rows()

	ui.Render(obj)
}

//
// Events propagates keyboard actions.
//
func (widget *Quotes) Events(e ui.Event) {
	switch e.ID {

	// Scroll item down.
	case "<Down>":
		widget.instance.ScrollDown()

	// Scroll item up.
	case "<Up>":
		widget.instance.ScrollUp()
	}
}

//
// Returns results header.
//
func (widget *Quotes) header() string {
	return common.PadRgt("#", quotesPad["Counter"]) + common.PadRgt(widget.Language.Translate("Symbol"), quotesPad["Symbol"]) + common.PadRgt(widget.Language.Translate("Name"), quotesPad["Name"]) + common.PadRgt(widget.Language.Translate("Price"), quotesPad["Price"]) + common.PadRgt(widget.Language.Translate("MarketCap"), quotesPad["MarketCap"]) + common.PadRgt(widget.Language.Translate("Volume24h"), quotesPad["Volume24h"]) + common.PadRgt(widget.Language.Translate("TotalSupply"), quotesPad["TotalSupply"]) + common.PadRgt(widget.Language.Translate("PercentChange1h"), quotesPad["PercentChange1h"]) + common.PadRgt(widget.Language.Translate("PercentChange24h"), quotesPad["PercentChange24h"]) + common.PadRgt(widget.Language.Translate("PercentChange7d"), quotesPad["PercentChange7d"])
}

//
// Returns results rows.
//
func (widget *Quotes) rows() []string {
	var rows []string

	for i, v := range widget.Request.Get().([]results.Quotes) {
		row := common.PadRgt(i + 1, quotesPad["Counter"]) + common.PadRgt(v.Symbol, quotesPad["Symbol"]) + common.PadRgt(v.Name, quotesPad["Name"]) + common.PadRgt(widget.price(v.Price), quotesPad["Price"]) + common.PadRgt(widget.marketCap(v.MarketCap), quotesPad["MarketCap"]) + common.PadRgt(widget.volume24h(v.Volume24h), quotesPad["Volume24h"]) + common.PadRgt(widget.totalSupply(v.TotalSupply), quotesPad["TotalSupply"]) + common.PadRgt(widget.percentChange(v.PercentChange1h), quotesPad["PercentChange1h"]) + common.PadRgt(widget.percentChange(v.PercentChange24h), quotesPad["PercentChange24h"]) + common.PadRgt(widget.percentChange(v.PercentChange7d), quotesPad["PercentChange7d"])

		rows = append(rows, row)
	}

	return rows
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
