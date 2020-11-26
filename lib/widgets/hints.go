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

// Widget properties.
var hintsRect = common.Widget{
	Left:   15,
	Top:    47,
	Right:  159,
	Bottom: 50,
}

//
// Hints declared data types.
//
type Hints struct {
	Language *common.Language
	content  [][]string
}

//
// NewHints creates a new widget instance.
//
func NewHints(language *common.Language) *Hints {
	widget := &Hints{}
	widget.Language = language
	return widget
}

//
// Render the widget.
//
func (widget *Hints) Render() {
	obj := widgets.NewTable()
	obj.Rows          = widget.content
	obj.ColumnWidths  = []int{35, 35, 35, 35}
	obj.TextAlignment = ui.AlignCenter

	obj.SetRect(
		hintsRect.Left,
		hintsRect.Top,
		hintsRect.Right,
		hintsRect.Bottom,
	)

	ui.Render(obj)
}

//
// Dashboard defines the instance content.
//
func (widget *Hints) Dashboard() *Hints {
	widget.content = [][]string{
		[]string{
			fmt.Sprintf("[?] %s",    widget.Language.Translate("HelpMenu")),
			fmt.Sprintf("[Home] %s", widget.Language.Translate("Configuration")),
			fmt.Sprintf("[Esc] %s",  widget.Language.Translate("ReloadScreen")),
			fmt.Sprintf("[End] %s",  widget.Language.Translate("ExitProgram")),
		},
	}

	return widget
}

//
// Setup defines the instance content.
//
func (widget *Hints) Setup() *Hints {
	widget.content = [][]string{
		[]string{
			fmt.Sprintf("[Enter] %s", widget.Language.Translate("SaveChanges")),
			fmt.Sprintf("[Tab] %s",   widget.Language.Translate("NextValue")),
			fmt.Sprintf("[Esc] %s",   widget.Language.Translate("CloseWindow")),
			fmt.Sprintf("[End] %s",   widget.Language.Translate("ExitProgram")),
		},
	}

	return widget
}
