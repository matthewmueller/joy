package window

import "github.com/matthewmueller/golly/js"

var _ Node = (*DocumentType)(nil)
var _ EventTarget = (*DocumentType)(nil)

// DocumentType struct
// js:"DocumentType,omit"
type DocumentType struct {
}

// AppendChild fn
// js:"appendChild"
func (*DocumentType) AppendChild(newChild Node) (n Node) {
	js.Rewrite("$_.appendChild($1)", newChild)
	return n
}

// CloneNode fn
// js:"cloneNode"
func (*DocumentType) CloneNode(deep *bool) (n Node) {
	js.Rewrite("$_.cloneNode($1)", deep)
	return n
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*DocumentType) CompareDocumentPosition(other Node) (u uint8) {
	js.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*DocumentType) Contains(child Node) (b bool) {
	js.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*DocumentType) HasAttributes() (b bool) {
	js.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*DocumentType) HasChildNodes() (b bool) {
	js.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*DocumentType) InsertBefore(newChild Node, refChild *Node) (n Node) {
	js.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*DocumentType) IsDefaultNamespace(namespaceURI string) (b bool) {
	js.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*DocumentType) IsEqualNode(arg Node) (b bool) {
	js.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*DocumentType) IsSameNode(other Node) (b bool) {
	js.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*DocumentType) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*DocumentType) LookupPrefix(namespaceURI string) (s string) {
	js.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*DocumentType) Normalize() {
	js.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*DocumentType) RemoveChild(oldChild Node) (n Node) {
	js.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

// ReplaceChild fn
// js:"replaceChild"
func (*DocumentType) ReplaceChild(newChild Node, oldChild Node) (n Node) {
	js.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

// AddEventListener fn
// js:"addEventListener"
func (*DocumentType) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*DocumentType) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*DocumentType) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Entities prop
// js:"entities"
func (*DocumentType) Entities() (entities *NamedNodeMap) {
	js.Rewrite("$_.entities")
	return entities
}

// InternalSubset prop
// js:"internalSubset"
func (*DocumentType) InternalSubset() (internalSubset string) {
	js.Rewrite("$_.internalSubset")
	return internalSubset
}

// Name prop
// js:"name"
func (*DocumentType) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// Notations prop
// js:"notations"
func (*DocumentType) Notations() (notations *NamedNodeMap) {
	js.Rewrite("$_.notations")
	return notations
}

// PublicID prop
// js:"publicId"
func (*DocumentType) PublicID() (publicId string) {
	js.Rewrite("$_.publicId")
	return publicId
}

// SystemID prop
// js:"systemId"
func (*DocumentType) SystemID() (systemId string) {
	js.Rewrite("$_.systemId")
	return systemId
}

// Attributes prop
// js:"attributes"
func (*DocumentType) Attributes() (attributes *NamedNodeMap) {
	js.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*DocumentType) BaseURI() (baseURI string) {
	js.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*DocumentType) ChildNodes() (childNodes *NodeList) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*DocumentType) FirstChild() (firstChild Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*DocumentType) LastChild() (lastChild Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*DocumentType) LocalName() (localName string) {
	js.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*DocumentType) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*DocumentType) NextSibling() (nextSibling Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*DocumentType) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*DocumentType) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*DocumentType) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*DocumentType) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*DocumentType) OwnerDocument() (ownerDocument Document) {
	js.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*DocumentType) ParentElement() (parentElement HTMLElement) {
	js.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*DocumentType) ParentNode() (parentNode Node) {
	js.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*DocumentType) PreviousSibling() (previousSibling Node) {
	js.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*DocumentType) TextContent() (textContent string) {
	js.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*DocumentType) SetTextContent(textContent string) {
	js.Rewrite("$_.textContent = $1", textContent)
}
