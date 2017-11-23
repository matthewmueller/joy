package window

// js:"Node,omit"
type Node interface {
	EventTarget

	// AppendChild
	AppendChild(newChild Node) (n Node)

	// CloneNode
	CloneNode(deep *bool) (n Node)

	// CompareDocumentPosition
	CompareDocumentPosition(other Node) (u uint8)

	// Contains
	Contains(child Node) (b bool)

	// HasAttributes
	HasAttributes() (b bool)

	// HasChildNodes
	HasChildNodes() (b bool)

	// InsertBefore
	InsertBefore(newChild Node, refChild *Node) (n Node)

	// IsDefaultNamespace
	IsDefaultNamespace(namespaceURI string) (b bool)

	// IsEqualNode
	IsEqualNode(arg Node) (b bool)

	// IsSameNode
	IsSameNode(other Node) (b bool)

	// LookupNamespaceURI
	LookupNamespaceURI(prefix string) (s string)

	// LookupPrefix
	LookupPrefix(namespaceURI string) (s string)

	// Normalize
	Normalize()

	// RemoveChild
	RemoveChild(oldChild Node) (n Node)

	// ReplaceChild
	ReplaceChild(newChild Node, oldChild Node) (n Node)

	// Attributes
	Attributes() (attributes *NamedNodeMap)

	// BaseURI
	BaseURI() (baseURI string)

	// ChildNodes
	ChildNodes() (childNodes *NodeList)

	// FirstChild
	FirstChild() (firstChild Node)

	// LastChild
	LastChild() (lastChild Node)

	// LocalName
	LocalName() (localName string)

	// NamespaceURI
	NamespaceURI() (namespaceURI string)

	// NextSibling
	NextSibling() (nextSibling Node)

	// NodeName
	NodeName() (nodeName string)

	// NodeType
	NodeType() (nodeType uint8)

	// NodeValue
	NodeValue() (nodeValue string)

	// NodeValue
	SetNodeValue(nodeValue string)

	// OwnerDocument
	OwnerDocument() (ownerDocument Document)

	// ParentElement
	ParentElement() (parentElement HTMLElement)

	// ParentNode
	ParentNode() (parentNode Node)

	// PreviousSibling
	PreviousSibling() (previousSibling Node)

	// TextContent
	TextContent() (textContent string)

	// TextContent
	SetTextContent(textContent string)
}
