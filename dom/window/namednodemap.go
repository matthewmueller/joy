package window

import "github.com/matthewmueller/golly/js"

// NamedNodeMap struct
// js:"NamedNodeMap,omit"
type NamedNodeMap struct {
}

// GetNamedItem fn
func (*NamedNodeMap) GetNamedItem(name string) (a *Attr) {
	js.Rewrite("$<.getNamedItem($1)", name)
	return a
}

// GetNamedItemNS fn
func (*NamedNodeMap) GetNamedItemNS(namespaceURI string, localName string) (a *Attr) {
	js.Rewrite("$<.getNamedItemNS($1, $2)", namespaceURI, localName)
	return a
}

// Item fn
func (*NamedNodeMap) Item(index uint) (a *Attr) {
	js.Rewrite("$<.item($1)", index)
	return a
}

// RemoveNamedItem fn
func (*NamedNodeMap) RemoveNamedItem(name string) (a *Attr) {
	js.Rewrite("$<.removeNamedItem($1)", name)
	return a
}

// RemoveNamedItemNS fn
func (*NamedNodeMap) RemoveNamedItemNS(namespaceURI string, localName string) (a *Attr) {
	js.Rewrite("$<.removeNamedItemNS($1, $2)", namespaceURI, localName)
	return a
}

// SetNamedItem fn
func (*NamedNodeMap) SetNamedItem(arg *Attr) (a *Attr) {
	js.Rewrite("$<.setNamedItem($1)", arg)
	return a
}

// SetNamedItemNS fn
func (*NamedNodeMap) SetNamedItemNS(arg *Attr) (a *Attr) {
	js.Rewrite("$<.setNamedItemNS($1)", arg)
	return a
}

// Length prop
func (*NamedNodeMap) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
