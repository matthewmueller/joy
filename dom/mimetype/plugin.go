package mimetype

import "github.com/matthewmueller/joy/js"

// Plugin struct
// js:"Plugin,omit"
type Plugin struct {
}

// Item fn
// js:"item"
func (*Plugin) Item(index uint) (m *MimeType) {
	js.Rewrite("$_.item($1)", index)
	return m
}

// NamedItem fn
// js:"namedItem"
func (*Plugin) NamedItem(kind string) (m *MimeType) {
	js.Rewrite("$_.namedItem($1)", kind)
	return m
}

// Description prop
// js:"description"
func (*Plugin) Description() (description string) {
	js.Rewrite("$_.description")
	return description
}

// Filename prop
// js:"filename"
func (*Plugin) Filename() (filename string) {
	js.Rewrite("$_.filename")
	return filename
}

// Length prop
// js:"length"
func (*Plugin) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}

// Name prop
// js:"name"
func (*Plugin) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// Version prop
// js:"version"
func (*Plugin) Version() (version string) {
	js.Rewrite("$_.version")
	return version
}
