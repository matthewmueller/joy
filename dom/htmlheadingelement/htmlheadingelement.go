package htmlheadingelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLHeadingElement struct
// js:"HTMLHeadingElement,omit"
type HTMLHeadingElement struct {
	window.HTMLElement
}

// Align prop Sets or retrieves a value that indicates the table alignment.
func (*HTMLHeadingElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Sets or retrieves a value that indicates the table alignment.
func (*HTMLHeadingElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}
