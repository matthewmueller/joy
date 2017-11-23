package htmltrackelement

import (
	"github.com/matthewmueller/golly/dom2/texttrack"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLTrackElement struct
// js:"HTMLTrackElement,omit"
type HTMLTrackElement struct {
	window.HTMLElement
}

// Default prop
func (*HTMLTrackElement) Default() (def bool) {
	js.Rewrite("$<.default")
	return def
}

// Default prop
func (*HTMLTrackElement) SetDefault(def bool) {
	js.Rewrite("$<.default = $1", def)
}

// Kind prop
func (*HTMLTrackElement) Kind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

// Kind prop
func (*HTMLTrackElement) SetKind(kind string) {
	js.Rewrite("$<.kind = $1", kind)
}

// Label prop
func (*HTMLTrackElement) Label() (label string) {
	js.Rewrite("$<.label")
	return label
}

// Label prop
func (*HTMLTrackElement) SetLabel(label string) {
	js.Rewrite("$<.label = $1", label)
}

// ReadyState prop
func (*HTMLTrackElement) ReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

// Src prop
func (*HTMLTrackElement) Src() (src string) {
	js.Rewrite("$<.src")
	return src
}

// Src prop
func (*HTMLTrackElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

// Srclang prop
func (*HTMLTrackElement) Srclang() (srclang string) {
	js.Rewrite("$<.srclang")
	return srclang
}

// Srclang prop
func (*HTMLTrackElement) SetSrclang(srclang string) {
	js.Rewrite("$<.srclang = $1", srclang)
}

// Track prop
func (*HTMLTrackElement) Track() (track *texttrack.TextTrack) {
	js.Rewrite("$<.track")
	return track
}
