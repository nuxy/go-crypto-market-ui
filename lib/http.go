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
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

//
// Request declared data types.
//
type Request struct {
	API       *API
	Error     interface{}
	cacheBody []byte
	cacheTime int64
}

//
// NewRequest creates a new request instance.
//
func NewRequest(api *API) *Request {
	request := &Request{}
	request.API = api
	return request
}

//
// Get data from a remote resource or cached response.
//
func (request *Request) Get() interface{} {
	defer func() {
		request.Error = recover()
	}()

	var body []byte

	nextTime := request.cacheTime + request.API.Config.RefreshRate()

	if time.Now().Unix() > nextTime {
		body = request.fetchData()
	} else {
		body = request.cacheBody
	}

	return request.API.Parse(body)
}

//
// Fetch data from remote resource and cache response.
//
func (request *Request) fetchData() []byte {
	URL := request.API.URL()

	validateURL(URL)

	resp, err := http.Get(URL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(io.Reader(resp.Body))

	if err != nil {
		log.Fatal(err)
	}

	request.cacheTime = time.Now().Unix()
	request.cacheBody = body

	return body
}

//
// Checks for a valid URL structure, fail otherwise.
//
func validateURL(rawURL string) {
	_, err := url.ParseRequestURI(rawURL)

	if err != nil {
		log.Fatal("Invalid URL: ", err)
	}
}
