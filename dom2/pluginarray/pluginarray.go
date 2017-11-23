package pluginarray

import (
	"github.com/matthewmueller/golly/dom2/mimetype"
	"github.com/matthewmueller/golly/js"
)

// js:"PluginArray,omit"
type PluginArray struct {
}

// Item
func (*PluginArray) Item(index uint) (m *mimetype.Plugin) {
	js.Rewrite("$<.Item($1)", index)
	return m
}

// NamedItem
func (*PluginArray) NamedItem(name string) (m *mimetype.Plugin) {
	js.Rewrite("$<.NamedItem($1)", name)
	return m
}

// Refresh
func (*PluginArray) Refresh(reload *bool) {
	js.Rewrite("$<.Refresh($1)", reload)
}

// Length
func (*PluginArray) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
