//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package quotes

//
// Provides endpoint JSON response schema.
//
// https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest
//
type currency struct {
	Price            float64 `json:"price"`
	Volume24h        float64 `json:"volume_24h"`
	PercentChange1h  float64 `json:"percent_change_1h"`
	PercentChange24h float64 `json:"percent_change_24h"`
	PercentChange7d  float64 `json:"percent_change_7d"`
	MarketCap        float64 `json:"market_cap"`
	LastUpdated      string  `json:"last_updated"`
}

type quote map[string]currency

type platform struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Symbol    string `json:"symbol"`
	Slug      string `json:"slug"`
	TokenAddr string `json:"token_address"`
}

type symbol struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Symbol      string    `json:"symbol"`
	Slug        string    `json:"slug"`
	MarketPairs int       `json:"num_market_pairs"`
	DateAdded   string    `json:"date_added"`
	Tags        []string  `json:"tags"`
	MaxSupply   *float64  `json:"max_supply"`
	Circulating float64   `json:"circulating_supply"`
	TotalSupply float64   `json:"total_supply"`
	Platform    *platform `json:"platform"`
	IsActive    int       `json:"is_active"`
	IsMarketInc *int      `json:"is_market_cap_included_in_calc"`
	CmcRank     int       `json:"cmc_rank"`
	IsFiat      int       `json:"is_fiat"`
	LastUpdated string    `json:"last_updated"`
	Quote       quote
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
