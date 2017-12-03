package window

import (
	"github.com/matthewmueller/joy/dom/childnode"
	"github.com/matthewmueller/joy/js"
)

var _ Text = (*CDATASection)(nil)
var _ CharacterData = (*CDATASection)(nil)
var _ childnode.ChildNode = (*CDATASection)(nil)
var _ Node = (*CDATASection)(nil)
var _ EventTarget = (*CDATASection)(nil)

// CDATASection struct
// js:"CDATASection,omit"
type CDATASection struct {
}

// SplitText fn
// js:"splitText"
func (*CDATASection) SplitText(offset uint) (t Text) {
	js.Rewrite("$_.splitText($1)", offset)
	return t
}

// AppendData fn
// js:"appendData"
func (*CDATASection) AppendData(arg string) {
	js.Rewrite("$_.appendData($1)", arg)
}

// DeleteData fn
// js:"deleteData"
func (*CDATASection) DeleteData(offset uint, count uint) {
	js.Rewrite("$_.deleteData($1, $2)", offset, count)
}

// InsertData fn
// js:"insertData"
func (*CDATASection) InsertData(offset uint, arg string) {
	js.Rewrite("$_.insertData($1, $2)", offset, arg)
}

// ReplaceData fn
// js:"replaceData"
func (*CDATASection) ReplaceData(offset uint, count uint, arg string) {
	js.Rewrite("$_.replaceData($1, $2, $3)", offset, count, arg)
}

// SubstringData fn
// js:"substringData"
func (*CDATASection) SubstringData(offset uint, count uint) (s string) {
	js.Rewrite("$_.substringData($1, $2)", offset, count)
	return s
}

// AppendChild fn
// js:"appendChild"
func (*CDATASection) AppendChild(newChild Node) (n Node) {
	js.Rewrite("$_.appendChild($1)", newChild)
	return n
}

// CloneNode fn
// js:"cloneNode"
func (*CDATASection) CloneNode(deep *bool) (n Node) {
	js.Rewrite("$_.cloneNode($1)", deep)
	return n
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*CDATASection) CompareDocumentPosition(other Node) (u uint8) {
	js.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*CDATASection) Contains(child Node) (b bool) {
	js.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*CDATASection) HasAttributes() (b bool) {
	js.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*CDATASection) HasChildNodes() (b bool) {
	js.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*CDATASection) InsertBefore(newChild Node, refChild *Node) (n Node) {
	js.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*CDATASection) IsDefaultNamespace(namespaceURI string) (b bool) {
	js.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*CDATASection) IsEqualNode(arg Node) (b bool) {
	js.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*CDATASection) IsSameNode(other Node) (b bool) {
	js.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*CDATASection) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*CDATASection) LookupPrefix(namespaceURI string) (s string) {
	js.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*CDATASection) Normalize() {
	js.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*CDATASection) RemoveChild(oldChild Node) (n Node) {
	js.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

// ReplaceChild fn
// js:"replaceChild"
func (*CDATASection) ReplaceChild(newChild Node, oldChild Node) (n Node) {
	js.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

// AddEventListener fn
// js:"addEventListener"
func (*CDATASection) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*CDATASection) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*CDATASection) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// WholeText prop
// js:"wholeText"
func (*CDATASection) WholeText() (wholeText string) {
	js.Rewrite("$_.wholeText")
	return wholeText
}

// Data prop
// js:"data"
func (*CDATASection) Data() (data string) {
	js.Rewrite("$_.data")
	return data
}

// SetData prop
// js:"data"
func (*CDATASection) SetData(data string) {
	js.Rewrite("$_.data = $1", data)
}

// Length prop
// js:"length"
func (*CDATASection) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}

// Attributes prop
// js:"attributes"
func (*CDATASection) Attributes() (attributes *NamedNodeMap) {
	js.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*CDATASection) BaseURI() (baseURI string) {
	js.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*CDATASection) ChildNodes() (childNodes *NodeList) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*CDATASection) FirstChild() (firstChild Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*CDATASection) LastChild() (lastChild Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*CDATASection) LocalName() (localName string) {
	js.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*CDATASection) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*CDATASection) NextSibling() (nextSibling Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*CDATASection) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*CDATASection) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*CDATASection) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*CDATASection) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*CDATASection) OwnerDocument() (ownerDocument Document) {
	js.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*CDATASection) ParentElement() (parentElement HTMLElement) {
	js.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*CDATASection) ParentNode() (parentNode Node) {
	js.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*CDATASection) PreviousSibling() (previousSibling Node) {
	js.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*CDATASection) TextContent() (textContent string) {
	js.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*CDATASection) SetTextContent(textContent string) {
	js.Rewrite("$_.textContent = $1", textContent)
}
