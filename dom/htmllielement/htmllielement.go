package htmllielement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLLIElement struct
// js:"HTMLLIElement,omit"
type HTMLLIElement struct {
	window.HTMLElement
}

// Type prop
func (*HTMLLIElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop
func (*HTMLLIElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

// Value prop Sets or retrieves the value of a list item.
func (*HTMLLIElement) Value() (value int) {
	js.Rewrite("$<.value")
	return value
}

// Value prop Sets or retrieves the value of a list item.
func (*HTMLLIElement) SetValue(value int) {
	js.Rewrite("$<.value = $1", value)
}
