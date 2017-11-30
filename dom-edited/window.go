package window

import "github.com/matthewmueller/golly/js"

// js:"Event,omit"
type Event interface {
	PreventDefault()
	StopImmediatePropagation()
	StopPropagation()
	Type() (kind string)
}

// js:"EventTarget,omit"
type EventTarget interface {
	AddEventListener(kind string, fn func(e Event), useCapture bool)
	DispatchEvent(evt Event) (b bool)
}

// js:"Node,omit"
type Node interface {
	EventTarget

	ChildNodes() (childNodes []Node)
	FirstChild() (firstChild Node)
	LastChild() (lastChild Node)
	NextSibling() (nextSibling Node)
	NodeName() (nodeName string)
	NodeType() (nodeType uint8)
	NodeValue() (nodeValue string)
	SetNodeValue(nodeValue string)
}

// New fn
func New() *Window {
	js.Rewrite("Window")
	return &Window{}
}

// Window struct
// js:"Window,omit"
type Window struct {
	EventTarget
}

// js:"Document,omit"
type Document interface {
	Node

	QuerySelector(selectors string) (e Element)
	DocumentElement() (documentElement Element)
}

var _ Document = (*HTMLDocument)(nil)

// HTMLDocument struct
// js:"HTMLDocument,omit"
type HTMLDocument struct {
}

// AddEventListener fn
func (h *HTMLDocument) AddEventListener(kind string, fn func(e Event), capture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, fn, capture)
}

// DispatchEvent fn
func (h *HTMLDocument) DispatchEvent(e Event) {

}

func (h *HTMLDocument) QuerySelector(selector string) Element {
	return nil
}

// DocumentElement fn
func (h *HTMLDocument) DocumentElement() (element Element) {
	js.Rewrite("$_.documentElement")
	return element
}

// js:"HTMLElement,omit"
type HTMLElement interface {
	Element

	Blur()
	Focus()
}

// NodeSelector struct
// js:"NodeSelector,omit"
type NodeSelector struct {
}

// QuerySelector fn
func (*NodeSelector) QuerySelector(selectors string) (e Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return e
}

// MouseEvent struct
// js:"MouseEvent,omit"
type MouseEvent struct {
	UIEvent
}

// X prop
func (*MouseEvent) X() (x int) {
	js.Rewrite("$_.x")
	return x
}

// Y prop
func (*MouseEvent) Y() (y int) {
	js.Rewrite("$_.y")
	return y
}

// js:"UIEvent,omit"
type UIEvent interface {
	Event
}

// js:"Element,omit"
type Element interface {
	Node

	// QuerySelector
	// js:"querySelector,rewrite=queryselector"
	QuerySelector(selectors string) (e Element)

	// ClassName prop
	// js:"className,rewrite=classname"
	ClassName() (className string)

	// ClassName prop
	// js:"setclassName,rewrite=setclassname"
	SetClassName(className string)

	// ID prop
	// js:"id,rewrite=id"
	ID() (id string)

	// ID prop
	// js:"setid,rewrite=setid"
	SetID(id string)

	// InnerHTML prop
	// js:"innerHTML,rewrite=innerhtml"
	InnerHTML() (innerHTML string)

	// InnerHTML prop
	// js:"setinnerHTML,rewrite=setinnerhtml"
	SetInnerHTML(innerHTML string)

	// OuterHTML prop
	// js:"outerHTML,rewrite=outerhtml"
	OuterHTML() (outerHTML string)

	// OuterHTML prop
	// js:"setouterHTML,rewrite=setouterhtml"
	SetOuterHTML(outerHTML string)

	// ScrollHeight prop
	// js:"scrollHeight,rewrite=scrollheight"
	ScrollHeight() (scrollHeight int)

	// ScrollLeft prop
	// js:"scrollLeft,rewrite=scrollleft"
	ScrollLeft() (scrollLeft int)

	// ScrollLeft prop
	// js:"setscrollLeft,rewrite=setscrollleft"
	SetScrollLeft(scrollLeft int)

	// ScrollTop prop
	// js:"scrollTop,rewrite=scrolltop"
	ScrollTop() (scrollTop int)

	// ScrollTop prop
	// js:"setscrollTop,rewrite=setscrolltop"
	SetScrollTop(scrollTop int)
}

// classname prop
func classname() (className string) {
	js.Rewrite("$_.className")
	return className
}

// setclassname prop
func setclassname(className string) {
	js.Rewrite("$_.className = className")
}

// id prop
func id() (id string) {
	js.Rewrite("$_.id")
	return id
}

// setid prop
func setid(id string) {
	js.Rewrite("$_.id = id")
}

// innerhtml prop
func innerhtml() (innerHTML string) {
	js.Rewrite("$_.innerHTML")
	return innerHTML
}

// setinnerhtml prop
func setinnerhtml(innerHTML string) {
	js.Rewrite("$_.innerHTML = innerHTML")
}

// outerhtml prop
func outerhtml() (outerHTML string) {
	js.Rewrite("$_.outerHTML")
	return outerHTML
}

// setouterhtml prop
func setouterhtml(outerHTML string) {
	js.Rewrite("$_.outerHTML = outerHTML")
}

// scrollheight prop
func scrollheight() (scrollHeight int) {
	js.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// scrollleft prop
func scrollleft() (scrollLeft int) {
	js.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// setscrollleft prop
func setscrollleft(scrollLeft int) {
	js.Rewrite("$_.scrollLeft = scrollLeft")
}

// scrolltop prop
func scrolltop() (scrollTop int) {
	js.Rewrite("$_.scrollTop")
	return scrollTop
}

// setscrolltop prop
func setscrolltop(scrollTop int) {
	js.Rewrite("$_.scrollTop = scrollTop")
}

// HTMLAnchorElement struct
// js:"HTMLAnchorElement,omit"
type HTMLAnchorElement struct {
	HTMLElement
}

// Href prop Sets or retrieves a destination URL or an anchor point.
func (*HTMLAnchorElement) Href() (href string) {
	js.Rewrite("$_.href")
	return href
}

// Href prop Sets or retrieves a destination URL or an anchor point.
func (*HTMLAnchorElement) SetHref(href string) {
	js.Rewrite("$_.href = $1", href)
}

// preventdefault fn
func preventdefault() {
	js.Rewrite("$_.preventDefault()")
}

// stopimmediatepropagation fn
func stopimmediatepropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// stoppropagation fn
func stoppropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// kind prop
func kind() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
