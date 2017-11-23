package mimetype

import "github.com/matthewmueller/golly/js"

// js:"Plugin,omit"
type Plugin struct {
}

// Item
func (*Plugin) Item(index uint) (m *MimeType) {
	js.Rewrite("$<.Item($1)", index)
	return m
}

// NamedItem
func (*Plugin) NamedItem(kind string) (m *MimeType) {
	js.Rewrite("$<.NamedItem($1)", kind)
	return m
}

// Description
func (*Plugin) Description() (description string) {
	js.Rewrite("$<.Description")
	return description
}

// Filename
func (*Plugin) Filename() (filename string) {
	js.Rewrite("$<.Filename")
	return filename
}

// Length
func (*Plugin) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}

// Name
func (*Plugin) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Version
func (*Plugin) Version() (version string) {
	js.Rewrite("$<.Version")
	return version
}
