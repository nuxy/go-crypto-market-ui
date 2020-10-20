//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package metadata

//
// Provides endpoint JSON response schema.
//
// https://pro-api.coinmarketcap.com/v1/cryptocurrency/info
//
type urls struct {
	Website      []string `json:"website"`
	Twitter      []string `json:"twitter"`
	MessageBoard []string `json:"message_board"`
	Chat         []string `json:"chat"`
	Explorer     []string `json:"explorer"`
	Reddit       []string `json:"reddit"`
	TechnicalDoc []string `json:"technical_doc"`
	SourceCode   []string `json:"source_code"`
	Announcement []string `json:"announcement"`
}

type symbol struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Slug        string   `json:"slug"`
	Logo        string   `json:"logo"`
	SubReddit   string   `json:"subreddit"`
	Notice      *string  `json:"notice"`
	Tags        []string `json:"tags"`
	TagNames    []string `json:"tag-names"`
	TagGroups   []string `json:"tag-groups"`
	URLs        urls
	DateAdded   string   `json:"date_added"`
	TwitterUser *string  `json:"twitter_username"`
	IsHidden    int      `json:"is_hidden"`
}

type data symbol

type status struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int    `json:"elapsed"`
	CreditCount  int    `json:"credit_count"`
	Notice       string `json:"notice"`
}

//
// Response declared data types.
//
type Response struct {
	Status status
	Data   map[string]data
}
