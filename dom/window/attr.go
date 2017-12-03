package window

import "github.com/matthewmueller/joy/js"

var _ Node = (*Attr)(nil)
var _ EventTarget = (*Attr)(nil)

// Attr struct
// js:"Attr,omit"
type Attr struct {
}

// AppendChild fn
// js:"appendChild"
func (*Attr) AppendChild(newChild Node) (n Node) {
	js.Rewrite("$_.appendChild($1)", newChild)
	return n
}

// CloneNode fn
// js:"cloneNode"
func (*Attr) CloneNode(deep *bool) (n Node) {
	js.Rewrite("$_.cloneNode($1)", deep)
	return n
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*Attr) CompareDocumentPosition(other Node) (u uint8) {
	js.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*Attr) Contains(child Node) (b bool) {
	js.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*Attr) HasAttributes() (b bool) {
	js.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*Attr) HasChildNodes() (b bool) {
	js.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*Attr) InsertBefore(newChild Node, refChild *Node) (n Node) {
	js.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*Attr) IsDefaultNamespace(namespaceURI string) (b bool) {
	js.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*Attr) IsEqualNode(arg Node) (b bool) {
	js.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*Attr) IsSameNode(other Node) (b bool) {
	js.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*Attr) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*Attr) LookupPrefix(namespaceURI string) (s string) {
	js.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*Attr) Normalize() {
	js.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*Attr) RemoveChild(oldChild Node) (n Node) {
	js.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

// ReplaceChild fn
// js:"replaceChild"
func (*Attr) ReplaceChild(newChild Node, oldChild Node) (n Node) {
	js.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

// AddEventListener fn
// js:"addEventListener"
func (*Attr) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*Attr) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*Attr) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Name prop
// js:"name"
func (*Attr) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// OwnerElement prop
// js:"ownerElement"
func (*Attr) OwnerElement() (ownerElement Element) {
	js.Rewrite("$_.ownerElement")
	return ownerElement
}

// Prefix prop
// js:"prefix"
func (*Attr) Prefix() (prefix string) {
	js.Rewrite("$_.prefix")
	return prefix
}

// Specified prop
// js:"specified"
func (*Attr) Specified() (specified bool) {
	js.Rewrite("$_.specified")
	return specified
}

// Value prop
// js:"value"
func (*Attr) Value() (value string) {
	js.Rewrite("$_.value")
	return value
}

// SetValue prop
// js:"value"
func (*Attr) SetValue(value string) {
	js.Rewrite("$_.value = $1", value)
}

// Attributes prop
// js:"attributes"
func (*Attr) Attributes() (attributes *NamedNodeMap) {
	js.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*Attr) BaseURI() (baseURI string) {
	js.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*Attr) ChildNodes() (childNodes *NodeList) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*Attr) FirstChild() (firstChild Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*Attr) LastChild() (lastChild Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*Attr) LocalName() (localName string) {
	js.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*Attr) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*Attr) NextSibling() (nextSibling Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*Attr) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*Attr) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*Attr) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*Attr) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*Attr) OwnerDocument() (ownerDocument Document) {
	js.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*Attr) ParentElement() (parentElement HTMLElement) {
	js.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*Attr) ParentNode() (parentNode Node) {
	js.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*Attr) PreviousSibling() (previousSibling Node) {
	js.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*Attr) TextContent() (textContent string) {
	js.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*Attr) SetTextContent(textContent string) {
	js.Rewrite("$_.textContent = $1", textContent)
}
