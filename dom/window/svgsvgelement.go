package window

import (
	"github.com/matthewmueller/golly/dom/svgangle"
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/dom/svganimatedrect"
	"github.com/matthewmueller/golly/dom/svglength"
	"github.com/matthewmueller/golly/dom/svgmatrix"
	"github.com/matthewmueller/golly/dom/svgnumber"
	"github.com/matthewmueller/golly/dom/svgpoint"
	"github.com/matthewmueller/golly/dom/svgrect"
	"github.com/matthewmueller/golly/dom/svgtransform"
	"github.com/matthewmueller/golly/js"
)

// SVGSVGElement struct
// js:"SVGSVGElement,omit"
type SVGSVGElement struct {
	SVGGraphicsElement
}

// CreateEvent fn
func (*SVGSVGElement) CreateEvent(eventInterface string) (e Event) {
	js.Rewrite("$<.createEvent($1)", eventInterface)
	return e
}

// CheckEnclosure fn
func (*SVGSVGElement) CheckEnclosure(element SVGElement, rect *svgrect.SVGRect) (b bool) {
	js.Rewrite("$<.checkEnclosure($1, $2)", element, rect)
	return b
}

// CheckIntersection fn
func (*SVGSVGElement) CheckIntersection(element SVGElement, rect *svgrect.SVGRect) (b bool) {
	js.Rewrite("$<.checkIntersection($1, $2)", element, rect)
	return b
}

// CreateSVGAngle fn
func (*SVGSVGElement) CreateSVGAngle() (s *svgangle.SVGAngle) {
	js.Rewrite("$<.createSVGAngle()")
	return s
}

// CreateSVGLength fn
func (*SVGSVGElement) CreateSVGLength() (s *svglength.SVGLength) {
	js.Rewrite("$<.createSVGLength()")
	return s
}

// CreateSVGMatrix fn
func (*SVGSVGElement) CreateSVGMatrix() (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.createSVGMatrix()")
	return s
}

// CreateSVGNumber fn
func (*SVGSVGElement) CreateSVGNumber() (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.createSVGNumber()")
	return s
}

// CreateSVGPoint fn
func (*SVGSVGElement) CreateSVGPoint() (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.createSVGPoint()")
	return s
}

// CreateSVGRect fn
func (*SVGSVGElement) CreateSVGRect() (s *svgrect.SVGRect) {
	js.Rewrite("$<.createSVGRect()")
	return s
}

// CreateSVGTransform fn
func (*SVGSVGElement) CreateSVGTransform() (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.createSVGTransform()")
	return s
}

// CreateSVGTransformFromMatrix fn
func (*SVGSVGElement) CreateSVGTransformFromMatrix(matrix *svgmatrix.SVGMatrix) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.createSVGTransformFromMatrix($1)", matrix)
	return s
}

// DeselectAll fn
func (*SVGSVGElement) DeselectAll() {
	js.Rewrite("$<.deselectAll()")
}

// ForceRedraw fn
func (*SVGSVGElement) ForceRedraw() {
	js.Rewrite("$<.forceRedraw()")
}

// GetComputedStyle fn
func (*SVGSVGElement) GetComputedStyle(elt Element, pseudoElt *string) (c *CSSStyleDeclaration) {
	js.Rewrite("$<.getComputedStyle($1, $2)", elt, pseudoElt)
	return c
}

// GetCurrentTime fn
func (*SVGSVGElement) GetCurrentTime() (f float32) {
	js.Rewrite("$<.getCurrentTime()")
	return f
}

// GetElementByID fn
func (*SVGSVGElement) GetElementByID(elementId string) (e Element) {
	js.Rewrite("$<.getElementById($1)", elementId)
	return e
}

// GetEnclosureList fn
func (*SVGSVGElement) GetEnclosureList(rect *svgrect.SVGRect, referenceElement SVGElement) (n *NodeList) {
	js.Rewrite("$<.getEnclosureList($1, $2)", rect, referenceElement)
	return n
}

// GetIntersectionList fn
func (*SVGSVGElement) GetIntersectionList(rect *svgrect.SVGRect, referenceElement SVGElement) (n *NodeList) {
	js.Rewrite("$<.getIntersectionList($1, $2)", rect, referenceElement)
	return n
}

// PauseAnimations fn
func (*SVGSVGElement) PauseAnimations() {
	js.Rewrite("$<.pauseAnimations()")
}

// SetCurrentTime fn
func (*SVGSVGElement) SetCurrentTime(seconds float32) {
	js.Rewrite("$<.setCurrentTime($1)", seconds)
}

// SuspendRedraw fn
func (*SVGSVGElement) SuspendRedraw(maxWaitMilliseconds uint) (u uint) {
	js.Rewrite("$<.suspendRedraw($1)", maxWaitMilliseconds)
	return u
}

// UnpauseAnimations fn
func (*SVGSVGElement) UnpauseAnimations() {
	js.Rewrite("$<.unpauseAnimations()")
}

// UnsuspendRedraw fn
func (*SVGSVGElement) UnsuspendRedraw(suspendHandleID uint) {
	js.Rewrite("$<.unsuspendRedraw($1)", suspendHandleID)
}

