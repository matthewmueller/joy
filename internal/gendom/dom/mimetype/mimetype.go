package mimetype

import (
	"plugin"

	"github.com/matthewmueller/golly/js"
)

type MimeType struct {
}

func (*MimeType) GetDescription() (description string) {
	js.Rewrite("$<.description")
	return description
}

func (*MimeType) GetEnabledPlugin() (enabledPlugin *plugin.Plugin) {
	js.Rewrite("$<.enabledPlugin")
	return enabledPlugin
}

func (*MimeType) GetSuffixes() (suffixes string) {
	js.Rewrite("$<.suffixes")
	return suffixes
}

func (*MimeType) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
