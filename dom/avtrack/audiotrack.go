package avtrack

import "github.com/matthewmueller/joy/js"

// AudioTrack struct
// js:"AudioTrack,omit"
type AudioTrack struct {
}

// Enabled prop
// js:"enabled"
func (*AudioTrack) Enabled() (enabled bool) {
	js.Rewrite("$_.enabled")
	return enabled
}

// SetEnabled prop
// js:"enabled"
func (*AudioTrack) SetEnabled(enabled bool) {
	js.Rewrite("$_.enabled = $1", enabled)
}

// ID prop
// js:"id"
func (*AudioTrack) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// Kind prop
// js:"kind"
func (*AudioTrack) Kind() (kind string) {
	js.Rewrite("$_.kind")
	return kind
}

// SetKind prop
// js:"kind"
func (*AudioTrack) SetKind(kind string) {
	js.Rewrite("$_.kind = $1", kind)
}

// Label prop
// js:"label"
func (*AudioTrack) Label() (label string) {
	js.Rewrite("$_.label")
	return label
}

// Language prop
// js:"language"
func (*AudioTrack) Language() (language string) {
	js.Rewrite("$_.language")
	return language
}

// SetLanguage prop
// js:"language"
func (*AudioTrack) SetLanguage(language string) {
	js.Rewrite("$_.language = $1", language)
}

// SourceBuffer prop
// js:"sourceBuffer"
func (*AudioTrack) SourceBuffer() (sourceBuffer *SourceBuffer) {
	js.Rewrite("$_.sourceBuffer")
	return sourceBuffer
}