// UnsuspendRedrawAll fn
func (*SVGSVGElement) UnsuspendRedrawAll() {
	js.Rewrite("$<.unsuspendRedrawAll()")
}

// PreserveAspectRatio prop
func (*SVGSVGElement) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.preserveAspectRatio")
	return preserveAspectRatio
}

// ViewBox prop
func (*SVGSVGElement) ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect) {
	js.Rewrite("$<.viewBox")
	return viewBox
}

// ZoomAndPan prop
func (*SVGSVGElement) ZoomAndPan() (zoomAndPan uint8) {
	js.Rewrite("$<.zoomAndPan")
	return zoomAndPan
}

// ContentScriptType prop
func (*SVGSVGElement) ContentScriptType() (contentScriptType string) {
	js.Rewrite("$<.contentScriptType")
	return contentScriptType
}

// ContentScriptType prop
func (*SVGSVGElement) SetContentScriptType(contentScriptType string) {
	js.Rewrite("$<.contentScriptType = $1", contentScriptType)
}

// ContentStyleType prop
func (*SVGSVGElement) ContentStyleType() (contentStyleType string) {
	js.Rewrite("$<.contentStyleType")
	return contentStyleType
}

// ContentStyleType prop
func (*SVGSVGElement) SetContentStyleType(contentStyleType string) {
	js.Rewrite("$<.contentStyleType = $1", contentStyleType)
}

// CurrentScale prop
func (*SVGSVGElement) CurrentScale() (currentScale float32) {
	js.Rewrite("$<.currentScale")
	return currentScale
}

// CurrentScale prop
func (*SVGSVGElement) SetCurrentScale(currentScale float32) {
	js.Rewrite("$<.currentScale = $1", currentScale)
}

// CurrentTranslate prop
func (*SVGSVGElement) CurrentTranslate() (currentTranslate *svgpoint.SVGPoint) {
	js.Rewrite("$<.currentTranslate")
	return currentTranslate
}

// Height prop
func (*SVGSVGElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Onabort prop
func (*SVGSVGElement) Onabort() (onabort func(Event)) {
	js.Rewrite("$<.onabort")
	return onabort
}

// Onabort prop
func (*SVGSVGElement) SetOnabort(onabort func(Event)) {
	js.Rewrite("$<.onabort = $1", onabort)
}

// Onerror prop
func (*SVGSVGElement) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*SVGSVGElement) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onresize prop
func (*SVGSVGElement) Onresize() (onresize func(Event)) {
	js.Rewrite("$<.onresize")
	return onresize
}

// Onresize prop
func (*SVGSVGElement) SetOnresize(onresize func(Event)) {
	js.Rewrite("$<.onresize = $1", onresize)
}

// Onscroll prop
func (*SVGSVGElement) Onscroll() (onscroll func(UIEvent)) {
	js.Rewrite("$<.onscroll")
	return onscroll
}

// Onscroll prop
func (*SVGSVGElement) SetOnscroll(onscroll func(UIEvent)) {
	js.Rewrite("$<.onscroll = $1", onscroll)
}

// Onunload prop
func (*SVGSVGElement) Onunload() (onunload func(Event)) {
	js.Rewrite("$<.onunload")
	return onunload
}

// Onunload prop
func (*SVGSVGElement) SetOnunload(onunload func(Event)) {
	js.Rewrite("$<.onunload = $1", onunload)
}

// Onzoom prop
func (*SVGSVGElement) Onzoom() (onzoom func(*SVGZoomEvent)) {
	js.Rewrite("$<.onzoom")
	return onzoom
}

// Onzoom prop
func (*SVGSVGElement) SetOnzoom(onzoom func(*SVGZoomEvent)) {
	js.Rewrite("$<.onzoom = $1", onzoom)
}

// PixelUnitToMillimeterX prop
func (*SVGSVGElement) PixelUnitToMillimeterX() (pixelUnitToMillimeterX float32) {
	js.Rewrite("$<.pixelUnitToMillimeterX")
	return pixelUnitToMillimeterX
}

// PixelUnitToMillimeterY prop
func (*SVGSVGElement) PixelUnitToMillimeterY() (pixelUnitToMillimeterY float32) {
	js.Rewrite("$<.pixelUnitToMillimeterY")
	return pixelUnitToMillimeterY
}

// ScreenPixelToMillimeterX prop
func (*SVGSVGElement) ScreenPixelToMillimeterX() (screenPixelToMillimeterX float32) {
	js.Rewrite("$<.screenPixelToMillimeterX")
	return screenPixelToMillimeterX
}

// ScreenPixelToMillimeterY prop
func (*SVGSVGElement) ScreenPixelToMillimeterY() (screenPixelToMillimeterY float32) {
	js.Rewrite("$<.screenPixelToMillimeterY")
	return screenPixelToMillimeterY
}

// Viewport prop
func (*SVGSVGElement) Viewport() (viewport *svgrect.SVGRect) {
	js.Rewrite("$<.viewport")
	return viewport
}

// Width prop
func (*SVGSVGElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGSVGElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGSVGElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
