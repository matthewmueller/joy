package htmldivelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLDivElement struct
// js:"HTMLDivElement,omit"
type HTMLDivElement struct {
	window.HTMLElement
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLDivElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLDivElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// NoWrap prop Sets or retrieves whether the browser automatically performs wordwrap.
func (*HTMLDivElement) NoWrap() (noWrap bool) {
	js.Rewrite("$<.noWrap")
	return noWrap
}

// NoWrap prop Sets or retrieves whether the browser automatically performs wordwrap.
func (*HTMLDivElement) SetNoWrap(noWrap bool) {
	js.Rewrite("$<.noWrap = $1", noWrap)
}
