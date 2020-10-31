//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package common

import (
	ui "github.com/gizak/termui/v3"
)

//
// Widget declared data types.
//
type Widget struct {
	Title  string
	Left   int
	Top    int
	Right  int
	Bottom int
}

//
// WidgetAction declared function.
//
type WidgetAction func()

//
//WidgetEvent declared function.
//
type WidgetEvent func(e ui.Event)

//
// WidgetActiveStyle returns new termui style instance.
//
func WidgetActiveStyle() ui.Style {
	return ui.NewStyle(ui.ColorYellow)
}

//
// WidgetDefaultStyle returns new termui style instance.
//
func WidgetDefaultStyle() ui.Style {
	return ui.NewStyle(ui.ColorWhite)
}

//
// WidgetBorderStyle returns new termui style instance.
//
func WidgetBorderStyle() ui.Style {
	return ui.NewStyle(ui.ColorWhite)
}

//
// WidgetTitleStyle returns new termui style instance.
//
func WidgetTitleStyle() ui.Style {
	return ui.NewStyle(ui.ColorBlack, ui.ColorWhite, ui.ModifierBold)
}
