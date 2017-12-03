package window

import "github.com/matthewmueller/joy/macro"

// NamedNodeMap struct
// js:"NamedNodeMap,omit"
type NamedNodeMap struct {
}

// GetNamedItem fn
// js:"getNamedItem"
func (*NamedNodeMap) GetNamedItem(name string) (a *Attr) {
	macro.Rewrite("$_.getNamedItem($1)", name)
	return a
}

// GetNamedItemNS fn
// js:"getNamedItemNS"
func (*NamedNodeMap) GetNamedItemNS(namespaceURI string, localName string) (a *Attr) {
	macro.Rewrite("$_.getNamedItemNS($1, $2)", namespaceURI, localName)
	return a
}

// Item fn
// js:"item"
func (*NamedNodeMap) Item(index uint) (a *Attr) {
	macro.Rewrite("$_.item($1)", index)
	return a
}

// RemoveNamedItem fn
// js:"removeNamedItem"
func (*NamedNodeMap) RemoveNamedItem(name string) (a *Attr) {
	macro.Rewrite("$_.removeNamedItem($1)", name)
	return a
}

// RemoveNamedItemNS fn
// js:"removeNamedItemNS"
func (*NamedNodeMap) RemoveNamedItemNS(namespaceURI string, localName string) (a *Attr) {
	macro.Rewrite("$_.removeNamedItemNS($1, $2)", namespaceURI, localName)
	return a
}

// SetNamedItem fn
// js:"setNamedItem"
func (*NamedNodeMap) SetNamedItem(arg *Attr) (a *Attr) {
	macro.Rewrite("$_.setNamedItem($1)", arg)
	return a
}

// SetNamedItemNS fn
// js:"setNamedItemNS"
func (*NamedNodeMap) SetNamedItemNS(arg *Attr) (a *Attr) {
	macro.Rewrite("$_.setNamedItemNS($1)", arg)
	return a
}

// Length prop
// js:"length"
func (*NamedNodeMap) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
