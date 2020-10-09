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
	"fmt"

	"github.com/dustin/go-humanize"
	pad "github.com/willf/pad/utf8"
)

//
// FormatCommas returns number in a comma delimited format.
//
func FormatCommas(v int64) string {
	return humanize.Commaf(float64(v))
}

//
// PadLft returns left-padded string for a given length.
//
func PadLft(v interface{}, length int) string {
	return pad.Left(fmt.Sprint(v), length, " ")
}

//
// PadRgt returns right-padded string for a given length.
//
func PadRgt(v interface{}, length int) string {
	return pad.Right(fmt.Sprint(v), length, " ")
}
