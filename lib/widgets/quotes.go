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
	"math"
	"sort"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/nuxy/go-crypto-market-ui/lib"
	"github.com/nuxy/go-crypto-market-ui/lib/common"
	"github.com/nuxy/go-crypto-market-ui/lib/results"
)

// Widget properties.
var quotesRect = common.Widget{
	Left:   1,
	Top:    23,
	Right:  176,
	Bottom: 47,
}

//
// Quotes declared data types.
//
type Quotes struct {
	Request  *lib.Request
	Currency *common.Currency
	Language *common.Language
	instance *widgets.List
	selected string
	sortCol  string
	sortDir  string
}

//
// NewQuotes creates a new widget instance.
//
func NewQuotes(config *lib.Config, currency *common.Currency, language *common.Language) *Quotes {
	widget := &Quotes{}
	widget.Request  = lib.NewRequest(lib.NewAPI(config, "Quotes"))
	widget.Currency = currency
	widget.Language = language
	widget.sortCol  = "Symbol"
	return widget
}

//
// Render the widget.
//
func (widget *Quotes) Render() {
	var obj = widget.instance

	if widget.instance == nil {
		obj = widgets.NewList()
		obj.BorderStyle      = common.WidgetBorderStyle()
		obj.SelectedRowStyle = common.WidgetActiveStyle()
		obj.TitleStyle       = common.WidgetTitleStyle()
		obj.PaddingLeft      = 1
		obj.PaddingTop       = 1
		obj.PaddingRight     = 1
		obj.PaddingBottom    = 1

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

	// Sort by column.
	case "1": widget.sortByCol("Symbol")
	case "2": widget.sortByCol("Name")
	case "3": widget.sortByCol("Price")
	case "4": widget.sortByCol("MarketCap")
	case "5": widget.sortByCol("Volume24h")
	case "6": widget.sortByCol("TotalSupply")
	case "7": widget.sortByCol("PercentChange1h")
	case "8": widget.sortByCol("PercentChange24h")
	case "9": widget.sortByCol("PercentChange7d")

	// Scroll item down.
	case "<Down>":
		widget.instance.ScrollDown()

	// Scroll item up.
	case "<Up>":
		widget.instance.ScrollUp()
	}
}

//
// Selected returns the symbol value.
//
func (widget *Quotes) Selected() string {
	return widget.selected
}

//
// Returns API response results.
//
func (widget *Quotes) results() []results.Quotes {
	items := widget.Request.Get().([]results.Quotes)

	// Sort items by field name.
	sort.SliceStable(items, func(i, j int) bool {
		var cmp bool

		switch widget.sortCol {
		case "Symbol":           cmp = widget.cmpStr(items[i].Symbol,           items[j].Symbol)
		case "Name":             cmp = widget.cmpStr(items[i].Name,             items[j].Name)
		case "Price":            cmp = widget.cmpFlt(items[i].Price,            items[j].Price)
		case "MarketCap":        cmp = widget.cmpFlt(items[i].MarketCap,        items[j].MarketCap)
		case "Volume24h":        cmp = widget.cmpFlt(items[i].Volume24h,        items[j].Volume24h)
		case "TotalSupply":      cmp = widget.cmpInt(items[i].TotalSupply,      items[j].TotalSupply)
		case "PercentChange1h":  cmp = widget.cmpFlt(items[i].PercentChange1h,  items[j].PercentChange1h)
		case "PercentChange24h": cmp = widget.cmpFlt(items[i].PercentChange24h, items[j].PercentChange24h)
		case "PercentChange7d":  cmp = widget.cmpFlt(items[i].PercentChange7d,  items[j].PercentChange7d)
		}

		return cmp
	})

	widget.selected = items[widget.instance.SelectedRow].Symbol

	return items
}

