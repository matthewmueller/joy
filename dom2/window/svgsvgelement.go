package window

import (
	"github.com/matthewmueller/golly/dom2/svgangle"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svgfittoviewbox"
	"github.com/matthewmueller/golly/dom2/svglength"
	"github.com/matthewmueller/golly/dom2/svgmatrix"
	"github.com/matthewmueller/golly/dom2/svgnumber"
	"github.com/matthewmueller/golly/dom2/svgpoint"
	"github.com/matthewmueller/golly/dom2/svgrect"
	"github.com/matthewmueller/golly/dom2/svgtransform"
	"github.com/matthewmueller/golly/dom2/svgzoomandpan"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGSVGElement,omit"
type SVGSVGElement struct {
	SVGGraphicsElement
	DocumentEvent
	svgfittoviewbox.SVGFitToViewBox
	svgzoomandpan.SVGZoomAndPan
}

// CheckEnclosure
func (*SVGSVGElement) CheckEnclosure(element SVGElement, rect *svgrect.SVGRect) (b bool) {
	js.Rewrite("$<.CheckEnclosure($1, $2)", element, rect)
	return b
}

// CheckIntersection
func (*SVGSVGElement) CheckIntersection(element SVGElement, rect *svgrect.SVGRect) (b bool) {
	js.Rewrite("$<.CheckIntersection($1, $2)", element, rect)
	return b
}

// CreateSVGAngle
func (*SVGSVGElement) CreateSVGAngle() (s *svgangle.SVGAngle) {
	js.Rewrite("$<.CreateSVGAngle()")
	return s
}

// CreateSVGLength
func (*SVGSVGElement) CreateSVGLength() (s *svglength.SVGLength) {
	js.Rewrite("$<.CreateSVGLength()")
	return s
}

// CreateSVGMatrix
func (*SVGSVGElement) CreateSVGMatrix() (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.CreateSVGMatrix()")
	return s
}

// CreateSVGNumber
func (*SVGSVGElement) CreateSVGNumber() (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.CreateSVGNumber()")
	return s
}

// CreateSVGPoint
func (*SVGSVGElement) CreateSVGPoint() (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.CreateSVGPoint()")
	return s
}

// CreateSVGRect
func (*SVGSVGElement) CreateSVGRect() (s *svgrect.SVGRect) {
	js.Rewrite("$<.CreateSVGRect()")
	return s
}

// CreateSVGTransform
func (*SVGSVGElement) CreateSVGTransform() (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.CreateSVGTransform()")
	return s
}

// CreateSVGTransformFromMatrix
func (*SVGSVGElement) CreateSVGTransformFromMatrix(matrix *svgmatrix.SVGMatrix) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.CreateSVGTransformFromMatrix($1)", matrix)
	return s
}

// DeselectAll
func (*SVGSVGElement) DeselectAll() {
	js.Rewrite("$<.DeselectAll()")
}

// ForceRedraw
func (*SVGSVGElement) ForceRedraw() {
	js.Rewrite("$<.ForceRedraw()")
}

// GetComputedStyle
func (*SVGSVGElement) GetComputedStyle(elt Element, pseudoElt *string) (c *CSSStyleDeclaration) {
	js.Rewrite("$<.GetComputedStyle($1, $2)", elt, pseudoElt)
	return c
}

// GetCurrentTime
func (*SVGSVGElement) GetCurrentTime() (f float32) {
	js.Rewrite("$<.GetCurrentTime()")
	return f
}

// GetElementByID
func (*SVGSVGElement) GetElementByID(elementId string) (e Element) {
	js.Rewrite("$<.GetElementByID($1)", elementId)
	return e
}

// GetEnclosureList
func (*SVGSVGElement) GetEnclosureList(rect *svgrect.SVGRect, referenceElement SVGElement) (n *NodeList) {
	js.Rewrite("$<.GetEnclosureList($1, $2)", rect, referenceElement)
	return n
}

// GetIntersectionList
func (*SVGSVGElement) GetIntersectionList(rect *svgrect.SVGRect, referenceElement SVGElement) (n *NodeList) {
	js.Rewrite("$<.GetIntersectionList($1, $2)", rect, referenceElement)
	return n
}

// PauseAnimations
func (*SVGSVGElement) PauseAnimations() {
	js.Rewrite("$<.PauseAnimations()")
}

// SetCurrentTime
func (*SVGSVGElement) SetCurrentTime(seconds float32) {
	js.Rewrite("$<.SetCurrentTime($1)", seconds)
}

// SuspendRedraw
func (*SVGSVGElement) SuspendRedraw(maxWaitMilliseconds uint) (u uint) {
	js.Rewrite("$<.SuspendRedraw($1)", maxWaitMilliseconds)
	return u
}

// UnpauseAnimations
func (*SVGSVGElement) UnpauseAnimations() {
	js.Rewrite("$<.UnpauseAnimations()")
}

