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
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/nuxy/go-crypto-market-ui/lib/common"
)

// Widget properties.
var clockRect = common.Widget{
	Left:   145,
	Top:    0,
	Right:  178,
	Bottom: 3,
}

//
// Clock declared data types.
//
type Clock struct {
	Language *common.Language
	instance *widgets.Paragraph
}

//
// NewClock creates a new widget instance.
//
func NewClock(language *common.Language) *Clock {
	widget := &Clock{}
	widget.Language = language
	return widget
}

//
// Render the widget.
//
func (widget *Clock) Render() {
	var obj = widget.instance

	if widget.instance == nil {
		obj = widgets.NewParagraph()
		obj.Border = false

		obj.SetRect(
			clockRect.Left,
			clockRect.Top,
			clockRect.Right,
			clockRect.Bottom,
		)

		widget.instance = obj
	}

	obj.Text = widget.timeStamp()

	ui.Render(obj)
}

//
// Returns the current timestamp.
//
func (Clock) timeStamp() string {
	return time.Now().Format("January 2, 2006 🕑 15:04:05")
}
