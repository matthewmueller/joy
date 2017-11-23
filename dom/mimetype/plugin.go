package mimetype

import "github.com/matthewmueller/golly/js"

// Plugin struct
// js:"Plugin,omit"
type Plugin struct {
}

// Item fn
func (*Plugin) Item(index uint) (m *MimeType) {
	js.Rewrite("$<.item($1)", index)
	return m
}

// NamedItem fn
func (*Plugin) NamedItem(kind string) (m *MimeType) {
	js.Rewrite("$<.namedItem($1)", kind)
	return m
}

// Description prop
func (*Plugin) Description() (description string) {
	js.Rewrite("$<.description")
	return description
}

// Filename prop
func (*Plugin) Filename() (filename string) {
	js.Rewrite("$<.filename")
	return filename
}

// Length prop
func (*Plugin) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}

// Name prop
func (*Plugin) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Version prop
func (*Plugin) Version() (version string) {
	js.Rewrite("$<.version")
	return version
}
