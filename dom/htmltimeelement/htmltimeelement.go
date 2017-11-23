package htmltimeelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLTimeElement struct
// js:"HTMLTimeElement,omit"
type HTMLTimeElement struct {
	window.HTMLElement
}

// DateTime prop
func (*HTMLTimeElement) DateTime() (dateTime string) {
	js.Rewrite("$<.dateTime")
	return dateTime
}

// DateTime prop
func (*HTMLTimeElement) SetDateTime(dateTime string) {
	js.Rewrite("$<.dateTime = $1", dateTime)
}
