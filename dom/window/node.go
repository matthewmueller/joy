package window

// Node interface
// js:"Node"
type Node interface {
	EventTarget

	// AppendChild
	// js:"appendChild"
	AppendChild(newChild Node) (n Node)

	// CloneNode
	// js:"cloneNode"
	CloneNode(deep *bool) (n Node)

	// CompareDocumentPosition
	// js:"compareDocumentPosition"
	CompareDocumentPosition(other Node) (u uint8)

	// Contains
	// js:"contains"
	Contains(child Node) (b bool)

	// HasAttributes
	// js:"hasAttributes"
	HasAttributes() (b bool)

	// HasChildNodes
	// js:"hasChildNodes"
	HasChildNodes() (b bool)

	// InsertBefore
	// js:"insertBefore"
	InsertBefore(newChild Node, refChild *Node) (n Node)

	// IsDefaultNamespace
	// js:"isDefaultNamespace"
	IsDefaultNamespace(namespaceURI string) (b bool)

	// IsEqualNode
	// js:"isEqualNode"
	IsEqualNode(arg Node) (b bool)

	// IsSameNode
	// js:"isSameNode"
	IsSameNode(other Node) (b bool)

	// LookupNamespaceURI
	// js:"lookupNamespaceURI"
	LookupNamespaceURI(prefix string) (s string)

	// LookupPrefix
	// js:"lookupPrefix"
	LookupPrefix(namespaceURI string) (s string)

	// Normalize
	// js:"normalize"
	Normalize()

	// RemoveChild
	// js:"removeChild"
	RemoveChild(oldChild Node) (n Node)

	// ReplaceChild
	// js:"replaceChild"
	ReplaceChild(newChild Node, oldChild Node) (n Node)

	// Attributes prop
	// js:"attributes"
	Attributes() (attributes *NamedNodeMap)

	// BaseURI prop
	// js:"baseURI"
	BaseURI() (baseURI string)

	// ChildNodes prop
	// js:"childNodes"
	ChildNodes() (childNodes *NodeList)

	// FirstChild prop
	// js:"firstChild"
	FirstChild() (firstChild Node)

	// LastChild prop
	// js:"lastChild"
	LastChild() (lastChild Node)

	// LocalName prop
	// js:"localName"
	LocalName() (localName string)

	// NamespaceURI prop
	// js:"namespaceURI"
	NamespaceURI() (namespaceURI string)

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

	// OwnerDocument prop
	// js:"ownerDocument"
	OwnerDocument() (ownerDocument Document)

	// ParentElement prop
	// js:"parentElement"
	ParentElement() (parentElement HTMLElement)

	// ParentNode prop
	// js:"parentNode"
	ParentNode() (parentNode Node)

	// PreviousSibling prop
	// js:"previousSibling"
	PreviousSibling() (previousSibling Node)

	// TextContent prop
	// js:"textContent"
	TextContent() (textContent string)

	// SetTextContent prop
	// js:"textContent"
	SetTextContent(textContent string)
}
