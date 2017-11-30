package node

import "github.com/matthewmueller/golly/dom/eventtarget"

// Node interface
// js:"Node"
type Node interface {
	eventtarget.EventTarget

	// ChildNodes prop
	// js:"childNodes"
	ChildNodes() (childNodes []Node)

	// FirstChild prop
	// js:"firstChild"
	FirstChild() (firstChild Node)

	// LastChild prop
	// js:"lastChild"
	LastChild() (lastChild Node)

	// NextSibling prop
	// js:"nextSibling"
	NextSibling() (nextSibling Node)

	// NodeName prop
	// js:"nodeName"
	NodeName() (nodeName string)

	// NodeType prop
	// js:"nodeType"
	NodeType() (nodeType uint8)

	// NodeValue prop
	// js:"nodeValue"
	NodeValue() (nodeValue string)

	// SetNodeValue prop
	// js:"nodeValue"
	SetNodeValue(nodeValue string)
}
