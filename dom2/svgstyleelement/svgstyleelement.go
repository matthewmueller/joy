package svgstyleelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGStyleElement struct
// js:"SVGStyleElement,omit"
type SVGStyleElement struct {
	window.SVGElement
}

// Disabled
func (*SVGStyleElement) Disabled() (disabled bool) {
	js.Rewrite("$<.Disabled")
	return disabled
}

// Disabled
func (*SVGStyleElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.Disabled = $1", disabled)
}

// Media
func (*SVGStyleElement) Media() (media string) {
	js.Rewrite("$<.Media")
	return media
}

// Media
func (*SVGStyleElement) SetMedia(media string) {
	js.Rewrite("$<.Media = $1", media)
}

// Title
func (*SVGStyleElement) Title() (title string) {
	js.Rewrite("$<.Title")
	return title
}

// Title
func (*SVGStyleElement) SetTitle(title string) {
	js.Rewrite("$<.Title = $1", title)
}

// Type
func (*SVGStyleElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type
func (*SVGStyleElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}
