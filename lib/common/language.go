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
	"io/ioutil"
	"log"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/markbates/pkger"

	lang "golang.org/x/text/language"
	yaml "gopkg.in/yaml.v2"
)

//
// Language declared data types.
//
type Language struct {
	Code     string
	instance *i18n.Localizer
}

//
// NewLanguage creates a new language instance.
//
func NewLanguage(locale string) *Language {
	language := &Language{}
	language.Code = strings.ToLower(locale)
	language.load()
	return language
}

//
// IsValid checks the language code is supported.
//
func (language *Language) IsValid() bool {
	var _, err = lang.Parse(language.Code)
	return err == nil
}

//
// Translate returns translated text value.
//
func (language *Language) Translate(key string) string {
	text, err := language.instance.LocalizeMessage(
		&i18n.Message{
			ID: key,
		},
	)

	if err != nil {
		log.Fatal("Missing translation: ", key)
	}

	return text
}

//
// Loads supported language files.
//
func (language *Language) load() {
	code := "en"

	if language.IsValid() {
		code = language.Code
	}

	path := "/locales/" + code + ".yaml"

	f, err := pkger.Open(path)

	if err != nil {
		log.Fatal("Failed to open: ", path)
	}

	defer f.Close()

	bytes, _ := ioutil.ReadAll(f)

	bundle := i18n.NewBundle(lang.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.ParseMessageFileBytes(bytes, path)

	language.instance = i18n.NewLocalizer(bundle, language.Code)
}
