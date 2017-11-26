package window

import "github.com/matthewmueller/golly/js"

// js:"Node,omit"
type Node interface {
	EventTarget

	// AppendChild
	// js:"appendChild,rewrite=appendchild"
	AppendChild(newChild Node) (n Node)

	// CloneNode
	// js:"cloneNode,rewrite=clonenode"
	CloneNode(deep *bool) (n Node)

	// CompareDocumentPosition
	// js:"compareDocumentPosition,rewrite=comparedocumentposition"
	CompareDocumentPosition(other Node) (u uint8)

	// Contains
	// js:"contains,rewrite=contains"
	Contains(child Node) (b bool)

	// HasAttributes
	// js:"hasAttributes,rewrite=hasattributes"
	HasAttributes() (b bool)

	// HasChildNodes
	// js:"hasChildNodes,rewrite=haschildnodes"
	HasChildNodes() (b bool)

	// InsertBefore
	// js:"insertBefore,rewrite=insertbefore"
	InsertBefore(newChild Node, refChild *Node) (n Node)

	// IsDefaultNamespace
	// js:"isDefaultNamespace,rewrite=isdefaultnamespace"
	IsDefaultNamespace(namespaceURI string) (b bool)

	// IsEqualNode
	// js:"isEqualNode,rewrite=isequalnode"
	IsEqualNode(arg Node) (b bool)

	// IsSameNode
	// js:"isSameNode,rewrite=issamenode"
	IsSameNode(other Node) (b bool)

	// LookupNamespaceURI
	// js:"lookupNamespaceURI,rewrite=lookupnamespaceuri"
	LookupNamespaceURI(prefix string) (s string)

	// LookupPrefix
	// js:"lookupPrefix,rewrite=lookupprefix"
	LookupPrefix(namespaceURI string) (s string)

	// Normalize
	// js:"normalize,rewrite=normalize"
	Normalize()

	// RemoveChild
	// js:"removeChild,rewrite=removechild"
	RemoveChild(oldChild Node) (n Node)

	// ReplaceChild
	// js:"replaceChild,rewrite=replacechild"
	ReplaceChild(newChild Node, oldChild Node) (n Node)

	// Attributes prop
	// js:"attributes,rewrite=attributes"
	Attributes() (attributes *NamedNodeMap)

	// BaseURI prop
	// js:"baseURI,rewrite=baseuri"
	BaseURI() (baseURI string)

	// ChildNodes prop
	// js:"childNodes,rewrite=childnodes"
	ChildNodes() (childNodes *NodeList)

	// FirstChild prop
	// js:"firstChild,rewrite=firstchild"
	FirstChild() (firstChild Node)

	// LastChild prop
	// js:"lastChild,rewrite=lastchild"
	LastChild() (lastChild Node)

	// LocalName prop
	// js:"localName,rewrite=localname"
	LocalName() (localName string)

	// NamespaceURI prop
	// js:"namespaceURI,rewrite=namespaceuri"
	NamespaceURI() (namespaceURI string)

	// NextSibling prop
	// js:"nextSibling,rewrite=nextsibling"
	NextSibling() (nextSibling Node)

	// NodeName prop
	// js:"nodeName,rewrite=nodename"
	NodeName() (nodeName string)

	// NodeType prop
	// js:"nodeType,rewrite=nodetype"
	NodeType() (nodeType uint8)

	// NodeValue prop
	// js:"nodeValue,rewrite=nodevalue"
	NodeValue() (nodeValue string)

	// NodeValue prop
	// js:"setnodeValue,rewrite=setnodevalue"
	SetNodeValue(nodeValue string)

	// OwnerDocument prop
	// js:"ownerDocument,rewrite=ownerdocument"
	OwnerDocument() (ownerDocument Document)

	// ParentElement prop
	// js:"parentElement,rewrite=parentelement"
	ParentElement() (parentElement HTMLElement)

	// ParentNode prop
	// js:"parentNode,rewrite=parentnode"
	ParentNode() (parentNode Node)

	// PreviousSibling prop
	// js:"previousSibling,rewrite=previoussibling"
	PreviousSibling() (previousSibling Node)

	// TextContent prop
	// js:"textContent,rewrite=textcontent"
	TextContent() (textContent string)

	// TextContent prop
	// js:"settextContent,rewrite=settextcontent"
	SetTextContent(textContent string)
}

// appendchild fn
func appendchild(newChild Node) (n Node) {
	js.Rewrite("$<.appendChild($1)", newChild)
	return n
}

