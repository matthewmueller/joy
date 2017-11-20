package svgelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssstyledeclaration"
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/focusevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/mouseevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedstring"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgsvgelement"
	"github.com/matthewmueller/golly/js"
)

type SVGElement struct {
	*element.Element
}

func (*SVGElement) GetClassName() (className *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.className")
	return className
}

func (*SVGElement) GetOnclick() (click *mouseevent.MouseEvent) {
	js.Rewrite("$<.onclick")
	return click
}

func (*SVGElement) SetOnclick(click *mouseevent.MouseEvent) {
	js.Rewrite("$<.onclick = $1", click)
}

func (*SVGElement) GetOndblclick() (dblclick *mouseevent.MouseEvent) {
	js.Rewrite("$<.ondblclick")
	return dblclick
}

func (*SVGElement) SetOndblclick(dblclick *mouseevent.MouseEvent) {
	js.Rewrite("$<.ondblclick = $1", dblclick)
}

func (*SVGElement) GetOnfocusin() (focusin *focusevent.FocusEvent) {
	js.Rewrite("$<.onfocusin")
	return focusin
}

func (*SVGElement) SetOnfocusin(focusin *focusevent.FocusEvent) {
	js.Rewrite("$<.onfocusin = $1", focusin)
}

func (*SVGElement) GetOnfocusout() (focusout *focusevent.FocusEvent) {
	js.Rewrite("$<.onfocusout")
	return focusout
}

func (*SVGElement) SetOnfocusout(focusout *focusevent.FocusEvent) {
	js.Rewrite("$<.onfocusout = $1", focusout)
}

func (*SVGElement) GetOnload() (e *event.Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*SVGElement) SetOnload(e *event.Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*SVGElement) GetOnmousedown() (mousedown *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmousedown")
	return mousedown
}

func (*SVGElement) SetOnmousedown(mousedown *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmousedown = $1", mousedown)
}

func (*SVGElement) GetOnmousemove() (mousemove *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmousemove")
	return mousemove
}

func (*SVGElement) SetOnmousemove(mousemove *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmousemove = $1", mousemove)
}

func (*SVGElement) GetOnmouseout() (mouseout *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseout")
	return mouseout
}

func (*SVGElement) SetOnmouseout(mouseout *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseout = $1", mouseout)
}

func (*SVGElement) GetOnmouseover() (mouseover *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseover")
	return mouseover
}

func (*SVGElement) SetOnmouseover(mouseover *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseover = $1", mouseover)
}

func (*SVGElement) GetOnmouseup() (mouseup *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseup")
	return mouseup
}

func (*SVGElement) SetOnmouseup(mouseup *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseup = $1", mouseup)
}

func (*SVGElement) GetOwnerSVGElement() (ownerSVGElement *svgsvgelement.SVGSVGElement) {
	js.Rewrite("$<.ownerSVGElement")
	return ownerSVGElement
}

func (*SVGElement) GetStyle() (style *cssstyledeclaration.CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}

func (*SVGElement) GetViewportElement() (viewportElement *SVGElement) {
	js.Rewrite("$<.viewportElement")
	return viewportElement
}

func (*SVGElement) GetXmlbase() (xmlbase string) {
	js.Rewrite("$<.xmlbase")
	return xmlbase
}

func (*SVGElement) SetXmlbase(xmlbase string) {
	js.Rewrite("$<.xmlbase = $1", xmlbase)
}
