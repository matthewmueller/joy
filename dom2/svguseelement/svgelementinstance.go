package svguseelement

import "github.com/matthewmueller/golly/js"

// js:"SVGElementInstance,omit"
type SVGElementInstance struct {
	window.EventTarget
}

// ChildNodes
func (*SVGElementInstance) ChildNodes() (childNodes *SVGElementInstanceList) {
	js.Rewrite("$<.ChildNodes")
	return childNodes
}

// CorrespondingElement
func (*SVGElementInstance) CorrespondingElement() (correspondingElement window.SVGElement) {
	js.Rewrite("$<.CorrespondingElement")
	return correspondingElement
}

// CorrespondingUseElement
func (*SVGElementInstance) CorrespondingUseElement() (correspondingUseElement *SVGUseElement) {
	js.Rewrite("$<.CorrespondingUseElement")
	return correspondingUseElement
}

// FirstChild
func (*SVGElementInstance) FirstChild() (firstChild *SVGElementInstance) {
	js.Rewrite("$<.FirstChild")
	return firstChild
}

// LastChild
func (*SVGElementInstance) LastChild() (lastChild *SVGElementInstance) {
	js.Rewrite("$<.LastChild")
	return lastChild
}

// NextSibling
func (*SVGElementInstance) NextSibling() (nextSibling *SVGElementInstance) {
	js.Rewrite("$<.NextSibling")
	return nextSibling
}

// ParentNode
func (*SVGElementInstance) ParentNode() (parentNode *SVGElementInstance) {
	js.Rewrite("$<.ParentNode")
	return parentNode
}

// PreviousSibling
func (*SVGElementInstance) PreviousSibling() (previousSibling *SVGElementInstance) {
	js.Rewrite("$<.PreviousSibling")
	return previousSibling
}
