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
	Title       string
	Left        int
	Top         int
	Right       int
	Bottom      int
	BorderColor ui.Color
	TextColor   ui.Color
}