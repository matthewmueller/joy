package window

import (
	"github.com/matthewmueller/joy/dom/childnode"
	"github.com/matthewmueller/joy/macro"
)

var _ CharacterData = (*ProcessingInstruction)(nil)
var _ childnode.ChildNode = (*ProcessingInstruction)(nil)
var _ Node = (*ProcessingInstruction)(nil)
var _ EventTarget = (*ProcessingInstruction)(nil)

// ProcessingInstruction struct
// js:"ProcessingInstruction,omit"
type ProcessingInstruction struct {
}

// AppendData fn
// js:"appendData"
func (*ProcessingInstruction) AppendData(arg string) {
	macro.Rewrite("$_.appendData($1)", arg)
}

// DeleteData fn
// js:"deleteData"
func (*ProcessingInstruction) DeleteData(offset uint, count uint) {
	macro.Rewrite("$_.deleteData($1, $2)", offset, count)
}

// InsertData fn
// js:"insertData"
func (*ProcessingInstruction) InsertData(offset uint, arg string) {
	macro.Rewrite("$_.insertData($1, $2)", offset, arg)
}

// ReplaceData fn
// js:"replaceData"
func (*ProcessingInstruction) ReplaceData(offset uint, count uint, arg string) {
	macro.Rewrite("$_.replaceData($1, $2, $3)", offset, count, arg)
}

// SubstringData fn
// js:"substringData"
func (*ProcessingInstruction) SubstringData(offset uint, count uint) (s string) {
	macro.Rewrite("$_.substringData($1, $2)", offset, count)
	return s
}

// AppendChild fn
// js:"appendChild"
func (*ProcessingInstruction) AppendChild(newChild Node) (n Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return n
}

// CloneNode fn
// js:"cloneNode"
func (*ProcessingInstruction) CloneNode(deep *bool) (n Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return n
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*ProcessingInstruction) CompareDocumentPosition(other Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*ProcessingInstruction) Contains(child Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*ProcessingInstruction) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*ProcessingInstruction) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*ProcessingInstruction) InsertBefore(newChild Node, refChild *Node) (n Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*ProcessingInstruction) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*ProcessingInstruction) IsEqualNode(arg Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*ProcessingInstruction) IsSameNode(other Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*ProcessingInstruction) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*ProcessingInstruction) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*ProcessingInstruction) Normalize() {
	macro.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*ProcessingInstruction) RemoveChild(oldChild Node) (n Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

// ReplaceChild fn
// js:"replaceChild"
func (*ProcessingInstruction) ReplaceChild(newChild Node, oldChild Node) (n Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

// AddEventListener fn
// js:"addEventListener"
func (*ProcessingInstruction) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*ProcessingInstruction) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*ProcessingInstruction) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Target prop
// js:"target"
func (*ProcessingInstruction) Target() (target string) {
	macro.Rewrite("$_.target")
	return target
}

// Data prop
// js:"data"
func (*ProcessingInstruction) Data() (data string) {
	macro.Rewrite("$_.data")
	return data
}

// SetData prop
// js:"data"
func (*ProcessingInstruction) SetData(data string) {
	macro.Rewrite("$_.data = $1", data)
}

// Length prop
// js:"length"
func (*ProcessingInstruction) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

// Attributes prop
// js:"attributes"
func (*ProcessingInstruction) Attributes() (attributes *NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*ProcessingInstruction) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*ProcessingInstruction) ChildNodes() (childNodes *NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*ProcessingInstruction) FirstChild() (firstChild Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*ProcessingInstruction) LastChild() (lastChild Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*ProcessingInstruction) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*ProcessingInstruction) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*ProcessingInstruction) NextSibling() (nextSibling Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*ProcessingInstruction) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*ProcessingInstruction) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*ProcessingInstruction) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*ProcessingInstruction) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*ProcessingInstruction) OwnerDocument() (ownerDocument Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*ProcessingInstruction) ParentElement() (parentElement HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*ProcessingInstruction) ParentNode() (parentNode Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*ProcessingInstruction) PreviousSibling() (previousSibling Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*ProcessingInstruction) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*ProcessingInstruction) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