//
// Returns results rows.
//
func (widget *Quotes) rows() []string {
	padWidth := widget.padWidth("default")

	rows := []string{}

	for i, v := range widget.results() {
		row := fmt.Sprint(
			common.PadRgt(i + 1,                                    padWidth[0]),
			common.PadRgt(v.Symbol,                                 padWidth[1]),
			common.PadRgt(widget.name(v.Name),                      padWidth[2]),
			common.PadRgt(widget.price(v.Price),                    padWidth[3]),
			common.PadRgt(widget.marketCap(v.MarketCap),            padWidth[4]),
			common.PadRgt(widget.volume24h(v.Volume24h),            padWidth[5]),
			common.PadRgt(widget.totalSupply(v.TotalSupply),        padWidth[6]),
			common.PadRgt(widget.percentChange(v.PercentChange1h),  padWidth[7]),
			common.PadRgt(widget.percentChange(v.PercentChange24h), padWidth[8]),
			common.PadRgt(widget.percentChange(v.PercentChange7d),  padWidth[9]),
		)

		// Highlight daily changes (e.g. loss).
		if math.Signbit(v.PercentChange24h) {
			row = fmt.Sprintf("[%s](fg:red)", row)
		}

		rows = append(rows, row)
	}

	return rows
}

//
// Returns results header.
//
func (widget *Quotes) header() string {
	padWidth := widget.padWidth(widget.Language.Code)

	return fmt.Sprint(
		common.PadRgt("#",                              padWidth[0]),
		common.PadRgt(widget.title("Symbol"),           padWidth[1]),
		common.PadRgt(widget.title("Name"),             padWidth[2]),
		common.PadRgt(widget.title("Price"),            padWidth[3]),
		common.PadRgt(widget.title("MarketCap"),        padWidth[4]),
		common.PadRgt(widget.title("Volume24h"),        padWidth[5]),
		common.PadRgt(widget.title("TotalSupply"),      padWidth[6]),
		common.PadRgt(widget.title("PercentChange1h"),  padWidth[7]),
		common.PadRgt(widget.title("PercentChange24h"), padWidth[8]),
		common.PadRgt(widget.title("PercentChange7d"),  padWidth[9]),
	)
}

//
// Returns a sort direction arrow for a field.
//
func (widget *Quotes) sortArrow(v string) string {
	var char string

	if widget.sortCol == v {
		switch widget.sortDir {
		case "ASC":  char = common.PadRgt("▴", 2)
		case "DESC": char = common.PadRgt("▾", 2)
		}
	}

	return char
}

//
// Toggle the sort direction; define otherwise.
//
func (widget *Quotes) sortByCol(v string) {
	widget.sortCol = v
	widget.sortByDir()
}

//
// Toggle the sort direction; define otherwise.
//
func (widget *Quotes) sortByDir() {
	if widget.sortDir != `` && widget.sortDir == "ASC" {
		widget.sortDir = "DESC"
	} else {
		widget.sortDir = "ASC"
	}
}

//
// Returns a sort comparison operation results.
//
func (widget *Quotes) cmpInt(a, b int64) bool {
	if widget.sortDir == "DESC" {
		return a > b
	}

	return a < b
}

func (widget *Quotes) cmpFlt(a, b float64) bool {
	if widget.sortDir == "DESC" {
		return a > b
	}

	return a < b
}

func (widget *Quotes) cmpStr(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	if widget.sortDir == "DESC" {
		return a > b
	}

	return a < b
}

//
// Returns post-processed column values.
//
func (widget *Quotes) title(v string) string {
	return fmt.Sprint(widget.sortArrow(v), widget.Language.Translate(v))
}

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

func (Quotes) name(v string) string {
	return common.TruncStr(v, 15)
}

func (Quotes) totalSupply(v int64) string {
	return common.FormatCommas(v)
}

//
// Returns column padding width by locale.
//
func (Quotes) padWidth(locale string) [10]int {
	switch locale {

	// Multi-byte characters.
	case "ja":
		return [10]int{6, 6, 15, 11, 21, 13, 18, 14, 14, 13}
	case "ko":
		return [10]int{6, 8, 15, 11, 21, 15, 18, 16, 16, 16}
	case "vi":
		return [10]int{6, 12, 17, 13, 25, 20, 21, 20, 20, 19}
	case "zh":
		return [10]int{6, 8, 15, 11, 23, 15, 18, 16, 16, 16}

	// Single-byte characters.
	default:
		return [10]int{6, 10, 17, 13, 25, 20, 21, 20, 20, 19}
	}
}
