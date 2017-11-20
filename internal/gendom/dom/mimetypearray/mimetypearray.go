package mimetypearray

import (
	"plugin"

	"github.com/matthewmueller/golly/js"
)

type MimeTypeArray struct {
}

func (*MimeTypeArray) Item(index uint) (p *plugin.Plugin) {
	js.Rewrite("$<.item($1)", index)
	return p
}

func (*MimeTypeArray) NamedItem(kind string) (p *plugin.Plugin) {
	js.Rewrite("$<.namedItem($1)", kind)
	return p
}

func (*MimeTypeArray) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
