package mimetypearray

import (
	"github.com/matthewmueller/golly/dom2/mimetype"
	"github.com/matthewmueller/golly/js"
)

// js:"MimeTypeArray,omit"
type MimeTypeArray struct {
}

// Item
func (*MimeTypeArray) Item(index uint) (m *mimetype.Plugin) {
	js.Rewrite("$<.Item($1)", index)
	return m
}

// NamedItem
func (*MimeTypeArray) NamedItem(kind string) (m *mimetype.Plugin) {
	js.Rewrite("$<.NamedItem($1)", kind)
	return m
}

// Length
func (*MimeTypeArray) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
