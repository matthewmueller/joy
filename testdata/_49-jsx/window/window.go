package window

import "github.com/matthewmueller/golly/js"

// EventTarget interface
type EventTarget interface {
}

// Node interface
type Node interface {
	EventTarget

	ChildNodes() []Node
	FirstChild() Node
	LastChild() Node
	NextSibling() Node
	NodeType() int8
	NodeName() string
	NodeValue() *string
}

// Element interface
type Element interface {
	Node

	ClassName() string
	ID() string
	InnerHTML() string
	OuterHTML() string
	ScrollHeight() int
	ScrollLeft() int
	ScrollTop() int

	QuerySelector(selector string) Element
}

// HTMLElement interface
type HTMLElement interface {
	Element

	Focus()
	Blur()
}

// Event interface
type Event interface {
	PreventDefault()
	StopImmediatePropagation()
	Type() string
}

var _ Event = (*event)(nil)

type event struct {
	kind string
}

func (*event) PreventDefault() {
	js.Rewrite("$<.preventDefault()")
}

func (*event) StopImmediatePropagation() {
	js.Rewrite("$<.stopImmediatePropagation()")
}

func (*event) Type() string {
	js.Rewrite("$<.type")
	return ""
}

var _ EventTarget = (*Window)(nil)

// Window fn
// js:"Window,omit"
type Window struct {
}

// New window
func New() *Window {
	js.Rewrite("window")
	return &Window{}
}

// Document fn
func (w *Window) Document() Document {
	js.Rewrite("$<.document")
	return NewDocument()
}

// DOCUMENT INTERFACE

// Document interface
type Document interface {
	Node

	DocumentElement() Element
	Body() Element
}

func documentElement() Element {
	js.Rewrite("$<.documentElement")
	return NewHTMLHTMLElement()
}

// HTMLDocument struct
var _ Document = (*HTMLDocument)(nil)

// HTMLDocument struct
// js:"HTMLDocument,omit"
type HTMLDocument struct {
	nodeName string
	nodeType int8
	parent   Node
	children []Node
}

// NewDocument fn
func NewDocument() *HTMLDocument {
	js.Rewrite("document")
	return &HTMLDocument{}
}

// DocumentElement fn
func (h *HTMLDocument) DocumentElement() Element {
	js.Rewrite("$<.documentElement")
	return NewHTMLHTMLElement()
}

// Body fn
func (h *HTMLDocument) Body() Element {
	js.Rewrite("$<.body")
	return nil
}

// ChildNodes fn
func (h *HTMLDocument) ChildNodes() (nodes []Node) {
	return h.children
}

// FirstChild fn
func (h *HTMLDocument) FirstChild() Node {
	return h.children[0]
}

// LastChild fn
func (h *HTMLDocument) LastChild() Node {
	return h.children[len(h.children)-1]
}

// NextSibling fn
func (h *HTMLDocument) NextSibling() Node {
	return nil
}

// NodeType fn
func (h *HTMLDocument) NodeType() int8 {
	return 1
}

// NodeName fn
func (h *HTMLDocument) NodeName() string {
	js.Rewrite("$<.nodeName")
	return "#document"
}

// NodeValue fn
func (h *HTMLDocument) NodeValue() *string {
	return nil
}

// HTMLHTMLElement struct
var _ HTMLElement = (*HTMLHTMLElement)(nil)

// HTMLHTMLElement fn
// js:"HTMLHTMLElement,omit"
type HTMLHTMLElement struct {
	nodeName string
	nodeType int8
	parent   Node
	children []Node
}

// NewHTMLHTMLElement fn
// js:"NewHTMLHTMLElement,omit"
func NewHTMLHTMLElement(children ...Node) *HTMLHTMLElement {
	return &HTMLHTMLElement{
		nodeName: "HTML",
		nodeType: 1,
		parent:   nil,
		children: children,
	}
}

// AddEventListener fn
func (h *HTMLHTMLElement) AddEventListener(kind string, fn func(e Event), capture bool) {
	js.Rewrite("$<.addEventListener($1, $2, $3)", kind, fn, capture)
}

// DispatchEvent fn
func (h *HTMLHTMLElement) DispatchEvent(e Event) {

}

// ChildNodes fn
func (h *HTMLHTMLElement) ChildNodes() (nodes []Node) {
	return h.children
}

// FirstChild fn
func (h *HTMLHTMLElement) FirstChild() Node {
	return h.children[0]
}

// LastChild fn
func (h *HTMLHTMLElement) LastChild() Node {
	return h.children[len(h.children)-1]
}

// NextSibling fn
func (h *HTMLHTMLElement) NextSibling() Node {
	return nil
}

