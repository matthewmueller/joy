package svgstyleelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGStyleElement struct
// js:"SVGStyleElement,omit"
type SVGStyleElement struct {
	window.SVGElement
}

// Disabled prop
func (*SVGStyleElement) Disabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

// Disabled prop
func (*SVGStyleElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

// Media prop
func (*SVGStyleElement) Media() (media string) {
	js.Rewrite("$<.media")
	return media
}

// Media prop
func (*SVGStyleElement) SetMedia(media string) {
	js.Rewrite("$<.media = $1", media)
}

// Title prop
func (*SVGStyleElement) Title() (title string) {
	js.Rewrite("$<.title")
	return title
}

// Title prop
func (*SVGStyleElement) SetTitle(title string) {
	js.Rewrite("$<.title = $1", title)
}

// Type prop
func (*SVGStyleElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop
func (*SVGStyleElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
