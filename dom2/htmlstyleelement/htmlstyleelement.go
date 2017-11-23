package htmlstyleelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLStyleElement struct
// js:"HTMLStyleElement,omit"
type HTMLStyleElement struct {
	window.HTMLElement
}

// Sheet
func (*HTMLStyleElement) Sheet() (sheet window.StyleSheet) {
	js.Rewrite("$<.Sheet")
	return sheet
}

// Disabled
func (*HTMLStyleElement) Disabled() (disabled bool) {
	js.Rewrite("$<.Disabled")
	return disabled
}

// Disabled
func (*HTMLStyleElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.Disabled = $1", disabled)
}

// Media Sets or retrieves the media type.
func (*HTMLStyleElement) Media() (media string) {
	js.Rewrite("$<.Media")
	return media
}

// Media Sets or retrieves the media type.
func (*HTMLStyleElement) SetMedia(media string) {
	js.Rewrite("$<.Media = $1", media)
}

// Type Retrieves the CSS language in which the style sheet is written.
func (*HTMLStyleElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type Retrieves the CSS language in which the style sheet is written.
func (*HTMLStyleElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}
