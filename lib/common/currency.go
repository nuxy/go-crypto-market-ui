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
	"log"
	"strings"

	"github.com/leekchan/accounting"
)

//
// Currency declared data types.
//
type Currency struct {
	Code   string
	Symbol string
}

//
// NewCurrency creates a new currency instance.
//
func NewCurrency(code string) *Currency {
	currency := &Currency{}
	currency.Code = strings.ToUpper(code)
	return currency
}

//
// IsValid checks the currency code is supported.
//
func (currency *Currency) IsValid() bool {
	var _, ok = accounting.LocaleInfo[currency.Code]
	return ok
}

//
// Format returns formatted currency value.
//
func (currency *Currency) Format(v float64, precision int) string {
	symbol := currency.load()

	ac := accounting.Accounting{
		Symbol:    symbol,
		Precision: precision,
	}

	return ac.FormatMoney(v)
}

//
// Returns the symbol for a given currency code.
//
func (currency *Currency) load() string {
	var symbol string

	if currency.IsValid() {
		symbol = accounting.LocaleInfo[currency.Code].ComSymbol
	} else {
		log.Fatal("Unsupported currency: ", currency.Code)
	}

	return symbol
}
