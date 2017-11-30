package dom

import "github.com/matthewmueller/golly/js"

// UIEvent interface
// js:"UIEvent"
type UIEvent interface {
	Event
}

var _ UIEvent = (*MouseEvent)(nil)
var _ Event = (*MouseEvent)(nil)

// MouseEvent struct
// js:"MouseEvent,omit"
type MouseEvent struct {
}

// PreventDefault fn
// js:"preventDefault"
func (*MouseEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MouseEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MouseEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// X prop
// js:"x"
func (*MouseEvent) X() (x int) {
	js.Rewrite("$_.x")
	return x
}

// Y prop
// js:"y"
func (*MouseEvent) Y() (y int) {
	js.Rewrite("$_.y")
	return y
}

// Type prop
// js:"type"
func (*MouseEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}

// EventTarget interface
// js:"EventTarget"
type EventTarget interface {

	// AddEventListener
	// js:"addEventListener"
	AddEventListener(kind string, listener func(evt Event), useCapture bool)

	// DispatchEvent
	// js:"dispatchEvent"
	DispatchEvent(evt Event) (b bool)
}

// HTMLElement interface
// js:"HTMLElement"
type HTMLElement interface {
	Element

	// Blur
	// js:"blur"
	Blur()

	// Focus
	// js:"focus"
	Focus()
}

// Element interface
// js:"Element"
type Element interface {
	Node

	// QuerySelector
	// js:"querySelector"
	QuerySelector(selectors string) (e Element)

	// ClassName prop
	// js:"className"
	ClassName() (className string)

	// SetClassName prop
	// js:"className"
	SetClassName(className string)

	// ID prop
	// js:"id"
	ID() (id string)

	// SetID prop
	// js:"id"
	SetID(id string)

	// InnerHTML prop
	// js:"innerHTML"
	InnerHTML() (innerHTML string)

	// SetInnerHTML prop
	// js:"innerHTML"
	SetInnerHTML(innerHTML string)

	// OuterHTML prop
	// js:"outerHTML"
	OuterHTML() (outerHTML string)

	// SetOuterHTML prop
	// js:"outerHTML"
	SetOuterHTML(outerHTML string)

	// ScrollHeight prop
	// js:"scrollHeight"
	ScrollHeight() (scrollHeight int)

	// ScrollLeft prop
	// js:"scrollLeft"
	ScrollLeft() (scrollLeft int)

	// SetScrollLeft prop
	// js:"scrollLeft"
	SetScrollLeft(scrollLeft int)

	// ScrollTop prop
	// js:"scrollTop"
	ScrollTop() (scrollTop int)

	// SetScrollTop prop
	// js:"scrollTop"
	SetScrollTop(scrollTop int)
}

var _ HTMLElement = (*HTMLAnchorElement)(nil)
var _ Element = (*HTMLAnchorElement)(nil)
var _ NodeSelector = (*HTMLAnchorElement)(nil)
var _ Node = (*HTMLAnchorElement)(nil)
var _ EventTarget = (*HTMLAnchorElement)(nil)

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
func (*HTMLAnchorElement) QuerySelector(selectors string) (e Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return e
}

// AddEventListener fn
// js:"addEventListener"
func (*HTMLAnchorElement) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*HTMLAnchorElement) DispatchEvent(evt Event) (b bool) {
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
func (*HTMLAnchorElement) ChildNodes() (childNodes []Node) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*HTMLAnchorElement) FirstChild() (firstChild Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*HTMLAnchorElement) LastChild() (lastChild Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// NextSibling prop
// js:"nextSibling"
func (*HTMLAnchorElement) NextSibling() (nextSibling Node) {
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

// NodeSelector interface
// js:"NodeSelector"
type NodeSelector interface {

	// QuerySelector
	// js:"querySelector"
	QuerySelector(selectors string) (e Element)
}

// Event interface
// js:"Event"
type Event interface {

	// PreventDefault
	// js:"preventDefault"
	PreventDefault()

	// StopImmediatePropagation
	// js:"stopImmediatePropagation"
	StopImmediatePropagation()

	// StopPropagation
	// js:"stopPropagation"
	StopPropagation()

	// Type prop
	// js:"type"
	Type() (kind string)
}

// Node interface
// js:"Node"
type Node interface {
	EventTarget

	// ChildNodes prop
	// js:"childNodes"
	ChildNodes() (childNodes []Node)

	// FirstChild prop
	// js:"firstChild"
	FirstChild() (firstChild Node)

	// LastChild prop
	// js:"lastChild"
	LastChild() (lastChild Node)

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
}

// New fn
func New() *Window {
	js.Rewrite("window")
	return &Window{}
}

var _ EventTarget = (*Window)(nil)

// Window struct
// js:"Window,omit"
type Window struct {
}

// AddEventListener fn
// js:"addEventListener"
func (*Window) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*Window) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// Document prop
// js:"document"
func (*Window) Document() (document Document) {
	js.Rewrite("$_.document")
	return document
}

// Document interface
// js:"Document"
type Document interface {
	Node

	// QuerySelector
	// js:"querySelector"
	QuerySelector(selectors string) (e Element)

	// DocumentElement prop Gets a reference to the root node of the document.
	// js:"documentElement"
	DocumentElement() (documentElement Element)
}

var _ Document = (*HTMLDocument)(nil)
var _ NodeSelector = (*HTMLDocument)(nil)
var _ Node = (*HTMLDocument)(nil)
var _ EventTarget = (*HTMLDocument)(nil)

// HTMLDocument struct
// js:"HTMLDocument,omit"
type HTMLDocument struct {
}

// QuerySelector fn
// js:"querySelector"
func (*HTMLDocument) QuerySelector(selectors string) (e Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return e
}

// AddEventListener fn
// js:"addEventListener"
func (*HTMLDocument) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*HTMLDocument) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// DocumentElement prop Gets a reference to the root node of the document.
// js:"documentElement"
func (*HTMLDocument) DocumentElement() (documentElement Element) {
	js.Rewrite("$_.documentElement")
	return documentElement
}

// ChildNodes prop
// js:"childNodes"
func (*HTMLDocument) ChildNodes() (childNodes []Node) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*HTMLDocument) FirstChild() (firstChild Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*HTMLDocument) LastChild() (lastChild Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// NextSibling prop
// js:"nextSibling"
func (*HTMLDocument) NextSibling() (nextSibling Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*HTMLDocument) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*HTMLDocument) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*HTMLDocument) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*HTMLDocument) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}
