//
//  go-crypto-market-ui
//  Command-line utility to track cryptocurrencies in realtime.
//
//  Copyright 2020, Marc S. Brooks (https://mbrooks.info)
//  Licensed under the MIT license:
//  http://www.opensource.org/licenses/mit-license.php
//

package widgets

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/nuxy/go-crypto-market-ui/lib/common"
)

// Version Makefile $VERSION
var Version string

// Widget properties.
var releaseRect = common.Widget{
	Left:   1,
	Top:    0,
	Right:  30,
	Bottom: 3,
}

//
// Release declared data types.
//
type Release struct {
	instance *widgets.Paragraph
}

//
// NewRelease creates a new widget instance.
//
func NewRelease() *Release {
	return &Release{}
}

//
// Render the widget.
//
func (widget *Release) Render() {
	var obj = widget.instance

	if widget.instance == nil {
		obj = widgets.NewParagraph()
		obj.Text   = widget.text()
		obj.Border = false

		obj.SetRect(
			releaseRect.Left,
			releaseRect.Top,
			releaseRect.Right,
			releaseRect.Bottom,
		)

		widget.instance = obj
	}

	ui.Render(obj)
}

//
// Returns this project release.
//
func (Release) text() string {
	return fmt.Sprintf("Crypto Market UI (v%s)", Version)
}
