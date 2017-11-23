package htmlstyleelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLStyleElement struct
// js:"HTMLStyleElement,omit"
type HTMLStyleElement struct {
	window.HTMLElement
}

// Sheet prop
func (*HTMLStyleElement) Sheet() (sheet window.StyleSheet) {
	js.Rewrite("$<.sheet")
	return sheet
}

// Disabled prop
func (*HTMLStyleElement) Disabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

// Disabled prop
func (*HTMLStyleElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

// Media prop Sets or retrieves the media type.
func (*HTMLStyleElement) Media() (media string) {
	js.Rewrite("$<.media")
	return media
}

// Media prop Sets or retrieves the media type.
func (*HTMLStyleElement) SetMedia(media string) {
	js.Rewrite("$<.media = $1", media)
}

// Type prop Retrieves the CSS language in which the style sheet is written.
func (*HTMLStyleElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop Retrieves the CSS language in which the style sheet is written.
func (*HTMLStyleElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
