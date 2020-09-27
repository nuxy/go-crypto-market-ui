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
)

//
// Request declared data types.
//
type Request struct {
	service *Service
	Error   interface{}
}

//
// NewRequest creates a new request instance.
//
func NewRequest(service *Service) *Request {
	request := &Request{}
	request.service = service
	return request
}

//
// Get fetches data from a remote resource.
//
func (request *Request) Get() (self *Request) {
	self = request

	defer func() {
		request.Error = recover()
	}()

	resp, err := http.Get(request.service.URL())

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)

	return request
}