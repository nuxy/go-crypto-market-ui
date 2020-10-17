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
	"reflect"
	"strings"

	ui "github.com/gizak/termui/v3"

	"github.com/nuxy/go-crypto-market-ui/lib"
	"github.com/nuxy/go-crypto-market-ui/lib/common"
)

// Widget properties.
var setupProp = common.Widget{
	Title:       "Setup",
	Left:        50,
	Top:         8,
	Right:       95,
	Bottom:      35,
	BorderColor: ui.ColorWhite,
	TextColor:   ui.ColorYellow,
}

//
// Setup declared data types.
//
type Setup struct {
	Config   *lib.Config
	Language *common.Language
	fields []*Field
}

//
// NewSetup creates a new widget instance.
//
func NewSetup(config *lib.Config, language *common.Language) *Setup {
	widget := &Setup{}
	widget.Config   = config
	widget.Language = language
	return widget
}

//
// Render the widget.
//
func (widget *Setup) Render() {
	obj := ui.NewBlock()
	obj.Title       = setupProp.Title
	obj.BorderStyle = widget.style(setupProp.BorderColor)
	obj.TitleStyle  = widget.style(setupProp.TextColor)

	obj.SetRect(
		setupProp.Left,
		setupProp.Top,
		setupProp.Right,
		setupProp.Bottom,
	)

	ui.Render(obj)

	widget.renderFields()
}

//
// Render editable fields.
//
func (widget *Setup) renderFields() {

	// Widget properties.
	var fieldProp = common.Widget{
		Left:  setupProp.Left  + 2,
		Right: setupProp.Right - 2,
	}

	r := reflect.ValueOf(widget.Config.Options())

	for i := 0; i < r.NumField(); i++ {
		name := r.Type().Field(i).Name

		fieldProp.Title = name 

		value := r.Field(i).Interface()

		switch name {
		case "Name":
			fieldProp.Bottom = 13
			fieldProp.Top    = 10

		case "APIKey":
			fieldProp.Bottom = 17
			fieldProp.Top    = 14

		case "Currency":
			fieldProp.Bottom = 21
			fieldProp.Top    = 18

		case "Language":
			fieldProp.Bottom = 25
			fieldProp.Top    = 22

		case "Symbols":
			fieldProp.Bottom = 29
			fieldProp.Top    = 26

			value = strings.Join(value.([]string), ",")

		case "RefreshRate":
			fieldProp.Bottom = 33
			fieldProp.Top    = 30
		}

		field := NewField(fieldProp, fmt.Sprintf("%v", value))

		widget.fields = append(widget.fields, field)

		field.Render()
	}

	widget.setActive()
}

//
// Events propagates keyboard actions.
//
func (widget *Setup) Events(e ui.Event) {
	for _, field := range widget.fields {
		field.Events(e)
	}

	switch e.ID {

	// Toggle active field.
	case "<Enter>":
		widget.setActive()

	case "<Tab>":
		widget.setActive()
	}
}

//
// Sets the active (focused) field.
//
func (widget *Setup) setActive() {
	var nextActive int

	for i, field := range widget.fields {
		active := field.Active()

		if active {
			field.Active(false)

			nextActive = i + 1
			break
		}
	}

	if nextActive == len(widget.fields) {
		nextActive = 0
	}

	widget.fields[nextActive].Active(true)
}

//
// Returns termui style instance.
//
func (Setup) style(color ui.Color) ui.Style {
	return ui.NewStyle(color)
}
