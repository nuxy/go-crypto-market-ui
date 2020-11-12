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

	ui "github.com/gizak/termui/v3"

	"github.com/nuxy/go-crypto-market-ui/lib"
	"github.com/nuxy/go-crypto-market-ui/lib/common"
	"github.com/nuxy/go-crypto-market-ui/lib/results"
	"github.com/nuxy/go-crypto-market-ui/lib/termui"
)

// Widget properties.
var holdingsRect = common.Widget{
	Left:   1,
	Top:    3,
	Right:  124,
	Bottom: 22,
}

//
// Holdings declared data types.
//
type Holdings struct {
	Request  *lib.Request
	Config   *lib.Config
	Currency *common.Currency
	Language *common.Language
	instance *termui.BarChart
	labels   []string
	values   []float64
	total    float64
}

//
// NewHoldings creates a new widget instance.
//
func NewHoldings(config *lib.Config, currency *common.Currency, language *common.Language) *Holdings {
	widget := &Holdings{}
	widget.Request  = lib.NewRequest(lib.NewAPI(config, "Quotes"))
	widget.Config   = config
	widget.Currency = currency
	widget.Language = language
	return widget
}

//
// Render the widget.
//
func (widget *Holdings) Render() {
	var obj = widget.instance

	if widget.instance == nil {
		obj = termui.NewBarChart()
		obj.NumFormatter  = widget.numFormatter
		obj.BorderStyle   = common.WidgetBorderStyle()
		obj.TitleStyle    = common.WidgetTitleStyle()
		obj.PaddingLeft   = 1
		obj.PaddingTop    = 2
		obj.PaddingRight  = 1
		obj.PaddingBottom = 0
		obj.BarWidth      = 5

		obj.SetRect(
			holdingsRect.Left,
			holdingsRect.Top,
			holdingsRect.Right,
			holdingsRect.Bottom,
		)

		widget.instance = obj
	}

	obj.Title = fmt.Sprintf(
		"%s — %s",
		widget.Language.Translate("MarketHoldings"),
		widget.Currency.Format(widget.total, 2),
	)

	obj.Labels = widget.labels
	obj.Data   = widget.values

	ui.Render(obj)
}

//
// Symbol defines the instance data.
//
func (widget *Holdings) Symbol(v string) *Holdings {
	symbols := widget.Config.Symbols()
	count   := len(symbols)

	labels := make([]string,  0, count)
	values := make([]float64, 0, count)

	var total float64

	for k := range symbols {
		labels = append(labels, k)

		// Calculate portfolio assets total.
		for i := range symbols[k] {
			total = widget.getPrice(k) * symbols[k][i].Units + total
		}
	}

	sort.Strings(labels)

	labels = shiftToValue(labels, v)

	for _, k := range labels {

		// Calculate percentage of each asset.
		for i := range symbols[k] {
			perc := widget.getPrice(k) * symbols[k][i].Units / total

			values = append(values, math.Round(perc * 100))
		}
	}

	if total != 0 {
		widget.labels = labels
		widget.values = values
		widget.total  = total
	}

	return widget
}

//
// Returns the price for a given symbol.
//
func (widget *Holdings) getPrice(symbol string) float64 {
	items := widget.Request.Get().([]results.Quotes)

	var price float64

	for _, v := range items {
		if v.Symbol == symbol {
			price = v.Price
			break
		}
	}

	return price
}

//
// Returns an ordered array starting with the given value.
//
func shiftToValue(slice []string, v string) []string {
	firstItem, otherItems := slice[0], slice[1:]

	if firstItem == v {
		return slice
	}

	otherItems = append(otherItems, firstItem)

	return shiftToValue(otherItems, v)
}

//
// Returns chart preformatted number data.
//
func (Holdings) numFormatter(v float64) string {
	return fmt.Sprintf("%d%%", int(math.RoundToEven(v)))
}
