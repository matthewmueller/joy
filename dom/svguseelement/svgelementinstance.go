package svguseelement

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.EventTarget = (*SVGElementInstance)(nil)

// SVGElementInstance struct
// js:"SVGElementInstance,omit"
type SVGElementInstance struct {
}

// AddEventListener fn
// js:"addEventListener"
func (*SVGElementInstance) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SVGElementInstance) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SVGElementInstance) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// ChildNodes prop
// js:"childNodes"
func (*SVGElementInstance) ChildNodes() (childNodes *SVGElementInstanceList) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// CorrespondingElement prop
// js:"correspondingElement"
func (*SVGElementInstance) CorrespondingElement() (correspondingElement window.SVGElement) {
	js.Rewrite("$_.correspondingElement")
	return correspondingElement
}

// CorrespondingUseElement prop
// js:"correspondingUseElement"
func (*SVGElementInstance) CorrespondingUseElement() (correspondingUseElement *SVGUseElement) {
	js.Rewrite("$_.correspondingUseElement")
	return correspondingUseElement
}

// FirstChild prop
// js:"firstChild"
func (*SVGElementInstance) FirstChild() (firstChild *SVGElementInstance) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*SVGElementInstance) LastChild() (lastChild *SVGElementInstance) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// NextSibling prop
// js:"nextSibling"
func (*SVGElementInstance) NextSibling() (nextSibling *SVGElementInstance) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// ParentNode prop
// js:"parentNode"
func (*SVGElementInstance) ParentNode() (parentNode *SVGElementInstance) {
	js.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*SVGElementInstance) PreviousSibling() (previousSibling *SVGElementInstance) {
	js.Rewrite("$_.previousSibling")
	return previousSibling
}