// UnsuspendRedraw
func (*SVGSVGElement) UnsuspendRedraw(suspendHandleID uint) {
	js.Rewrite("$<.UnsuspendRedraw($1)", suspendHandleID)
}

// UnsuspendRedrawAll
func (*SVGSVGElement) UnsuspendRedrawAll() {
	js.Rewrite("$<.UnsuspendRedrawAll()")
}

// ContentScriptType
func (*SVGSVGElement) ContentScriptType() (contentScriptType string) {
	js.Rewrite("$<.ContentScriptType")
	return contentScriptType
}

// ContentScriptType
func (*SVGSVGElement) SetContentScriptType(contentScriptType string) {
	js.Rewrite("$<.ContentScriptType = $1", contentScriptType)
}

// ContentStyleType
func (*SVGSVGElement) ContentStyleType() (contentStyleType string) {
	js.Rewrite("$<.ContentStyleType")
	return contentStyleType
}

// ContentStyleType
func (*SVGSVGElement) SetContentStyleType(contentStyleType string) {
	js.Rewrite("$<.ContentStyleType = $1", contentStyleType)
}

// CurrentScale
func (*SVGSVGElement) CurrentScale() (currentScale float32) {
	js.Rewrite("$<.CurrentScale")
	return currentScale
}

// CurrentScale
func (*SVGSVGElement) SetCurrentScale(currentScale float32) {
	js.Rewrite("$<.CurrentScale = $1", currentScale)
}

// CurrentTranslate
func (*SVGSVGElement) CurrentTranslate() (currentTranslate *svgpoint.SVGPoint) {
	js.Rewrite("$<.CurrentTranslate")
	return currentTranslate
}

// Height
func (*SVGSVGElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Onabort
func (*SVGSVGElement) Onabort() (onabort func(Event)) {
	js.Rewrite("$<.Onabort")
	return onabort
}

// Onabort
func (*SVGSVGElement) SetOnabort(onabort func(Event)) {
	js.Rewrite("$<.Onabort = $1", onabort)
}

// Onerror
func (*SVGSVGElement) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*SVGSVGElement) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onresize
func (*SVGSVGElement) Onresize() (onresize func(Event)) {
	js.Rewrite("$<.Onresize")
	return onresize
}

// Onresize
func (*SVGSVGElement) SetOnresize(onresize func(Event)) {
	js.Rewrite("$<.Onresize = $1", onresize)
}

// Onscroll
func (*SVGSVGElement) Onscroll() (onscroll func(UIEvent)) {
	js.Rewrite("$<.Onscroll")
	return onscroll
}

// Onscroll
func (*SVGSVGElement) SetOnscroll(onscroll func(UIEvent)) {
	js.Rewrite("$<.Onscroll = $1", onscroll)
}

// Onunload
func (*SVGSVGElement) Onunload() (onunload func(Event)) {
	js.Rewrite("$<.Onunload")
	return onunload
}

// Onunload
func (*SVGSVGElement) SetOnunload(onunload func(Event)) {
	js.Rewrite("$<.Onunload = $1", onunload)
}

// Onzoom
func (*SVGSVGElement) Onzoom() (onzoom func(*SVGZoomEvent)) {
	js.Rewrite("$<.Onzoom")
	return onzoom
}

// Onzoom
func (*SVGSVGElement) SetOnzoom(onzoom func(*SVGZoomEvent)) {
	js.Rewrite("$<.Onzoom = $1", onzoom)
}

// PixelUnitToMillimeterX
func (*SVGSVGElement) PixelUnitToMillimeterX() (pixelUnitToMillimeterX float32) {
	js.Rewrite("$<.PixelUnitToMillimeterX")
	return pixelUnitToMillimeterX
}

// PixelUnitToMillimeterY
func (*SVGSVGElement) PixelUnitToMillimeterY() (pixelUnitToMillimeterY float32) {
	js.Rewrite("$<.PixelUnitToMillimeterY")
	return pixelUnitToMillimeterY
}

// ScreenPixelToMillimeterX
func (*SVGSVGElement) ScreenPixelToMillimeterX() (screenPixelToMillimeterX float32) {
	js.Rewrite("$<.ScreenPixelToMillimeterX")
	return screenPixelToMillimeterX
}

// ScreenPixelToMillimeterY
func (*SVGSVGElement) ScreenPixelToMillimeterY() (screenPixelToMillimeterY float32) {
	js.Rewrite("$<.ScreenPixelToMillimeterY")
	return screenPixelToMillimeterY
}

// Viewport
func (*SVGSVGElement) Viewport() (viewport *svgrect.SVGRect) {
	js.Rewrite("$<.Viewport")
	return viewport
}

// Width
func (*SVGSVGElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGSVGElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGSVGElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}
