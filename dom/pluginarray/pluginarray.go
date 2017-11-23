package pluginarray

import (
	"github.com/matthewmueller/golly/dom2/mimetype"
	"github.com/matthewmueller/golly/js"
)

// PluginArray struct
// js:"PluginArray,omit"
type PluginArray struct {
}

// Item fn
func (*PluginArray) Item(index uint) (m *mimetype.Plugin) {
	js.Rewrite("$<.item($1)", index)
	return m
}

// NamedItem fn
func (*PluginArray) NamedItem(name string) (m *mimetype.Plugin) {
	js.Rewrite("$<.namedItem($1)", name)
	return m
}

// Refresh fn
func (*PluginArray) Refresh(reload *bool) {
	js.Rewrite("$<.refresh($1)", reload)
}

// Length prop
func (*PluginArray) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
