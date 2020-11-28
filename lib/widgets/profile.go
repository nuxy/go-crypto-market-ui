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
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/nuxy/go-crypto-market-ui/lib"
	"github.com/nuxy/go-crypto-market-ui/lib/common"
	"github.com/nuxy/go-crypto-market-ui/lib/results"
)

// Widget properties.
var profileRect = common.Widget{
	Left:   127,
	Top:    3,
	Right:  176,
	Bottom: 22,
}

//
// Profile declared data types.
//
type Profile struct {
	Request  *lib.Request
	Language *common.Language
	instance *widgets.Paragraph
	content  string
}

//
// NewProfile creates a new widget instance.
//
func NewProfile(config *lib.Config, language *common.Language) *Profile {
	widget := &Profile{}
	widget.Request  = lib.NewRequest(lib.NewAPI(config, "Metadata"))
	widget.Language = language
	return widget
}

//
// Render the widget.
//
func (widget *Profile) Render() {
	var obj = widget.instance

	if widget.instance == nil {
		obj = widgets.NewParagraph()
		obj.Title         = widget.Language.Translate("MarketProfile")
		obj.BorderStyle   = common.WidgetBorderStyle()
		obj.TitleStyle    = common.WidgetTitleStyle()
		obj.PaddingLeft   = 1
		obj.PaddingTop    = 1
		obj.PaddingRight  = 1
		obj.PaddingBottom = 1

		obj.SetRect(
			profileRect.Left,
			profileRect.Top,
			profileRect.Right,
			profileRect.Bottom,
		)

		widget.instance = obj
	}

	obj.Text = widget.content

	ui.Render(obj)
}

//
// Symbol defines the instance content.
//
func (widget *Profile) Symbol(v string) *Profile {
	data := widget.Request.Get()

	if widget.Request.Error != nil {
		panic(widget.Request.Error)
	}

	items := data.([]results.Metadata)

	var item results.Metadata

	for i := 0; i < len(items); i++ {
		item = items[i]

		if v == item.Symbol {
			break
		}
	}

	widget.content = item.Description

	return widget
}
