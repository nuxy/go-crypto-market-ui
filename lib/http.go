//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package lib

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//
// Request declared data types.
//
type Request struct {
	api    *API
	Error  interface{}
}

//
// NewRequest creates a new request instance.
//
func NewRequest(api *API) *Request {
	request := &Request{}
	request.api = api
	return request
}

//
// Get fetches data from a remote resource.
//
func (request *Request) Get() interface{} {
	defer func() {
		request.Error = recover()
	}()

	url := request.api.URL()

	request.validateURL(url)

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return request.api.Parse(body)
}

//
// Checks for a valid URL structure, fail otherwise.
//
func (request *Request) validateURL(rawURL string) {
	_, err := url.ParseRequestURI(rawURL)

	if err != nil {
		log.Fatal(err)
	}
}
