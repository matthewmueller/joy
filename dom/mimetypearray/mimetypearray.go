package mimetypearray

import (
	"github.com/matthewmueller/joy/dom/mimetype"
	"github.com/matthewmueller/joy/macro"
)

// MimeTypeArray struct
// js:"MimeTypeArray,omit"
type MimeTypeArray struct {
}

// Item fn
// js:"item"
func (*MimeTypeArray) Item(index uint) (m *mimetype.Plugin) {
	macro.Rewrite("$_.item($1)", index)
	return m
}

// NamedItem fn
// js:"namedItem"
func (*MimeTypeArray) NamedItem(kind string) (m *mimetype.Plugin) {
	macro.Rewrite("$_.namedItem($1)", kind)
	return m
}

// Length prop
// js:"length"
func (*MimeTypeArray) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
