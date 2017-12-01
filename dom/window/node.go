package window

// Node interface
// js:"Node"
type Node interface {
	EventTarget

	// AppendChild
	// js:"appendChild"
	// jsrewrite:"$_.appendChild($1)"
	AppendChild(newChild Node) (n Node)

	// CloneNode
	// js:"cloneNode"
	// jsrewrite:"$_.cloneNode($1)"
	CloneNode(deep *bool) (n Node)

	// CompareDocumentPosition
	// js:"compareDocumentPosition"
	// jsrewrite:"$_.compareDocumentPosition($1)"
	CompareDocumentPosition(other Node) (u uint8)

	// Contains
	// js:"contains"
	// jsrewrite:"$_.contains($1)"
	Contains(child Node) (b bool)

	// HasAttributes
	// js:"hasAttributes"
	// jsrewrite:"$_.hasAttributes()"
	HasAttributes() (b bool)

	// HasChildNodes
	// js:"hasChildNodes"
	// jsrewrite:"$_.hasChildNodes()"
	HasChildNodes() (b bool)

	// InsertBefore
	// js:"insertBefore"
	// jsrewrite:"$_.insertBefore($1, $2)"
	InsertBefore(newChild Node, refChild *Node) (n Node)

	// IsDefaultNamespace
	// js:"isDefaultNamespace"
	// jsrewrite:"$_.isDefaultNamespace($1)"
	IsDefaultNamespace(namespaceURI string) (b bool)

	// IsEqualNode
	// js:"isEqualNode"
	// jsrewrite:"$_.isEqualNode($1)"
	IsEqualNode(arg Node) (b bool)

	// IsSameNode
	// js:"isSameNode"
	// jsrewrite:"$_.isSameNode($1)"
	IsSameNode(other Node) (b bool)

	// LookupNamespaceURI
	// js:"lookupNamespaceURI"
	// jsrewrite:"$_.lookupNamespaceURI($1)"
	LookupNamespaceURI(prefix string) (s string)

	// LookupPrefix
	// js:"lookupPrefix"
	// jsrewrite:"$_.lookupPrefix($1)"
	LookupPrefix(namespaceURI string) (s string)

	// Normalize
	// js:"normalize"
	// jsrewrite:"$_.normalize()"
	Normalize()

	// RemoveChild
	// js:"removeChild"
	// jsrewrite:"$_.removeChild($1)"
	RemoveChild(oldChild Node) (n Node)

	// ReplaceChild
	// js:"replaceChild"
	// jsrewrite:"$_.replaceChild($1, $2)"
	ReplaceChild(newChild Node, oldChild Node) (n Node)

	// Attributes prop
	// js:"attributes"
	// jsrewrite:"$_.attributes"
	Attributes() (attributes *NamedNodeMap)

	// BaseURI prop
	// js:"baseURI"
	// jsrewrite:"$_.baseURI"
	BaseURI() (baseURI string)

	// ChildNodes prop
	// js:"childNodes"
	// jsrewrite:"$_.childNodes"
	ChildNodes() (childNodes *NodeList)

	// FirstChild prop
	// js:"firstChild"
	// jsrewrite:"$_.firstChild"
	FirstChild() (firstChild Node)

	// LastChild prop
	// js:"lastChild"
	// jsrewrite:"$_.lastChild"
	LastChild() (lastChild Node)

	// LocalName prop
	// js:"localName"
	// jsrewrite:"$_.localName"
	LocalName() (localName string)

	// NamespaceURI prop
	// js:"namespaceURI"
	// jsrewrite:"$_.namespaceURI"
	NamespaceURI() (namespaceURI string)

	// NextSibling prop
	// js:"nextSibling"
	// jsrewrite:"$_.nextSibling"
	NextSibling() (nextSibling Node)

	// NodeName prop
	// js:"nodeName"
	// jsrewrite:"$_.nodeName"
	NodeName() (nodeName string)

	// NodeType prop
	// js:"nodeType"
	// jsrewrite:"$_.nodeType"
	NodeType() (nodeType uint8)

	// NodeValue prop
	// js:"nodeValue"
	// jsrewrite:"$_.nodeValue"
	NodeValue() (nodeValue string)

	// SetNodeValue prop
	// js:"nodeValue"
	// jsrewrite:"$_.nodeValue = $1"
	SetNodeValue(nodeValue string)

	// OwnerDocument prop
	// js:"ownerDocument"
	// jsrewrite:"$_.ownerDocument"
	OwnerDocument() (ownerDocument Document)

	// ParentElement prop
	// js:"parentElement"
	// jsrewrite:"$_.parentElement"
	ParentElement() (parentElement HTMLElement)

	// ParentNode prop
	// js:"parentNode"
	// jsrewrite:"$_.parentNode"
	ParentNode() (parentNode Node)

	// PreviousSibling prop
	// js:"previousSibling"
	// jsrewrite:"$_.previousSibling"
	PreviousSibling() (previousSibling Node)

	// TextContent prop
	// js:"textContent"
	// jsrewrite:"$_.textContent"
	TextContent() (textContent string)

	// SetTextContent prop
	// js:"textContent"
	// jsrewrite:"$_.textContent = $1"
	SetTextContent(textContent string)
}
