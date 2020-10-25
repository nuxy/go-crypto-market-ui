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

	"github.com/nuxy/go-crypto-market-ui/lib/common"
)

// Widget properties.
var helpRect = common.Widget{
	Left:   44,
	Top:    8,
	Right:  101,
	Bottom: 32,
}

//
// Help declared data types.
//
type Help struct {
	Language *common.Language
}

//
// NewHelp creates a new widget instance.
//
func NewHelp(language *common.Language) *Help {
	widget := &Help{}
	widget.Language = language
	return widget
}

//
// Render the widget.
//
func (widget *Help) Render() {
	obj := widgets.NewParagraph()
	obj.Title         = widget.Language.Translate("KeyboardShortcuts")
	obj.Text          = widget.Language.Translate("KeyboardCommands")
	obj.BorderStyle   = common.WidgetBorderStyle()
	obj.TitleStyle    = common.WidgetTitleStyle()
	obj.PaddingLeft   = 1
	obj.PaddingTop    = 1
	obj.PaddingRight  = 1
	obj.PaddingBottom = 1

	obj.SetRect(
		helpRect.Left,
		helpRect.Top,
		helpRect.Right,
		helpRect.Bottom,
	)

	ui.Render(obj)
}
