package svgsvgelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssstyledeclaration"
	"github.com/matthewmueller/golly/internal/gendom/dom/documentevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodelist"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgangle"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfittoviewbox"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svglength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgmatrix"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgnumber"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpoint"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgrect"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgtransform"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgzoomandpan"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgzoomevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/js"
)

type SVGSVGElement struct {
	*svggraphicselement.SVGGraphicsElement
	*documentevent.DocumentEvent
	*svgfittoviewbox.SVGFitToViewBox
	*svgzoomandpan.SVGZoomAndPan
}

func (*SVGSVGElement) CheckEnclosure(element *svgelement.SVGElement, rect *svgrect.SVGRect) (b bool) {
	js.Rewrite("$<.checkEnclosure($1, $2)", element, rect)
	return b
}

func (*SVGSVGElement) CheckIntersection(element *svgelement.SVGElement, rect *svgrect.SVGRect) (b bool) {
	js.Rewrite("$<.checkIntersection($1, $2)", element, rect)
	return b
}

func (*SVGSVGElement) CreateSVGAngle() (s *svgangle.SVGAngle) {
	js.Rewrite("$<.createSVGAngle()")
	return s
}

func (*SVGSVGElement) CreateSVGLength() (s *svglength.SVGLength) {
	js.Rewrite("$<.createSVGLength()")
	return s
}

func (*SVGSVGElement) CreateSVGMatrix() (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.createSVGMatrix()")
	return s
}

func (*SVGSVGElement) CreateSVGNumber() (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.createSVGNumber()")
	return s
}

func (*SVGSVGElement) CreateSVGPoint() (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.createSVGPoint()")
	return s
}

func (*SVGSVGElement) CreateSVGRect() (s *svgrect.SVGRect) {
	js.Rewrite("$<.createSVGRect()")
	return s
}

func (*SVGSVGElement) CreateSVGTransform() (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.createSVGTransform()")
	return s
}

func (*SVGSVGElement) CreateSVGTransformFromMatrix(matrix *svgmatrix.SVGMatrix) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.createSVGTransformFromMatrix($1)", matrix)
	return s
}

func (*SVGSVGElement) DeselectAll() {
	js.Rewrite("$<.deselectAll()")
}

func (*SVGSVGElement) ForceRedraw() {
	js.Rewrite("$<.forceRedraw()")
}

func (*SVGSVGElement) GetComputedStyle(elt *element.Element, pseudoElt string) (c *cssstyledeclaration.CSSStyleDeclaration) {
	js.Rewrite("$<.getComputedStyle($1, $2)", elt, pseudoElt)
	return c
}

func (*SVGSVGElement) GetCurrentTime() (f float32) {
	js.Rewrite("$<.getCurrentTime()")
	return f
}

func (*SVGSVGElement) GetElementByID(elementId string) (e *element.Element) {
	js.Rewrite("$<.getElementById($1)", elementId)
	return e
}

func (*SVGSVGElement) GetEnclosureList(rect *svgrect.SVGRect, referenceElement *svgelement.SVGElement) (n *nodelist.NodeList) {
	js.Rewrite("$<.getEnclosureList($1, $2)", rect, referenceElement)
	return n
}

func (*SVGSVGElement) GetIntersectionList(rect *svgrect.SVGRect, referenceElement *svgelement.SVGElement) (n *nodelist.NodeList) {
	js.Rewrite("$<.getIntersectionList($1, $2)", rect, referenceElement)
	return n
}

func (*SVGSVGElement) PauseAnimations() {
	js.Rewrite("$<.pauseAnimations()")
}

func (*SVGSVGElement) SetCurrentTime(seconds float32) {
	js.Rewrite("$<.setCurrentTime($1)", seconds)
}

func (*SVGSVGElement) SuspendRedraw(maxWaitMilliseconds uint) (u uint) {
	js.Rewrite("$<.suspendRedraw($1)", maxWaitMilliseconds)
	return u
}

func (*SVGSVGElement) UnpauseAnimations() {
	js.Rewrite("$<.unpauseAnimations()")
}

func (*SVGSVGElement) UnsuspendRedraw(suspendHandleID uint) {
	js.Rewrite("$<.unsuspendRedraw($1)", suspendHandleID)
}

