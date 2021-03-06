//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package common

//
// ServiceInterface provides runtime methods.
//
type ServiceInterface interface {
	URL   (endpointName string)              string
	Params(endpointName string)              []string
	Parse (endpointName string, body []byte) interface{}
}

//
// EndpointInterface provides runtime methods.
//
type EndpointInterface interface {
	URI()      string
	Params()   []string
	Parse(body []byte) interface{}
}
