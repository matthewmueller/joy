package mimetype

import "github.com/matthewmueller/golly/js"

// MimeType struct
// js:"MimeType,omit"
type MimeType struct {
}

// Description prop
func (*MimeType) Description() (description string) {
	js.Rewrite("$<.description")
	return description
}

// EnabledPlugin prop
func (*MimeType) EnabledPlugin() (enabledPlugin *Plugin) {
	js.Rewrite("$<.enabledPlugin")
	return enabledPlugin
}

// Suffixes prop
func (*MimeType) Suffixes() (suffixes string) {
	js.Rewrite("$<.suffixes")
	return suffixes
}

// Type prop
func (*MimeType) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
