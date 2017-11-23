package htmllielement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLLIElement,omit"
type HTMLLIElement struct {
	window.HTMLElement
}

// Type
func (*HTMLLIElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type
func (*HTMLLIElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}

// Value Sets or retrieves the value of a list item.
func (*HTMLLIElement) Value() (value int) {
	js.Rewrite("$<.Value")
	return value
}

// Value Sets or retrieves the value of a list item.
func (*HTMLLIElement) SetValue(value int) {
	js.Rewrite("$<.Value = $1", value)
}
