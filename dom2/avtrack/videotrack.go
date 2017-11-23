package avtrack

import "github.com/matthewmueller/golly/js"

// VideoTrack struct
// js:"VideoTrack,omit"
type VideoTrack struct {
}

// ID
func (*VideoTrack) ID() (id string) {
	js.Rewrite("$<.ID")
	return id
}

// Kind
func (*VideoTrack) Kind() (kind string) {
	js.Rewrite("$<.Kind")
	return kind
}

// Kind
func (*VideoTrack) SetKind(kind string) {
	js.Rewrite("$<.Kind = $1", kind)
}

// Label
func (*VideoTrack) Label() (label string) {
	js.Rewrite("$<.Label")
	return label
}

// Language
func (*VideoTrack) Language() (language string) {
	js.Rewrite("$<.Language")
	return language
}

// Language
func (*VideoTrack) SetLanguage(language string) {
	js.Rewrite("$<.Language = $1", language)
}

// Selected
func (*VideoTrack) Selected() (selected bool) {
	js.Rewrite("$<.Selected")
	return selected
}

// Selected
func (*VideoTrack) SetSelected(selected bool) {
	js.Rewrite("$<.Selected = $1", selected)
}

// SourceBuffer
func (*VideoTrack) SourceBuffer() (sourceBuffer *SourceBuffer) {
	js.Rewrite("$<.SourceBuffer")
	return sourceBuffer
}
