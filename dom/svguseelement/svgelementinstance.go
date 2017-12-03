package svguseelement

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.EventTarget = (*SVGElementInstance)(nil)

// SVGElementInstance struct
// js:"SVGElementInstance,omit"
type SVGElementInstance struct {
}

// AddEventListener fn
// js:"addEventListener"
func (*SVGElementInstance) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SVGElementInstance) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SVGElementInstance) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// ChildNodes prop
// js:"childNodes"
func (*SVGElementInstance) ChildNodes() (childNodes *SVGElementInstanceList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

// CorrespondingElement prop
// js:"correspondingElement"
func (*SVGElementInstance) CorrespondingElement() (correspondingElement window.SVGElement) {
	macro.Rewrite("$_.correspondingElement")
	return correspondingElement
}

// CorrespondingUseElement prop
// js:"correspondingUseElement"
func (*SVGElementInstance) CorrespondingUseElement() (correspondingUseElement *SVGUseElement) {
	macro.Rewrite("$_.correspondingUseElement")
	return correspondingUseElement
}

// FirstChild prop
// js:"firstChild"
func (*SVGElementInstance) FirstChild() (firstChild *SVGElementInstance) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*SVGElementInstance) LastChild() (lastChild *SVGElementInstance) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

// NextSibling prop
// js:"nextSibling"
func (*SVGElementInstance) NextSibling() (nextSibling *SVGElementInstance) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

// ParentNode prop
// js:"parentNode"
func (*SVGElementInstance) ParentNode() (parentNode *SVGElementInstance) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*SVGElementInstance) PreviousSibling() (previousSibling *SVGElementInstance) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}
