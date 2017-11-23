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

// Value prop
func (*HTMLDataElement) Value() (value string) {
	js.Rewrite("$<.value")
	return value
}

// Value prop
func (*HTMLDataElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}
