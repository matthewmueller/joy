package htmldocument

import (
	"github.com/matthewmueller/golly/dom/document"
	"github.com/matthewmueller/golly/dom/element"
	"github.com/matthewmueller/golly/dom/event"
	"github.com/matthewmueller/golly/dom/eventtarget"
	"github.com/matthewmueller/golly/dom/node"
	"github.com/matthewmueller/golly/js"
)

var _ document.Document = (*HTMLDocument)(nil)
var _ element.NodeSelector = (*HTMLDocument)(nil)
var _ node.Node = (*HTMLDocument)(nil)
var _ eventtarget.EventTarget = (*HTMLDocument)(nil)

// HTMLDocument struct
// js:"HTMLDocument,omit"
type HTMLDocument struct {
}

// QuerySelector fn
// js:"querySelector"
func (*HTMLDocument) QuerySelector(selectors string) (e element.Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return e
}

// AddEventListener fn
// js:"addEventListener"
func (*HTMLDocument) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*HTMLDocument) DispatchEvent(evt event.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// DocumentElement prop Gets a reference to the root node of the document.
// js:"documentElement"
func (*HTMLDocument) DocumentElement() (documentElement element.Element) {
	js.Rewrite("$_.documentElement")
	return documentElement
}

// ChildNodes prop
// js:"childNodes"
func (*HTMLDocument) ChildNodes() (childNodes []node.Node) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*HTMLDocument) FirstChild() (firstChild node.Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*HTMLDocument) LastChild() (lastChild node.Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// NextSibling prop
// js:"nextSibling"
func (*HTMLDocument) NextSibling() (nextSibling node.Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*HTMLDocument) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*HTMLDocument) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*HTMLDocument) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*HTMLDocument) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}
