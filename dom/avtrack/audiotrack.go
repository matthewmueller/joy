package avtrack

import "github.com/matthewmueller/golly/js"

// AudioTrack struct
// js:"AudioTrack,omit"
type AudioTrack struct {
}

// Enabled prop
func (*AudioTrack) Enabled() (enabled bool) {
	js.Rewrite("$<.enabled")
	return enabled
}

// Enabled prop
func (*AudioTrack) SetEnabled(enabled bool) {
	js.Rewrite("$<.enabled = $1", enabled)
}

// ID prop
func (*AudioTrack) ID() (id string) {
	js.Rewrite("$<.id")
	return id
}

// Kind prop
func (*AudioTrack) Kind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

// Kind prop
func (*AudioTrack) SetKind(kind string) {
	js.Rewrite("$<.kind = $1", kind)
}

// Label prop
func (*AudioTrack) Label() (label string) {
	js.Rewrite("$<.label")
	return label
}

// Language prop
func (*AudioTrack) Language() (language string) {
	js.Rewrite("$<.language")
	return language
}

// Language prop
func (*AudioTrack) SetLanguage(language string) {
	js.Rewrite("$<.language = $1", language)
}

// SourceBuffer prop
func (*AudioTrack) SourceBuffer() (sourceBuffer *SourceBuffer) {
	js.Rewrite("$<.sourceBuffer")
	return sourceBuffer
}
