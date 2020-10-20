//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package results

//
// Metadata declared data types.
//
type Metadata struct {
	Name         string
	Symbol       string
	Category     string
	Description  string
	Website      []string
	Twitter      []string
	MessageBoard []string
	Chat         []string
	Explorer     []string
	Reddit       []string
	TechnicalDoc []string
	Announcement []string
}
