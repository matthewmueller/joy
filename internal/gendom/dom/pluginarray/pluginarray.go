package pluginarray

import (
	"plugin"

	"github.com/matthewmueller/golly/js"
)

type PluginArray struct {
}

func (*PluginArray) Item(index uint) (p *plugin.Plugin) {
	js.Rewrite("$<.item($1)", index)
	return p
}

func (*PluginArray) NamedItem(name string) (p *plugin.Plugin) {
	js.Rewrite("$<.namedItem($1)", name)
	return p
}

func (*PluginArray) Refresh(reload bool) {
	js.Rewrite("$<.refresh($1)", reload)
}

func (*PluginArray) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
