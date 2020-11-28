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
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/nuxy/go-crypto-market-ui/lib/common"
)

// Widget properties.
var errorRect = common.Widget{
	Left:   58,
	Top:    4,
	Right:  119,
	Bottom: 9,
}

//
// Error declared data types.
//
type Error struct {
	Language *common.Language
	content  string
}

//
// NewError creates a new widget instance.
//
func NewError(language *common.Language) *Error {
	widget := &Error{}
	widget.Language = language
	return widget
}

//
// Render the widget.
//
func (widget *Error) Render() {
	obj := widgets.NewParagraph()
	obj.Text          = widget.content
	obj.BorderStyle   = common.WidgetBorderStyle()
	obj.TitleStyle    = common.WidgetTitleStyle()
	obj.PaddingLeft   = 2
	obj.PaddingTop    = 1
	obj.PaddingRight  = 1
	obj.PaddingBottom = 1

	obj.SetRect(
		errorRect.Left,
		errorRect.Top,
		errorRect.Right,
		errorRect.Bottom,
	)

	ui.Render(obj)
}

//
// Messages defines the instance content.
//
func (widget *Error) Messages(v []string) *Error {
	for i, text := range v {
		v[i] = common.PadRgt("⚠", 3) + text
	}

	widget.content = strings.Join(v, "\n")

	return widget
}
