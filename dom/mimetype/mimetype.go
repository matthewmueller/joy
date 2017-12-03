package mimetype

import "github.com/matthewmueller/joy/macro"

// MimeType struct
// js:"MimeType,omit"
type MimeType struct {
}

// Description prop
// js:"description"
func (*MimeType) Description() (description string) {
	macro.Rewrite("$_.description")
	return description
}

// EnabledPlugin prop
// js:"enabledPlugin"
func (*MimeType) EnabledPlugin() (enabledPlugin *Plugin) {
	macro.Rewrite("$_.enabledPlugin")
	return enabledPlugin
}

// Suffixes prop
// js:"suffixes"
func (*MimeType) Suffixes() (suffixes string) {
	macro.Rewrite("$_.suffixes")
	return suffixes
}

// Type prop
// js:"type"
func (*MimeType) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