// NodeType fn
func (h *HTMLHTMLElement) NodeType() int8 {
	return 1
}

// NodeName fn
func (h *HTMLHTMLElement) NodeName() string {
	js.Rewrite("$<.nodeName")
	return "HTML"
}

// NodeValue fn
func (h *HTMLHTMLElement) NodeValue() *string {
	return nil
}

// ClassName fn
func (h *HTMLHTMLElement) ClassName() string {
	return ""
}

// ID fn
func (h *HTMLHTMLElement) ID() string {
	return ""
}

// InnerHTML fn
func (h *HTMLHTMLElement) InnerHTML() string {
	js.Rewrite("$<.innerHTML")
	return ""
}

// OuterHTML fn
func (h *HTMLHTMLElement) OuterHTML() string {
	return ""
}

// ScrollHeight fn
func (h *HTMLHTMLElement) ScrollHeight() int {
	return 0
}

// ScrollLeft fn
func (h *HTMLHTMLElement) ScrollLeft() int {
	return 0
}

// ScrollTop fn
func (h *HTMLHTMLElement) ScrollTop() int {
	return 0
}

// Focus fn
func (h *HTMLHTMLElement) Focus() {

}

// Blur fn
func (h *HTMLHTMLElement) Blur() {

}

func (h *HTMLHTMLElement) QuerySelector(selector string) Element {
	js.Rewrite("$<.querySelector($1)", selector)
	return NewHTMLAnchorElement()
}

// HTMLHyperlinkElementUtils interface
type HTMLHyperlinkElementUtils interface {
	Href() string
}

// HTMLAnchorElement struct
var _ HTMLElement = (*HTMLAnchorElement)(nil)
var _ HTMLHyperlinkElementUtils = (*HTMLAnchorElement)(nil)

// HTMLAnchorElement fn
// js:"HTMLAnchorElement,omit"
type HTMLAnchorElement struct {
	nodeName string
	nodeType int8
	children []Node
}

// NewHTMLAnchorElement fn
// js:"NewHTMLAnchorElement,omit"
func NewHTMLAnchorElement() *HTMLAnchorElement {
	return &HTMLAnchorElement{
		nodeName: "A",
		nodeType: 1,
	}
}

// AddEventListener fn
func (h *HTMLAnchorElement) AddEventListener(kind string, fn func(e Event), capture bool) {
	js.Rewrite("$<.addEventListener($1, $2, $3)", kind, fn, capture)
}

// DispatchEvent fn
func (h *HTMLAnchorElement) DispatchEvent(e Event) {

}

// ChildNodes fn
func (h *HTMLAnchorElement) ChildNodes() (nodes []Node) {
	return h.children
}

// FirstChild fn
func (h *HTMLAnchorElement) FirstChild() Node {
	return h.children[0]
}

// LastChild fn
func (h *HTMLAnchorElement) LastChild() Node {
	return h.children[len(h.children)-1]
}

// NextSibling fn
func (h *HTMLAnchorElement) NextSibling() Node {
	return nil
}

// NodeType fn
func (h *HTMLAnchorElement) NodeType() int8 {
	return 1
}

// NodeName fn
func (h *HTMLAnchorElement) NodeName() string {
	js.Rewrite("$<.nodeName")
	return "HTML"
}

// NodeValue fn
func (h *HTMLAnchorElement) NodeValue() *string {
	return nil
}

// ClassName fn
func (h *HTMLAnchorElement) ClassName() string {
	return ""
}

// ID fn
func (h *HTMLAnchorElement) ID() string {
	return ""
}

// InnerHTML fn
func (h *HTMLAnchorElement) InnerHTML() string {
	js.Rewrite("$<.innerHTML")
	return ""
}

// OuterHTML fn
func (h *HTMLAnchorElement) OuterHTML() string {
	return ""
}

// ScrollHeight fn
func (h *HTMLAnchorElement) ScrollHeight() int {
	return 0
}

// ScrollLeft fn
func (h *HTMLAnchorElement) ScrollLeft() int {
	return 0
}

// ScrollTop fn
func (h *HTMLAnchorElement) ScrollTop() int {
	return 0
}

// Focus fn
func (h *HTMLAnchorElement) Focus() {

}

// Blur fn
func (h *HTMLAnchorElement) Blur() {

}

// QuerySelector fn
func (h *HTMLAnchorElement) QuerySelector(selector string) Element {
	js.Rewrite("$<.querySelector($1)", selector)
	return nil
}

// Href string
func (H *HTMLAnchorElement) Href() string {
	js.Rewrite("$<.href")
	return "www.google.com"
}
