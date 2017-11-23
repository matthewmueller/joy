package idb

import "github.com/matthewmueller/golly/js"

// NamedNodeMap struct
// js:"NamedNodeMap,omit"
type NamedNodeMap struct {
}

// GetNamedItem
func (*NamedNodeMap) GetNamedItem(name string) (a *Attr) {
	js.Rewrite("$<.GetNamedItem($1)", name)
	return a
}

// GetNamedItemNS
func (*NamedNodeMap) GetNamedItemNS(namespaceURI string, localName string) (a *Attr) {
	js.Rewrite("$<.GetNamedItemNS($1, $2)", namespaceURI, localName)
	return a
}

// Item
func (*NamedNodeMap) Item(index uint) (a *Attr) {
	js.Rewrite("$<.Item($1)", index)
	return a
}

// RemoveNamedItem
func (*NamedNodeMap) RemoveNamedItem(name string) (a *Attr) {
	js.Rewrite("$<.RemoveNamedItem($1)", name)
	return a
}

// RemoveNamedItemNS
func (*NamedNodeMap) RemoveNamedItemNS(namespaceURI string, localName string) (a *Attr) {
	js.Rewrite("$<.RemoveNamedItemNS($1, $2)", namespaceURI, localName)
	return a
}

// SetNamedItem
func (*NamedNodeMap) SetNamedItem(arg *Attr) (a *Attr) {
	js.Rewrite("$<.SetNamedItem($1)", arg)
	return a
}

// SetNamedItemNS
func (*NamedNodeMap) SetNamedItemNS(arg *Attr) (a *Attr) {
	js.Rewrite("$<.SetNamedItemNS($1)", arg)
	return a
}

// Length
func (*NamedNodeMap) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
