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

// Align Sets or retrieves a value that indicates the table alignment.
func (*HTMLHeadingElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Sets or retrieves a value that indicates the table alignment.
func (*HTMLHeadingElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}
