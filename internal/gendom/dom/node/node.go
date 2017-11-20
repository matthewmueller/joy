package node

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/document"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/namednodemap"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodelist"
	"github.com/matthewmueller/golly/js"
)

type Node struct {
	*eventtarget.EventTarget
}

func (*Node) AppendChild(newChild *Node) (n *Node) {
	js.Rewrite("$<.appendChild($1)", newChild)
	return n
}

func (*Node) CloneNode(deep bool) (n *Node) {
	js.Rewrite("$<.cloneNode($1)", deep)
	return n
}

func (*Node) CompareDocumentPosition(other *Node) (u uint8) {
	js.Rewrite("$<.compareDocumentPosition($1)", other)
	return u
}

func (*Node) Contains(child *Node) (b bool) {
	js.Rewrite("$<.contains($1)", child)
	return b
}

func (*Node) HasAttributes() (b bool) {
	js.Rewrite("$<.hasAttributes()")
	return b
}

func (*Node) HasChildNodes() (b bool) {
	js.Rewrite("$<.hasChildNodes()")
	return b
}

func (*Node) InsertBefore(newChild *Node, refChild *Node) (n *Node) {
	js.Rewrite("$<.insertBefore($1, $2)", newChild, refChild)
	return n
}

func (*Node) IsDefaultNamespace(namespaceURI string) (b bool) {
	js.Rewrite("$<.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*Node) IsEqualNode(arg *Node) (b bool) {
	js.Rewrite("$<.isEqualNode($1)", arg)
	return b
}

func (*Node) IsSameNode(other *Node) (b bool) {
	js.Rewrite("$<.isSameNode($1)", other)
	return b
}

func (*Node) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$<.lookupNamespaceURI($1)", prefix)
	return s
}

func (*Node) LookupPrefix(namespaceURI string) (s string) {
	js.Rewrite("$<.lookupPrefix($1)", namespaceURI)
	return s
}

func (*Node) Normalize() {
	js.Rewrite("$<.normalize()")
}

func (*Node) RemoveChild(oldChild *Node) (n *Node) {
	js.Rewrite("$<.removeChild($1)", oldChild)
	return n
}

func (*Node) ReplaceChild(newChild *Node, oldChild *Node) (n *Node) {
	js.Rewrite("$<.replaceChild($1, $2)", newChild, oldChild)
	return n
}

func (*Node) GetAttributes() (attributes *namednodemap.NamedNodeMap) {
	js.Rewrite("$<.attributes")
	return attributes
}

func (*Node) GetBaseURI() (baseURI string) {
	js.Rewrite("$<.baseURI")
	return baseURI
}

func (*Node) GetChildNodes() (childNodes *nodelist.NodeList) {
	js.Rewrite("$<.childNodes")
	return childNodes
}

func (*Node) GetFirstChild() (firstChild *Node) {
	js.Rewrite("$<.firstChild")
	return firstChild
}

func (*Node) GetLastChild() (lastChild *Node) {
	js.Rewrite("$<.lastChild")
	return lastChild
}

func (*Node) GetLocalName() (localName string) {
	js.Rewrite("$<.localName")
	return localName
}

func (*Node) GetNamespaceURI() (namespaceURI string) {
	js.Rewrite("$<.namespaceURI")
	return namespaceURI
}

func (*Node) GetNextSibling() (nextSibling *Node) {
	js.Rewrite("$<.nextSibling")
	return nextSibling
}

func (*Node) GetNodeName() (nodeName string) {
	js.Rewrite("$<.nodeName")
	return nodeName
}

func (*Node) GetNodeType() (nodeType uint8) {
	js.Rewrite("$<.nodeType")
	return nodeType
}

func (*Node) GetNodeValue() (nodeValue string) {
	js.Rewrite("$<.nodeValue")
	return nodeValue
}

func (*Node) SetNodeValue(nodeValue string) {
	js.Rewrite("$<.nodeValue = $1", nodeValue)
}

func (*Node) GetOwnerDocument() (ownerDocument *document.Document) {
	js.Rewrite("$<.ownerDocument")
	return ownerDocument
}

func (*Node) GetParentElement() (parentElement *htmlelement.HTMLElement) {
	js.Rewrite("$<.parentElement")
	return parentElement
}

func (*Node) GetParentNode() (parentNode *Node) {
	js.Rewrite("$<.parentNode")
	return parentNode
}

func (*Node) GetPreviousSibling() (previousSibling *Node) {
	js.Rewrite("$<.previousSibling")
	return previousSibling
}

func (*Node) GetTextContent() (textContent string) {
	js.Rewrite("$<.textContent")
	return textContent
}

func (*Node) SetTextContent(textContent string) {
	js.Rewrite("$<.textContent = $1", textContent)
}
