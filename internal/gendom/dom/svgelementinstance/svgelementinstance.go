package svgelementinstance

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelementinstancelist"
	"github.com/matthewmueller/golly/internal/gendom/dom/svguseelement"
	"github.com/matthewmueller/golly/js"
)

type SVGElementInstance struct {
	*eventtarget.EventTarget
}

func (*SVGElementInstance) GetChildNodes() (childNodes *svgelementinstancelist.SVGElementInstanceList) {
	js.Rewrite("$<.childNodes")
	return childNodes
}

func (*SVGElementInstance) GetCorrespondingElement() (correspondingElement *svgelement.SVGElement) {
	js.Rewrite("$<.correspondingElement")
	return correspondingElement
}

func (*SVGElementInstance) GetCorrespondingUseElement() (correspondingUseElement *svguseelement.SVGUseElement) {
	js.Rewrite("$<.correspondingUseElement")
	return correspondingUseElement
}

func (*SVGElementInstance) GetFirstChild() (firstChild *SVGElementInstance) {
	js.Rewrite("$<.firstChild")
	return firstChild
}

func (*SVGElementInstance) GetLastChild() (lastChild *SVGElementInstance) {
	js.Rewrite("$<.lastChild")
	return lastChild
}

func (*SVGElementInstance) GetNextSibling() (nextSibling *SVGElementInstance) {
	js.Rewrite("$<.nextSibling")
	return nextSibling
}

func (*SVGElementInstance) GetParentNode() (parentNode *SVGElementInstance) {
	js.Rewrite("$<.parentNode")
	return parentNode
}

func (*SVGElementInstance) GetPreviousSibling() (previousSibling *SVGElementInstance) {
	js.Rewrite("$<.previousSibling")
	return previousSibling
}
