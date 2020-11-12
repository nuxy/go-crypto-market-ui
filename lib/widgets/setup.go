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

	ui "github.com/gizak/termui/v3"

	"github.com/nuxy/go-crypto-market-ui/lib"
	"github.com/nuxy/go-crypto-market-ui/lib/common"
)

// Widget properties.
var setupRect = common.Widget{
	Left:   67,
	Top:    12,
	Right:  110,
	Bottom: 35,
}

//
// Setup declared data types.
//
type Setup struct {
	Config   *lib.Config
	Language *common.Language
	fields   []*Field
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
	obj.Title       = widget.Language.Translate("Configuration")
	obj.BorderStyle = common.WidgetBorderStyle()
	obj.TitleStyle  = common.WidgetTitleStyle()

	obj.SetRect(
		setupRect.Left,
		setupRect.Top,
		setupRect.Right,
		setupRect.Bottom,
	)

	ui.Render(obj)

	widget.RenderFields()
}

//
// RenderFields creates editable fields.
//
func (widget *Setup) RenderFields() {
	margin := 4

	// Widget properties.
	var fieldProp = common.Widget{
		Left:   setupRect.Left  + margin / 2,
		Top:    setupRect.Top   - margin / 2,
		Right:  setupRect.Right - margin / 2,
		Bottom: setupRect.Top   + 1,
	}

	r := reflect.ValueOf(widget.Config.Options())

	for i := 0; i < r.NumField(); i++ {
		name := r.Type().Field(i).Name

		if name == "Symbols" {
			continue
		}

		value := r.Field(i).Interface()

		fieldProp.Title   = widget.Language.Translate(name)
		fieldProp.Top    += margin
		fieldProp.Bottom += margin

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

	// Save field values.
	case "<Enter>":
		widget.saveFields()

	// Toggle active field.
	case "<Tab>":
		widget.setActive()
	}
}

//
// Save the field config values.
//
func (widget *Setup) saveFields() {
	r := reflect.ValueOf(widget.Config.Options())

	for i := 0; i < r.NumField(); i++ {
		name := r.Type().Field(i).Name

		if name == "Symbols" {
			continue
		}

		// TODO: Support multiple services.
		if name == "ServiceName" {
			continue
		}

		value := widget.fields[i].Value()

		switch name {
		case "ServiceName": widget.Config.ServiceName(value)
		case "APIKey":      widget.Config.APIKey(value)
		case "Currency":    widget.Config.Currency(value)
		case "Language":    widget.Config.Language(value)
		case "RefreshRate": widget.Config.RefreshRate(value)
		}
	}

	widget.Config.Save()
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
