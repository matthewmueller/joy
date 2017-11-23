package htmltrackelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLTrackElement,omit"
type HTMLTrackElement struct {
	window.HTMLElement
}

// Default
func (*HTMLTrackElement) Default() (def bool) {
	js.Rewrite("$<.Default")
	return def
}

// Default
func (*HTMLTrackElement) SetDefault(def bool) {
	js.Rewrite("$<.Default = $1", def)
}

// Kind
func (*HTMLTrackElement) Kind() (kind string) {
	js.Rewrite("$<.Kind")
	return kind
}

// Kind
func (*HTMLTrackElement) SetKind(kind string) {
	js.Rewrite("$<.Kind = $1", kind)
}

// Label
func (*HTMLTrackElement) Label() (label string) {
	js.Rewrite("$<.Label")
	return label
}

// Label
func (*HTMLTrackElement) SetLabel(label string) {
	js.Rewrite("$<.Label = $1", label)
}

// ReadyState
func (*HTMLTrackElement) ReadyState() (readyState uint8) {
	js.Rewrite("$<.ReadyState")
	return readyState
}

// Src
func (*HTMLTrackElement) Src() (src string) {
	js.Rewrite("$<.Src")
	return src
}

// Src
func (*HTMLTrackElement) SetSrc(src string) {
	js.Rewrite("$<.Src = $1", src)
}

// Srclang
func (*HTMLTrackElement) Srclang() (srclang string) {
	js.Rewrite("$<.Srclang")
	return srclang
}

// Srclang
func (*HTMLTrackElement) SetSrclang(srclang string) {
	js.Rewrite("$<.Srclang = $1", srclang)
}

// Track
func (*HTMLTrackElement) Track() (track *texttrack.TextTrack) {
	js.Rewrite("$<.Track")
	return track
}
