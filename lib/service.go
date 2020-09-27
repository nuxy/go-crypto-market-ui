//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package lib

//
// ServiceConfig declared data types.
//
type ServiceConfig struct {
	Name    string   `json:"name"`
	APIKey  string   `json:"apiKey"`
	URL     string   `json:"url"`
	Symbols []string `json:"symbols"`
}