func (*SVGSVGElement) UnsuspendRedrawAll() {
	js.Rewrite("$<.unsuspendRedrawAll()")
}

func (*SVGSVGElement) GetContentScriptType() (contentScriptType string) {
	js.Rewrite("$<.contentScriptType")
	return contentScriptType
}

func (*SVGSVGElement) SetContentScriptType(contentScriptType string) {
	js.Rewrite("$<.contentScriptType = $1", contentScriptType)
}

func (*SVGSVGElement) GetContentStyleType() (contentStyleType string) {
	js.Rewrite("$<.contentStyleType")
	return contentStyleType
}

func (*SVGSVGElement) SetContentStyleType(contentStyleType string) {
	js.Rewrite("$<.contentStyleType = $1", contentStyleType)
}

func (*SVGSVGElement) GetCurrentScale() (currentScale float32) {
	js.Rewrite("$<.currentScale")
	return currentScale
}

func (*SVGSVGElement) SetCurrentScale(currentScale float32) {
	js.Rewrite("$<.currentScale = $1", currentScale)
}

func (*SVGSVGElement) GetCurrentTranslate() (currentTranslate *svgpoint.SVGPoint) {
	js.Rewrite("$<.currentTranslate")
	return currentTranslate
}

func (*SVGSVGElement) GetHeight() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGSVGElement) GetOnabort() (e *event.Event) {
	js.Rewrite("$<.onabort")
	return e
}

func (*SVGSVGElement) SetOnabort(e *event.Event) {
	js.Rewrite("$<.onabort = $1", e)
}

func (*SVGSVGElement) GetOnerror() (e *event.Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*SVGSVGElement) SetOnerror(e *event.Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*SVGSVGElement) GetOnresize() (e *event.Event) {
	js.Rewrite("$<.onresize")
	return e
}

func (*SVGSVGElement) SetOnresize(e *event.Event) {
	js.Rewrite("$<.onresize = $1", e)
}

func (*SVGSVGElement) GetOnscroll() (scroll *uievent.UIEvent) {
	js.Rewrite("$<.onscroll")
	return scroll
}

func (*SVGSVGElement) SetOnscroll(scroll *uievent.UIEvent) {
	js.Rewrite("$<.onscroll = $1", scroll)
}

func (*SVGSVGElement) GetOnunload() (e *event.Event) {
	js.Rewrite("$<.onunload")
	return e
}

func (*SVGSVGElement) SetOnunload(e *event.Event) {
	js.Rewrite("$<.onunload = $1", e)
}

func (*SVGSVGElement) GetOnzoom() (SVGZoom *svgzoomevent.SVGZoomEvent) {
	js.Rewrite("$<.onzoom")
	return SVGZoom
}

func (*SVGSVGElement) SetOnzoom(SVGZoom *svgzoomevent.SVGZoomEvent) {
	js.Rewrite("$<.onzoom = $1", SVGZoom)
}

func (*SVGSVGElement) GetPixelUnitToMillimeterX() (pixelUnitToMillimeterX float32) {
	js.Rewrite("$<.pixelUnitToMillimeterX")
	return pixelUnitToMillimeterX
}

func (*SVGSVGElement) GetPixelUnitToMillimeterY() (pixelUnitToMillimeterY float32) {
	js.Rewrite("$<.pixelUnitToMillimeterY")
	return pixelUnitToMillimeterY
}

func (*SVGSVGElement) GetScreenPixelToMillimeterX() (screenPixelToMillimeterX float32) {
	js.Rewrite("$<.screenPixelToMillimeterX")
	return screenPixelToMillimeterX
}

func (*SVGSVGElement) GetScreenPixelToMillimeterY() (screenPixelToMillimeterY float32) {
	js.Rewrite("$<.screenPixelToMillimeterY")
	return screenPixelToMillimeterY
}

func (*SVGSVGElement) GetViewport() (viewport *svgrect.SVGRect) {
	js.Rewrite("$<.viewport")
	return viewport
}

func (*SVGSVGElement) GetWidth() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGSVGElement) GetX() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGSVGElement) GetY() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
