//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package metadata

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/nuxy/go-crypto-market-ui/lib/results"
)

//
// Parse returns API response body as results type.
//
func Parse(body []byte) interface{} {
	var schema Response

	err := json.Unmarshal(body, &schema)

	if err != nil {
		log.Fatal("Parser error: ", err)
	}

	status := schema.Status

	if len(status.ErrorMessage) > 0 {
		panic(status.ErrorMessage)
	}

	data  := schema.Data
	count := reflect.ValueOf(data).Len()

	if count == 0 {
		panic("Malformed response.")
	}

	items := make([]results.Metadata, count)

	i := 0

	// Store applicable items.
	for _, v1 := range data {
		items[i].Name         = v1.Name
		items[i].Symbol       = v1.Symbol
		items[i].Category     = v1.Category
		items[i].Description  = v1.Description
		items[i].Website      = v1.URLs.Website
		items[i].Twitter      = v1.URLs.Twitter
		items[i].MessageBoard = v1.URLs.MessageBoard
		items[i].Chat         = v1.URLs.Chat
		items[i].Explorer     = v1.URLs.Explorer
		items[i].Reddit       = v1.URLs.Reddit
		items[i].TechnicalDoc = v1.URLs.TechnicalDoc
		items[i].Announcement = v1.URLs.Announcement

		i++
	}

	return items
}
