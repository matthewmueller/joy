package htmltimeelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLTimeElement struct
// js:"HTMLTimeElement,omit"
type HTMLTimeElement struct {
	window.HTMLElement
}

// DateTime
func (*HTMLTimeElement) DateTime() (dateTime string) {
	js.Rewrite("$<.DateTime")
	return dateTime
}

// DateTime
func (*HTMLTimeElement) SetDateTime(dateTime string) {
	js.Rewrite("$<.DateTime = $1", dateTime)
}
