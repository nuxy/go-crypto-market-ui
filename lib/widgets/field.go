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
	"regexp"
	"unicode/utf8"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/nuxy/go-crypto-market-ui/lib/common"
)

//
// Field declared data types.
//
type Field struct {
	prop  common.Widget
	value string
}

//
// NewField creates a new widget instance.
//
func NewField(prop common.Widget, value string) *Field {
	widget := &Field{}
	widget.prop  = prop
	widget.value = value
	return widget
}

//
// Render the widget.
//
func (widget *Field) Render() {
	obj := widgets.NewParagraph()
	obj.Title       = widget.prop.Title
	obj.Text        = widget.value
	obj.BorderStyle = widget.style(widget.prop.BorderColor)
	obj.TextStyle   = widget.style(widget.prop.TextColor)

	obj.SetRect(
		widget.prop.Left,
		widget.prop.Top,
		widget.prop.Right,
		widget.prop.Bottom,
	)

	ui.Render(obj)
}

//
// Events propagates keyboard actions.
//
func (widget *Field) Events(e ui.Event) {
	switch e.ID {

	// Delete last character.
	case "<Backspace>":
		widget.delValue()

	case "<Delete>":
		widget.delValue()

	// Everything else.
	default:
		widget.setValue(e.ID)
	}
}

//
// Sets the keyboard event value.
//
func (widget *Field) setValue(v string) {
	r, _ := regexp.Compile(`^[a-z0-9-_:;~.\/&]{1,100}$`)

	if r.MatchString(v) {
		widget.value += v
	}
}

//
// Deletes the last character value.
//
func (widget *Field) delValue() {
	v := widget.value

	for len(v) > 0 {
		_, size := utf8.DecodeLastRuneInString(v)

		widget.value = v[:len(v) - size]
		break
	}
}

//
// Returns termui style instance.
//
func (Field) style(color ui.Color) ui.Style {
	return ui.NewStyle(color)
}
