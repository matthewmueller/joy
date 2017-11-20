package plugin

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/mimetype"
	"github.com/matthewmueller/golly/js"
)

type Plugin struct {
}

func (*Plugin) Item(index uint) (m *mimetype.MimeType) {
	js.Rewrite("$<.item($1)", index)
	return m
}

func (*Plugin) NamedItem(kind string) (m *mimetype.MimeType) {
	js.Rewrite("$<.namedItem($1)", kind)
	return m
}

func (*Plugin) GetDescription() (description string) {
	js.Rewrite("$<.description")
	return description
}

func (*Plugin) GetFilename() (filename string) {
	js.Rewrite("$<.filename")
	return filename
}

func (*Plugin) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*Plugin) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*Plugin) GetVersion() (version string) {
	js.Rewrite("$<.version")
	return version
}
