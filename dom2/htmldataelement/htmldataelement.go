package htmldataelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLDataElement struct
// js:"HTMLDataElement,omit"
type HTMLDataElement struct {
	window.HTMLElement
}

// Value
func (*HTMLDataElement) Value() (value string) {
	js.Rewrite("$<.Value")
	return value
}

// Value
func (*HTMLDataElement) SetValue(value string) {
	js.Rewrite("$<.Value = $1", value)
}
