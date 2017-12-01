package window

import "github.com/matthewmueller/golly/js"

var _ Node = (*DocumentFragment)(nil)
var _ EventTarget = (*DocumentFragment)(nil)

// DocumentFragment struct
// js:"DocumentFragment,omit"
type DocumentFragment struct {
}

// QuerySelector fn
// js:"querySelector"
func (*DocumentFragment) QuerySelector(selectors string) (e Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return e
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*DocumentFragment) QuerySelectorAll(selectors string) (n *NodeList) {
	js.Rewrite("$_.querySelectorAll($1)", selectors)
	return n
}

// AppendChild fn
// js:"appendChild"
func (*DocumentFragment) AppendChild(newChild Node) (n Node) {
	js.Rewrite("$_.appendChild($1)", newChild)
	return n
}

// CloneNode fn
// js:"cloneNode"
func (*DocumentFragment) CloneNode(deep *bool) (n Node) {
	js.Rewrite("$_.cloneNode($1)", deep)
	return n
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*DocumentFragment) CompareDocumentPosition(other Node) (u uint8) {
	js.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*DocumentFragment) Contains(child Node) (b bool) {
	js.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*DocumentFragment) HasAttributes() (b bool) {
	js.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*DocumentFragment) HasChildNodes() (b bool) {
	js.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*DocumentFragment) InsertBefore(newChild Node, refChild *Node) (n Node) {
	js.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*DocumentFragment) IsDefaultNamespace(namespaceURI string) (b bool) {
	js.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*DocumentFragment) IsEqualNode(arg Node) (b bool) {
	js.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*DocumentFragment) IsSameNode(other Node) (b bool) {
	js.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*DocumentFragment) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*DocumentFragment) LookupPrefix(namespaceURI string) (s string) {
	js.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*DocumentFragment) Normalize() {
	js.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*DocumentFragment) RemoveChild(oldChild Node) (n Node) {
	js.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

// ReplaceChild fn
// js:"replaceChild"
func (*DocumentFragment) ReplaceChild(newChild Node, oldChild Node) (n Node) {
	js.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

// AddEventListener fn
// js:"addEventListener"
func (*DocumentFragment) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*DocumentFragment) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*DocumentFragment) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Attributes prop
// js:"attributes"
func (*DocumentFragment) Attributes() (attributes *NamedNodeMap) {
	js.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*DocumentFragment) BaseURI() (baseURI string) {
	js.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*DocumentFragment) ChildNodes() (childNodes *NodeList) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*DocumentFragment) FirstChild() (firstChild Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*DocumentFragment) LastChild() (lastChild Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*DocumentFragment) LocalName() (localName string) {
	js.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*DocumentFragment) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*DocumentFragment) NextSibling() (nextSibling Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*DocumentFragment) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*DocumentFragment) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*DocumentFragment) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*DocumentFragment) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*DocumentFragment) OwnerDocument() (ownerDocument Document) {
	js.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*DocumentFragment) ParentElement() (parentElement HTMLElement) {
	js.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*DocumentFragment) ParentNode() (parentNode Node) {
	js.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*DocumentFragment) PreviousSibling() (previousSibling Node) {
	js.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*DocumentFragment) TextContent() (textContent string) {
	js.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*DocumentFragment) SetTextContent(textContent string) {
	js.Rewrite("$_.textContent = $1", textContent)
}
