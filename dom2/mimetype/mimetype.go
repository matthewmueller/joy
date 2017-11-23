package mimetype

import "github.com/matthewmueller/golly/js"

// MimeType struct
// js:"MimeType,omit"
type MimeType struct {
}

// Description
func (*MimeType) Description() (description string) {
	js.Rewrite("$<.Description")
	return description
}

// EnabledPlugin
func (*MimeType) EnabledPlugin() (enabledPlugin *Plugin) {
	js.Rewrite("$<.EnabledPlugin")
	return enabledPlugin
}

// Suffixes
func (*MimeType) Suffixes() (suffixes string) {
	js.Rewrite("$<.Suffixes")
	return suffixes
}

// Type
func (*MimeType) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}
