package avtrack

import "github.com/matthewmueller/joy/macro"

// VideoTrack struct
// js:"VideoTrack,omit"
type VideoTrack struct {
}

// ID prop
// js:"id"
func (*VideoTrack) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// Kind prop
// js:"kind"
func (*VideoTrack) Kind() (kind string) {
	macro.Rewrite("$_.kind")
	return kind
}

// SetKind prop
// js:"kind"
func (*VideoTrack) SetKind(kind string) {
	macro.Rewrite("$_.kind = $1", kind)
}

// Label prop
// js:"label"
func (*VideoTrack) Label() (label string) {
	macro.Rewrite("$_.label")
	return label
}

// Language prop
// js:"language"
func (*VideoTrack) Language() (language string) {
	macro.Rewrite("$_.language")
	return language
}

// SetLanguage prop
// js:"language"
func (*VideoTrack) SetLanguage(language string) {
	macro.Rewrite("$_.language = $1", language)
}

// Selected prop
// js:"selected"
func (*VideoTrack) Selected() (selected bool) {
	macro.Rewrite("$_.selected")
	return selected
}

// SetSelected prop
// js:"selected"
func (*VideoTrack) SetSelected(selected bool) {
	macro.Rewrite("$_.selected = $1", selected)
}

// SourceBuffer prop
// js:"sourceBuffer"
func (*VideoTrack) SourceBuffer() (sourceBuffer *SourceBuffer) {
	macro.Rewrite("$_.sourceBuffer")
	return sourceBuffer
}