// clonenode fn
func clonenode(deep *bool) (n Node) {
	js.Rewrite("$<.cloneNode($1)", deep)
	return n
}

// comparedocumentposition fn
func comparedocumentposition(other Node) (u uint8) {
	js.Rewrite("$<.compareDocumentPosition($1)", other)
	return u
}

// contains fn
func contains(child Node) (b bool) {
	js.Rewrite("$<.contains($1)", child)
	return b
}

// hasattributes fn
func hasattributes() (b bool) {
	js.Rewrite("$<.hasAttributes()")
	return b
}

// haschildnodes fn
func haschildnodes() (b bool) {
	js.Rewrite("$<.hasChildNodes()")
	return b
}

// insertbefore fn
func insertbefore(newChild Node, refChild *Node) (n Node) {
	js.Rewrite("$<.insertBefore($1, $2)", newChild, refChild)
	return n
}

// isdefaultnamespace fn
func isdefaultnamespace(namespaceURI string) (b bool) {
	js.Rewrite("$<.isDefaultNamespace($1)", namespaceURI)
	return b
}

// isequalnode fn
func isequalnode(arg Node) (b bool) {
	js.Rewrite("$<.isEqualNode($1)", arg)
	return b
}

// issamenode fn
func issamenode(other Node) (b bool) {
	js.Rewrite("$<.isSameNode($1)", other)
	return b
}

// lookupnamespaceuri fn
func lookupnamespaceuri(prefix string) (s string) {
	js.Rewrite("$<.lookupNamespaceURI($1)", prefix)
	return s
}

// lookupprefix fn
func lookupprefix(namespaceURI string) (s string) {
	js.Rewrite("$<.lookupPrefix($1)", namespaceURI)
	return s
}

// normalize fn
func normalize() {
	js.Rewrite("$<.normalize()")
}

// removechild fn
func removechild(oldChild Node) (n Node) {
	js.Rewrite("$<.removeChild($1)", oldChild)
	return n
}

// replacechild fn
func replacechild(newChild Node, oldChild Node) (n Node) {
	js.Rewrite("$<.replaceChild($1, $2)", newChild, oldChild)
	return n
}

// attributes prop
func attributes() (attributes *NamedNodeMap) {
	js.Rewrite("$<.attributes")
	return attributes
}

// baseuri prop
func baseuri() (baseURI string) {
	js.Rewrite("$<.baseURI")
	return baseURI
}

// childnodes prop
func childnodes() (childNodes *NodeList) {
	js.Rewrite("$<.childNodes")
	return childNodes
}

// firstchild prop
func firstchild() (firstChild Node) {
	js.Rewrite("$<.firstChild")
	return firstChild
}

// lastchild prop
func lastchild() (lastChild Node) {
	js.Rewrite("$<.lastChild")
	return lastChild
}

// localname prop
func localname() (localName string) {
	js.Rewrite("$<.localName")
	return localName
}

// namespaceuri prop
func namespaceuri() (namespaceURI string) {
	js.Rewrite("$<.namespaceURI")
	return namespaceURI
}

// nextsibling prop
func nextsibling() (nextSibling Node) {
	js.Rewrite("$<.nextSibling")
	return nextSibling
}

// nodename prop
func nodename() (nodeName string) {
	js.Rewrite("$<.nodeName")
	return nodeName
}

// nodetype prop
func nodetype() (nodeType uint8) {
	js.Rewrite("$<.nodeType")
	return nodeType
}

// nodevalue prop
func nodevalue() (nodeValue string) {
	js.Rewrite("$<.nodeValue")
	return nodeValue
}

// setnodevalue prop
func setnodevalue(nodeValue string) {
	js.Rewrite("$<.nodeValue = nodeValue")
}

// ownerdocument prop
func ownerdocument() (ownerDocument Document) {
	js.Rewrite("$<.ownerDocument")
	return ownerDocument
}

// parentelement prop
func parentelement() (parentElement HTMLElement) {
	js.Rewrite("$<.parentElement")
	return parentElement
}

// parentnode prop
func parentnode() (parentNode Node) {
	js.Rewrite("$<.parentNode")
	return parentNode
}

// previoussibling prop
func previoussibling() (previousSibling Node) {
	js.Rewrite("$<.previousSibling")
	return previousSibling
}

// textcontent prop
func textcontent() (textContent string) {
	js.Rewrite("$<.textContent")
	return textContent
}

// settextcontent prop
func settextcontent(textContent string) {
	js.Rewrite("$<.textContent = textContent")
}
