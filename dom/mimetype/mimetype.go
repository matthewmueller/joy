package mimetype

import "github.com/matthewmueller/joy/js"

// MimeType struct
// js:"MimeType,omit"
type MimeType struct {
}

// Description prop
// js:"description"
func (*MimeType) Description() (description string) {
	js.Rewrite("$_.description")
	return description
}

// EnabledPlugin prop
// js:"enabledPlugin"
func (*MimeType) EnabledPlugin() (enabledPlugin *Plugin) {
	js.Rewrite("$_.enabledPlugin")
	return enabledPlugin
}

// Suffixes prop
// js:"suffixes"
func (*MimeType) Suffixes() (suffixes string) {
	js.Rewrite("$_.suffixes")
	return suffixes
}

// Type prop
// js:"type"
func (*MimeType) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
