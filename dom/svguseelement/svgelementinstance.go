package svguseelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGElementInstance struct
// js:"SVGElementInstance,omit"
type SVGElementInstance struct {
	window.EventTarget
}

// ChildNodes prop
func (*SVGElementInstance) ChildNodes() (childNodes *SVGElementInstanceList) {
	js.Rewrite("$<.childNodes")
	return childNodes
}

// CorrespondingElement prop
func (*SVGElementInstance) CorrespondingElement() (correspondingElement window.SVGElement) {
	js.Rewrite("$<.correspondingElement")
	return correspondingElement
}

// CorrespondingUseElement prop
func (*SVGElementInstance) CorrespondingUseElement() (correspondingUseElement *SVGUseElement) {
	js.Rewrite("$<.correspondingUseElement")
	return correspondingUseElement
}

// FirstChild prop
func (*SVGElementInstance) FirstChild() (firstChild *SVGElementInstance) {
	js.Rewrite("$<.firstChild")
	return firstChild
}

// LastChild prop
func (*SVGElementInstance) LastChild() (lastChild *SVGElementInstance) {
	js.Rewrite("$<.lastChild")
	return lastChild
}

// NextSibling prop
func (*SVGElementInstance) NextSibling() (nextSibling *SVGElementInstance) {
	js.Rewrite("$<.nextSibling")
	return nextSibling
}

// ParentNode prop
func (*SVGElementInstance) ParentNode() (parentNode *SVGElementInstance) {
	js.Rewrite("$<.parentNode")
	return parentNode
}

// PreviousSibling prop
func (*SVGElementInstance) PreviousSibling() (previousSibling *SVGElementInstance) {
	js.Rewrite("$<.previousSibling")
	return previousSibling
}
