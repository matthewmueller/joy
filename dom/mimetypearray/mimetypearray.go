package mimetypearray

import (
	"github.com/matthewmueller/golly/dom2/mimetype"
	"github.com/matthewmueller/golly/js"
)

// MimeTypeArray struct
// js:"MimeTypeArray,omit"
type MimeTypeArray struct {
}

// Item fn
func (*MimeTypeArray) Item(index uint) (m *mimetype.Plugin) {
	js.Rewrite("$<.item($1)", index)
	return m
}

// NamedItem fn
func (*MimeTypeArray) NamedItem(kind string) (m *mimetype.Plugin) {
	js.Rewrite("$<.namedItem($1)", kind)
	return m
}

// Length prop
func (*MimeTypeArray) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
