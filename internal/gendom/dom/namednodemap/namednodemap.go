package namednodemap

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/attr"
	"github.com/matthewmueller/golly/js"
)

type NamedNodeMap struct {
}

func (*NamedNodeMap) GetNamedItem(name string) (a *attr.Attr) {
	js.Rewrite("$<.getNamedItem($1)", name)
	return a
}

func (*NamedNodeMap) GetNamedItemNS(namespaceURI string, localName string) (a *attr.Attr) {
	js.Rewrite("$<.getNamedItemNS($1, $2)", namespaceURI, localName)
	return a
}

func (*NamedNodeMap) Item(index uint) (a *attr.Attr) {
	js.Rewrite("$<.item($1)", index)
	return a
}

func (*NamedNodeMap) RemoveNamedItem(name string) (a *attr.Attr) {
	js.Rewrite("$<.removeNamedItem($1)", name)
	return a
}

func (*NamedNodeMap) RemoveNamedItemNS(namespaceURI string, localName string) (a *attr.Attr) {
	js.Rewrite("$<.removeNamedItemNS($1, $2)", namespaceURI, localName)
	return a
}

func (*NamedNodeMap) SetNamedItem(arg *attr.Attr) (a *attr.Attr) {
	js.Rewrite("$<.setNamedItem($1)", arg)
	return a
}

func (*NamedNodeMap) SetNamedItemNS(arg *attr.Attr) (a *attr.Attr) {
	js.Rewrite("$<.setNamedItemNS($1)", arg)
	return a
}

func (*NamedNodeMap) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
