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
	prop     common.Widget
	newValue string
	oldValue string
	isActive bool
}

//
// NewField creates a new widget instance.
//
func NewField(prop common.Widget, v string) *Field {
	widget := &Field{}
	widget.prop     = prop
	widget.newValue = v
	widget.oldValue = v
	return widget
}

//
// Render the widget.
//
func (widget *Field) Render() {
	obj := widgets.NewParagraph()
	obj.Title = widget.prop.Title
	obj.Text  = widget.newValue

	if widget.isActive {
		obj.BorderStyle = widget.style(ui.ColorYellow)
	} else {
		obj.BorderStyle = widget.style(ui.ColorWhite)
	}

	if widget.newValue != widget.oldValue {
		obj.TextStyle = widget.style(ui.ColorYellow)
	} else {
		obj.TextStyle = widget.style(ui.ColorWhite)
	}

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
	if widget.isActive {
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
}

//
// Active sets/returns the input state.
//
func (widget *Field) Active(v ...bool) bool {
	if len(v) == 1 {
		widget.isActive = v[0]
		widget.Render()
	}

	return widget.isActive
}

//
// Value returns the field input.
//
func (widget *Field) Value() string {
	return widget.newValue
}

//
// Sets the keyboard event value.
//
func (widget *Field) setValue(v string) {
	r, _ := regexp.Compile(`^[a-zA-Z0-9-_:;~.\/&]{1,100}$`)

	if r.MatchString(v) {
		widget.newValue += v
	}

	widget.Render()
}

//
// Deletes the last character value.
//
func (widget *Field) delValue() {
	v := widget.newValue

	for len(v) > 0 {
		_, size := utf8.DecodeLastRuneInString(v)

		widget.newValue = v[:len(v) - size]
		break
	}

	widget.Render()
}

//
// Returns termui style instance.
//
func (Field) style(color ui.Color) ui.Style {
	return ui.NewStyle(color)
}
