package avtrack

import "github.com/matthewmueller/golly/js"

// VideoTrack struct
// js:"VideoTrack,omit"
type VideoTrack struct {
}

// ID prop
func (*VideoTrack) ID() (id string) {
	js.Rewrite("$<.id")
	return id
}

// Kind prop
func (*VideoTrack) Kind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

// Kind prop
func (*VideoTrack) SetKind(kind string) {
	js.Rewrite("$<.kind = $1", kind)
}

// Label prop
func (*VideoTrack) Label() (label string) {
	js.Rewrite("$<.label")
	return label
}

// Language prop
func (*VideoTrack) Language() (language string) {
	js.Rewrite("$<.language")
	return language
}

// Language prop
func (*VideoTrack) SetLanguage(language string) {
	js.Rewrite("$<.language = $1", language)
}

// Selected prop
func (*VideoTrack) Selected() (selected bool) {
	js.Rewrite("$<.selected")
	return selected
}

// Selected prop
func (*VideoTrack) SetSelected(selected bool) {
	js.Rewrite("$<.selected = $1", selected)
}

// SourceBuffer prop
func (*VideoTrack) SourceBuffer() (sourceBuffer *SourceBuffer) {
	js.Rewrite("$<.sourceBuffer")
	return sourceBuffer
}
