package htmlanchorelement

import (
	"github.com/matthewmueller/golly/dom/element"
	"github.com/matthewmueller/golly/dom/event"
	"github.com/matthewmueller/golly/dom/eventtarget"
	"github.com/matthewmueller/golly/dom/htmlelement"
	"github.com/matthewmueller/golly/dom/node"
	"github.com/matthewmueller/golly/js"
)

var _ htmlelement.HTMLElement = (*HTMLAnchorElement)(nil)
var _ element.Element = (*HTMLAnchorElement)(nil)
var _ element.NodeSelector = (*HTMLAnchorElement)(nil)
var _ node.Node = (*HTMLAnchorElement)(nil)
var _ eventtarget.EventTarget = (*HTMLAnchorElement)(nil)

// HTMLAnchorElement struct
// js:"HTMLAnchorElement,omit"
type HTMLAnchorElement struct {
}

// Blur fn
// js:"blur"
func (*HTMLAnchorElement) Blur() {
	js.Rewrite("$_.blur()")
}

// Focus fn
// js:"focus"
func (*HTMLAnchorElement) Focus() {
	js.Rewrite("$_.focus()")
}

// QuerySelector fn
// js:"querySelector"
func (*HTMLAnchorElement) QuerySelector(selectors string) (e element.Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return e
}

// AddEventListener fn
// js:"addEventListener"
func (*HTMLAnchorElement) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*HTMLAnchorElement) DispatchEvent(evt event.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// Href prop Sets or retrieves a destination URL or an anchor point.
// js:"href"
func (*HTMLAnchorElement) Href() (href string) {
	js.Rewrite("$_.href")
	return href
}

// SetHref prop Sets or retrieves a destination URL or an anchor point.
// js:"href"
func (*HTMLAnchorElement) SetHref(href string) {
	js.Rewrite("$_.href = $1", href)
}

// ClassName prop
// js:"className"
func (*HTMLAnchorElement) ClassName() (className string) {
	js.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*HTMLAnchorElement) SetClassName(className string) {
	js.Rewrite("$_.className = $1", className)
}

// ID prop
// js:"id"
func (*HTMLAnchorElement) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*HTMLAnchorElement) SetID(id string) {
	js.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*HTMLAnchorElement) InnerHTML() (innerHTML string) {
	js.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*HTMLAnchorElement) SetInnerHTML(innerHTML string) {
	js.Rewrite("$_.innerHTML = $1", innerHTML)
}

// OuterHTML prop
// js:"outerHTML"
func (*HTMLAnchorElement) OuterHTML() (outerHTML string) {
	js.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*HTMLAnchorElement) SetOuterHTML(outerHTML string) {
	js.Rewrite("$_.outerHTML = $1", outerHTML)
}

// ScrollHeight prop
// js:"scrollHeight"
func (*HTMLAnchorElement) ScrollHeight() (scrollHeight int) {
	js.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*HTMLAnchorElement) ScrollLeft() (scrollLeft int) {
	js.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*HTMLAnchorElement) SetScrollLeft(scrollLeft int) {
	js.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*HTMLAnchorElement) ScrollTop() (scrollTop int) {
	js.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*HTMLAnchorElement) SetScrollTop(scrollTop int) {
	js.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ChildNodes prop
// js:"childNodes"
func (*HTMLAnchorElement) ChildNodes() (childNodes []node.Node) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*HTMLAnchorElement) FirstChild() (firstChild node.Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*HTMLAnchorElement) LastChild() (lastChild node.Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// NextSibling prop
// js:"nextSibling"
func (*HTMLAnchorElement) NextSibling() (nextSibling node.Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*HTMLAnchorElement) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*HTMLAnchorElement) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*HTMLAnchorElement) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*HTMLAnchorElement) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}
