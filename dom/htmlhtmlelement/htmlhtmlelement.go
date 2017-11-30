package htmlhtmlelement

import (
	"github.com/matthewmueller/golly/dom/element"
	"github.com/matthewmueller/golly/dom/event"
	"github.com/matthewmueller/golly/dom/eventtarget"
	"github.com/matthewmueller/golly/dom/htmlelement"
	"github.com/matthewmueller/golly/dom/node"
	"github.com/matthewmueller/golly/js"
)

var _ htmlelement.HTMLElement = (*HTMLHTMLElement)(nil)
var _ element.Element = (*HTMLHTMLElement)(nil)
var _ element.NodeSelector = (*HTMLHTMLElement)(nil)
var _ node.Node = (*HTMLHTMLElement)(nil)
var _ eventtarget.EventTarget = (*HTMLHTMLElement)(nil)

// HTMLHTMLElement struct
// js:"HTMLHTMLElement,omit"
type HTMLHTMLElement struct {
}

// Blur fn
// js:"blur"
func (*HTMLHTMLElement) Blur() {
	js.Rewrite("$_.blur()")
}

// Focus fn
// js:"focus"
func (*HTMLHTMLElement) Focus() {
	js.Rewrite("$_.focus()")
}

// QuerySelector fn
// js:"querySelector"
func (*HTMLHTMLElement) QuerySelector(selectors string) (e element.Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return e
}

// AddEventListener fn
// js:"addEventListener"
func (*HTMLHTMLElement) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*HTMLHTMLElement) DispatchEvent(evt event.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// ClassName prop
// js:"className"
func (*HTMLHTMLElement) ClassName() (className string) {
	js.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*HTMLHTMLElement) SetClassName(className string) {
	js.Rewrite("$_.className = $1", className)
}

// ID prop
// js:"id"
func (*HTMLHTMLElement) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*HTMLHTMLElement) SetID(id string) {
	js.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*HTMLHTMLElement) InnerHTML() (innerHTML string) {
	js.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*HTMLHTMLElement) SetInnerHTML(innerHTML string) {
	js.Rewrite("$_.innerHTML = $1", innerHTML)
}

// OuterHTML prop
// js:"outerHTML"
func (*HTMLHTMLElement) OuterHTML() (outerHTML string) {
	js.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*HTMLHTMLElement) SetOuterHTML(outerHTML string) {
	js.Rewrite("$_.outerHTML = $1", outerHTML)
}

// ScrollHeight prop
// js:"scrollHeight"
func (*HTMLHTMLElement) ScrollHeight() (scrollHeight int) {
	js.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*HTMLHTMLElement) ScrollLeft() (scrollLeft int) {
	js.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*HTMLHTMLElement) SetScrollLeft(scrollLeft int) {
	js.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*HTMLHTMLElement) ScrollTop() (scrollTop int) {
	js.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*HTMLHTMLElement) SetScrollTop(scrollTop int) {
	js.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ChildNodes prop
// js:"childNodes"
func (*HTMLHTMLElement) ChildNodes() (childNodes []node.Node) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*HTMLHTMLElement) FirstChild() (firstChild node.Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*HTMLHTMLElement) LastChild() (lastChild node.Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// NextSibling prop
// js:"nextSibling"
func (*HTMLHTMLElement) NextSibling() (nextSibling node.Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*HTMLHTMLElement) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*HTMLHTMLElement) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*HTMLHTMLElement) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*HTMLHTMLElement) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}
