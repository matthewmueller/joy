package htmldivelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLDivElement,omit"
type HTMLDivElement struct {
	window.HTMLElement
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLDivElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLDivElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// NoWrap Sets or retrieves whether the browser automatically performs wordwrap.
func (*HTMLDivElement) NoWrap() (noWrap bool) {
	js.Rewrite("$<.NoWrap")
	return noWrap
}

// NoWrap Sets or retrieves whether the browser automatically performs wordwrap.
func (*HTMLDivElement) SetNoWrap(noWrap bool) {
	js.Rewrite("$<.NoWrap = $1", noWrap)
}
