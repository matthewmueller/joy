package avtrack

import "github.com/matthewmueller/golly/js"

// AudioTrack struct
// js:"AudioTrack,omit"
type AudioTrack struct {
}

// Enabled
func (*AudioTrack) Enabled() (enabled bool) {
	js.Rewrite("$<.Enabled")
	return enabled
}

// Enabled
func (*AudioTrack) SetEnabled(enabled bool) {
	js.Rewrite("$<.Enabled = $1", enabled)
}

// ID
func (*AudioTrack) ID() (id string) {
	js.Rewrite("$<.ID")
	return id
}

// Kind
func (*AudioTrack) Kind() (kind string) {
	js.Rewrite("$<.Kind")
	return kind
}

// Kind
func (*AudioTrack) SetKind(kind string) {
	js.Rewrite("$<.Kind = $1", kind)
}

// Label
func (*AudioTrack) Label() (label string) {
	js.Rewrite("$<.Label")
	return label
}

// Language
func (*AudioTrack) Language() (language string) {
	js.Rewrite("$<.Language")
	return language
}

// Language
func (*AudioTrack) SetLanguage(language string) {
	js.Rewrite("$<.Language = $1", language)
}

// SourceBuffer
func (*AudioTrack) SourceBuffer() (sourceBuffer *SourceBuffer) {
	js.Rewrite("$<.SourceBuffer")
	return sourceBuffer
}
