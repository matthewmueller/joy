package window

import (
	"github.com/matthewmueller/joy/dom/childnode"
	"github.com/matthewmueller/joy/dom/clientrect"
	"github.com/matthewmueller/joy/dom/clientrectlist"
	"github.com/matthewmueller/joy/dom/domstringmap"
	"github.com/matthewmueller/joy/dom/domtokenlist"
	"github.com/matthewmueller/joy/dom/mszoomtooptions"
	"github.com/matthewmueller/joy/dom/svgangle"
	"github.com/matthewmueller/joy/dom/svganimatedlength"
	"github.com/matthewmueller/joy/dom/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/joy/dom/svganimatedrect"
	"github.com/matthewmueller/joy/dom/svganimatedtransformlist"
	"github.com/matthewmueller/joy/dom/svglength"
	"github.com/matthewmueller/joy/dom/svgmatrix"
	"github.com/matthewmueller/joy/dom/svgnumber"
	"github.com/matthewmueller/joy/dom/svgpoint"
	"github.com/matthewmueller/joy/dom/svgrect"
	"github.com/matthewmueller/joy/dom/svgstringlist"
	"github.com/matthewmueller/joy/dom/svgtests"
	"github.com/matthewmueller/joy/dom/svgtransform"
	"github.com/matthewmueller/joy/js"
)

var _ SVGGraphicsElement = (*SVGSVGElement)(nil)
var _ svgtests.SVGTests = (*SVGSVGElement)(nil)
var _ SVGElement = (*SVGSVGElement)(nil)
var _ Element = (*SVGSVGElement)(nil)
var _ GlobalEventHandlers = (*SVGSVGElement)(nil)
var _ ElementTraversal = (*SVGSVGElement)(nil)
var _ NodeSelector = (*SVGSVGElement)(nil)
var _ childnode.ChildNode = (*SVGSVGElement)(nil)
var _ Node = (*SVGSVGElement)(nil)
var _ EventTarget = (*SVGSVGElement)(nil)

// SVGSVGElement struct
// js:"SVGSVGElement,omit"
type SVGSVGElement struct {
}

// CheckEnclosure fn
// js:"checkEnclosure"
func (*SVGSVGElement) CheckEnclosure(element SVGElement, rect *svgrect.SVGRect) (b bool) {
	js.Rewrite("$_.checkEnclosure($1, $2)", element, rect)
	return b
}

// CheckIntersection fn
// js:"checkIntersection"
func (*SVGSVGElement) CheckIntersection(element SVGElement, rect *svgrect.SVGRect) (b bool) {
	js.Rewrite("$_.checkIntersection($1, $2)", element, rect)
	return b
}

// CreateSVGAngle fn
// js:"createSVGAngle"
func (*SVGSVGElement) CreateSVGAngle() (s *svgangle.SVGAngle) {
	js.Rewrite("$_.createSVGAngle()")
	return s
}

// CreateSVGLength fn
// js:"createSVGLength"
func (*SVGSVGElement) CreateSVGLength() (s *svglength.SVGLength) {
	js.Rewrite("$_.createSVGLength()")
	return s
}

// CreateSVGMatrix fn
// js:"createSVGMatrix"
func (*SVGSVGElement) CreateSVGMatrix() (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$_.createSVGMatrix()")
	return s
}

// CreateSVGNumber fn
// js:"createSVGNumber"
func (*SVGSVGElement) CreateSVGNumber() (s *svgnumber.SVGNumber) {
	js.Rewrite("$_.createSVGNumber()")
	return s
}

// CreateSVGPoint fn
// js:"createSVGPoint"
func (*SVGSVGElement) CreateSVGPoint() (s *svgpoint.SVGPoint) {
	js.Rewrite("$_.createSVGPoint()")
	return s
}

// CreateSVGRect fn
// js:"createSVGRect"
func (*SVGSVGElement) CreateSVGRect() (s *svgrect.SVGRect) {
	js.Rewrite("$_.createSVGRect()")
	return s
}

// CreateSVGTransform fn
// js:"createSVGTransform"
func (*SVGSVGElement) CreateSVGTransform() (s *svgtransform.SVGTransform) {
	js.Rewrite("$_.createSVGTransform()")
	return s
}

// CreateSVGTransformFromMatrix fn
// js:"createSVGTransformFromMatrix"
func (*SVGSVGElement) CreateSVGTransformFromMatrix(matrix *svgmatrix.SVGMatrix) (s *svgtransform.SVGTransform) {
	js.Rewrite("$_.createSVGTransformFromMatrix($1)", matrix)
	return s
}

// DeselectAll fn
// js:"deselectAll"
func (*SVGSVGElement) DeselectAll() {
	js.Rewrite("$_.deselectAll()")
}

// ForceRedraw fn
// js:"forceRedraw"
func (*SVGSVGElement) ForceRedraw() {
	js.Rewrite("$_.forceRedraw()")
}

// GetComputedStyle fn
// js:"getComputedStyle"
func (*SVGSVGElement) GetComputedStyle(elt Element, pseudoElt *string) (c *CSSStyleDeclaration) {
	js.Rewrite("$_.getComputedStyle($1, $2)", elt, pseudoElt)
	return c
}

// GetCurrentTime fn
// js:"getCurrentTime"
func (*SVGSVGElement) GetCurrentTime() (f float32) {
	js.Rewrite("$_.getCurrentTime()")
	return f
}

// GetElementByID fn
// js:"getElementById"
func (*SVGSVGElement) GetElementByID(elementId string) (e Element) {
	js.Rewrite("$_.getElementById($1)", elementId)
	return e
}

// GetEnclosureList fn
// js:"getEnclosureList"
func (*SVGSVGElement) GetEnclosureList(rect *svgrect.SVGRect, referenceElement SVGElement) (n *NodeList) {
	js.Rewrite("$_.getEnclosureList($1, $2)", rect, referenceElement)
	return n
}

// GetIntersectionList fn
// js:"getIntersectionList"
func (*SVGSVGElement) GetIntersectionList(rect *svgrect.SVGRect, referenceElement SVGElement) (n *NodeList) {
	js.Rewrite("$_.getIntersectionList($1, $2)", rect, referenceElement)
	return n
}

// PauseAnimations fn
// js:"pauseAnimations"
func (*SVGSVGElement) PauseAnimations() {
	js.Rewrite("$_.pauseAnimations()")
}

// SetCurrentTime fn
// js:"setCurrentTime"
func (*SVGSVGElement) SetCurrentTime(seconds float32) {
	js.Rewrite("$_.setCurrentTime($1)", seconds)
}

// SuspendRedraw fn
// js:"suspendRedraw"
func (*SVGSVGElement) SuspendRedraw(maxWaitMilliseconds uint) (u uint) {
	js.Rewrite("$_.suspendRedraw($1)", maxWaitMilliseconds)
	return u
}

// UnpauseAnimations fn
// js:"unpauseAnimations"
func (*SVGSVGElement) UnpauseAnimations() {
	js.Rewrite("$_.unpauseAnimations()")
}

// UnsuspendRedraw fn
// js:"unsuspendRedraw"
func (*SVGSVGElement) UnsuspendRedraw(suspendHandleID uint) {
	js.Rewrite("$_.unsuspendRedraw($1)", suspendHandleID)
}

// UnsuspendRedrawAll fn
// js:"unsuspendRedrawAll"
func (*SVGSVGElement) UnsuspendRedrawAll() {
	js.Rewrite("$_.unsuspendRedrawAll()")
}

// CreateEvent fn
// js:"createEvent"
func (*SVGSVGElement) CreateEvent(eventInterface string) (e Event) {
	js.Rewrite("$_.createEvent($1)", eventInterface)
	return e
}

// GetBBox fn
// js:"getBBox"
func (*SVGSVGElement) GetBBox() (s *svgrect.SVGRect) {
	js.Rewrite("$_.getBBox()")
	return s
}

// GetCTM fn
// js:"getCTM"
func (*SVGSVGElement) GetCTM() (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$_.getCTM()")
	return s
}

// GetScreenCTM fn
// js:"getScreenCTM"
func (*SVGSVGElement) GetScreenCTM() (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$_.getScreenCTM()")
	return s
}

// GetTransformToElement fn
// js:"getTransformToElement"
func (*SVGSVGElement) GetTransformToElement(element SVGElement) (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$_.getTransformToElement($1)", element)
	return s
}

// HasExtension fn
// js:"hasExtension"
func (*SVGSVGElement) HasExtension(extension string) (b bool) {
	js.Rewrite("$_.hasExtension($1)", extension)
	return b
}

// GetAttribute fn
// js:"getAttribute"
func (*SVGSVGElement) GetAttribute(qualifiedName string) (s string) {
	js.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

// GetAttributeNode fn
// js:"getAttributeNode"
func (*SVGSVGElement) GetAttributeNode(name string) (a *Attr) {
	js.Rewrite("$_.getAttributeNode($1)", name)
	return a
}

// GetAttributeNodeNS fn
// js:"getAttributeNodeNS"
func (*SVGSVGElement) GetAttributeNodeNS(namespaceURI string, localName string) (a *Attr) {
	js.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return a
}

// GetAttributeNS fn
// js:"getAttributeNS"
func (*SVGSVGElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	js.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*SVGSVGElement) GetBoundingClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*SVGSVGElement) GetClientRects() (c *clientrectlist.ClientRectList) {
	js.Rewrite("$_.getClientRects()")
	return c
}

// GetElementsByTagName fn
// js:"getElementsByTagName"
func (*SVGSVGElement) GetElementsByTagName(name string) (n *NodeList) {
	js.Rewrite("$_.getElementsByTagName($1)", name)
	return n
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*SVGSVGElement) GetElementsByTagNameNS(namespaceURI string, localName string) (n *NodeList) {
	js.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return n
}

// HasAttribute fn
// js:"hasAttribute"
func (*SVGSVGElement) HasAttribute(name string) (b bool) {
	js.Rewrite("$_.hasAttribute($1)", name)
	return b
}

// HasAttributeNS fn
// js:"hasAttributeNS"
func (*SVGSVGElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	js.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// MsGetRegionContent fn
// js:"msGetRegionContent"
func (*SVGSVGElement) MsGetRegionContent() (m *MSRangeCollection) {
	js.Rewrite("$_.msGetRegionContent()")
	return m
}

// MsGetUntransformedBounds fn
// js:"msGetUntransformedBounds"
func (*SVGSVGElement) MsGetUntransformedBounds() (c *clientrect.ClientRect) {
	js.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

// MsMatchesSelector fn
// js:"msMatchesSelector"
func (*SVGSVGElement) MsMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

// MsReleasePointerCapture fn
// js:"msReleasePointerCapture"
func (*SVGSVGElement) MsReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

// MsSetPointerCapture fn
// js:"msSetPointerCapture"
func (*SVGSVGElement) MsSetPointerCapture(pointerId int) {
	js.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

// MsZoomTo fn
// js:"msZoomTo"
func (*SVGSVGElement) MsZoomTo(args *mszoomtooptions.MsZoomToOptions) {
	js.Rewrite("$_.msZoomTo($1)", args)
}

// ReleasePointerCapture fn
// js:"releasePointerCapture"
func (*SVGSVGElement) ReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

// RemoveAttribute fn
// js:"removeAttribute"
func (*SVGSVGElement) RemoveAttribute(qualifiedName string) {
	js.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

// RemoveAttributeNode fn
// js:"removeAttributeNode"
func (*SVGSVGElement) RemoveAttributeNode(oldAttr *Attr) (a *Attr) {
	js.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return a
}

// RemoveAttributeNS fn
// js:"removeAttributeNS"
func (*SVGSVGElement) RemoveAttributeNS(namespaceURI string, localName string) {
	js.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// RequestFullscreen fn
// js:"requestFullscreen"
func (*SVGSVGElement) RequestFullscreen() {
	js.Rewrite("$_.requestFullscreen()")
}

// RequestPointerLock fn
// js:"requestPointerLock"
func (*SVGSVGElement) RequestPointerLock() {
	js.Rewrite("$_.requestPointerLock()")
}

// SetAttribute fn
// js:"setAttribute"
func (*SVGSVGElement) SetAttribute(qualifiedName string, value string) {
	js.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

// SetAttributeNode fn
// js:"setAttributeNode"
func (*SVGSVGElement) SetAttributeNode(newAttr *Attr) (a *Attr) {
	js.Rewrite("$_.setAttributeNode($1)", newAttr)
	return a
}

// SetAttributeNodeNS fn
// js:"setAttributeNodeNS"
func (*SVGSVGElement) SetAttributeNodeNS(newAttr *Attr) (a *Attr) {
	js.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return a
}

// SetAttributeNS fn
// js:"setAttributeNS"
func (*SVGSVGElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	js.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// SetPointerCapture fn
// js:"setPointerCapture"
func (*SVGSVGElement) SetPointerCapture(pointerId int) {
	js.Rewrite("$_.setPointerCapture($1)", pointerId)
}

// WebkitMatchesSelector fn
// js:"webkitMatchesSelector"
func (*SVGSVGElement) WebkitMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

// WebkitRequestFullscreen fn
// js:"webkitRequestFullscreen"
func (*SVGSVGElement) WebkitRequestFullscreen() {
	js.Rewrite("$_.webkitRequestFullscreen()")
}

// WebkitRequestFullScreen fn
// js:"webkitRequestFullScreen"
func (*SVGSVGElement) WebkitRequestFullScreen() {
	js.Rewrite("$_.webkitRequestFullScreen()")
}

// QuerySelector fn
// js:"querySelector"
func (*SVGSVGElement) QuerySelector(selectors string) (e Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return e
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*SVGSVGElement) QuerySelectorAll(selectors string) (n *NodeList) {
	js.Rewrite("$_.querySelectorAll($1)", selectors)
	return n
}

// AppendChild fn
// js:"appendChild"
func (*SVGSVGElement) AppendChild(newChild Node) (n Node) {
	js.Rewrite("$_.appendChild($1)", newChild)
	return n
}

// CloneNode fn
// js:"cloneNode"
func (*SVGSVGElement) CloneNode(deep *bool) (n Node) {
	js.Rewrite("$_.cloneNode($1)", deep)
	return n
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*SVGSVGElement) CompareDocumentPosition(other Node) (u uint8) {
	js.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*SVGSVGElement) Contains(child Node) (b bool) {
	js.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*SVGSVGElement) HasAttributes() (b bool) {
	js.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*SVGSVGElement) HasChildNodes() (b bool) {
	js.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*SVGSVGElement) InsertBefore(newChild Node, refChild *Node) (n Node) {
	js.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*SVGSVGElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	js.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*SVGSVGElement) IsEqualNode(arg Node) (b bool) {
	js.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*SVGSVGElement) IsSameNode(other Node) (b bool) {
	js.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*SVGSVGElement) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*SVGSVGElement) LookupPrefix(namespaceURI string) (s string) {
	js.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*SVGSVGElement) Normalize() {
	js.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*SVGSVGElement) RemoveChild(oldChild Node) (n Node) {
	js.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

// ReplaceChild fn
// js:"replaceChild"
func (*SVGSVGElement) ReplaceChild(newChild Node, oldChild Node) (n Node) {
	js.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

// AddEventListener fn
// js:"addEventListener"
func (*SVGSVGElement) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SVGSVGElement) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SVGSVGElement) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// ContentScriptType prop
// js:"contentScriptType"
func (*SVGSVGElement) ContentScriptType() (contentScriptType string) {
	js.Rewrite("$_.contentScriptType")
	return contentScriptType
}

// SetContentScriptType prop
// js:"contentScriptType"
func (*SVGSVGElement) SetContentScriptType(contentScriptType string) {
	js.Rewrite("$_.contentScriptType = $1", contentScriptType)
}

// ContentStyleType prop
// js:"contentStyleType"
func (*SVGSVGElement) ContentStyleType() (contentStyleType string) {
	js.Rewrite("$_.contentStyleType")
	return contentStyleType
}

// SetContentStyleType prop
// js:"contentStyleType"
func (*SVGSVGElement) SetContentStyleType(contentStyleType string) {
	js.Rewrite("$_.contentStyleType = $1", contentStyleType)
}

// CurrentScale prop
// js:"currentScale"
func (*SVGSVGElement) CurrentScale() (currentScale float32) {
	js.Rewrite("$_.currentScale")
	return currentScale
}

// SetCurrentScale prop
// js:"currentScale"
func (*SVGSVGElement) SetCurrentScale(currentScale float32) {
	js.Rewrite("$_.currentScale = $1", currentScale)
}

// CurrentTranslate prop
// js:"currentTranslate"
func (*SVGSVGElement) CurrentTranslate() (currentTranslate *svgpoint.SVGPoint) {
	js.Rewrite("$_.currentTranslate")
	return currentTranslate
}

// Height prop
// js:"height"
func (*SVGSVGElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$_.height")
	return height
}

// Onabort prop
// js:"onabort"
func (*SVGSVGElement) Onabort() (onabort func(Event)) {
	js.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*SVGSVGElement) SetOnabort(onabort func(Event)) {
	js.Rewrite("$_.onabort = $1", onabort)
}

// Onerror prop
// js:"onerror"
func (*SVGSVGElement) Onerror() (onerror func(Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*SVGSVGElement) SetOnerror(onerror func(Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onresize prop
// js:"onresize"
func (*SVGSVGElement) Onresize() (onresize func(Event)) {
	js.Rewrite("$_.onresize")
	return onresize
}

// SetOnresize prop
// js:"onresize"
func (*SVGSVGElement) SetOnresize(onresize func(Event)) {
	js.Rewrite("$_.onresize = $1", onresize)
}

// Onscroll prop
// js:"onscroll"
func (*SVGSVGElement) Onscroll() (onscroll func(UIEvent)) {
	js.Rewrite("$_.onscroll")
	return onscroll
}

// SetOnscroll prop
// js:"onscroll"
func (*SVGSVGElement) SetOnscroll(onscroll func(UIEvent)) {
	js.Rewrite("$_.onscroll = $1", onscroll)
}

// Onunload prop
// js:"onunload"
func (*SVGSVGElement) Onunload() (onunload func(Event)) {
	js.Rewrite("$_.onunload")
	return onunload
}

// SetOnunload prop
// js:"onunload"
func (*SVGSVGElement) SetOnunload(onunload func(Event)) {
	js.Rewrite("$_.onunload = $1", onunload)
}

// Onzoom prop
// js:"onzoom"
func (*SVGSVGElement) Onzoom() (onzoom func(*SVGZoomEvent)) {
	js.Rewrite("$_.onzoom")
	return onzoom
}

// SetOnzoom prop
// js:"onzoom"
func (*SVGSVGElement) SetOnzoom(onzoom func(*SVGZoomEvent)) {
	js.Rewrite("$_.onzoom = $1", onzoom)
}

// PixelUnitToMillimeterX prop
// js:"pixelUnitToMillimeterX"
func (*SVGSVGElement) PixelUnitToMillimeterX() (pixelUnitToMillimeterX float32) {
	js.Rewrite("$_.pixelUnitToMillimeterX")
	return pixelUnitToMillimeterX
}

// PixelUnitToMillimeterY prop
// js:"pixelUnitToMillimeterY"
func (*SVGSVGElement) PixelUnitToMillimeterY() (pixelUnitToMillimeterY float32) {
	js.Rewrite("$_.pixelUnitToMillimeterY")
	return pixelUnitToMillimeterY
}

// ScreenPixelToMillimeterX prop
// js:"screenPixelToMillimeterX"
func (*SVGSVGElement) ScreenPixelToMillimeterX() (screenPixelToMillimeterX float32) {
	js.Rewrite("$_.screenPixelToMillimeterX")
	return screenPixelToMillimeterX
}

// ScreenPixelToMillimeterY prop
// js:"screenPixelToMillimeterY"
func (*SVGSVGElement) ScreenPixelToMillimeterY() (screenPixelToMillimeterY float32) {
	js.Rewrite("$_.screenPixelToMillimeterY")
	return screenPixelToMillimeterY
}

// Viewport prop
// js:"viewport"
func (*SVGSVGElement) Viewport() (viewport *svgrect.SVGRect) {
	js.Rewrite("$_.viewport")
	return viewport
}

// Width prop
// js:"width"
func (*SVGSVGElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$_.width")
	return width
}

// X prop
// js:"x"
func (*SVGSVGElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$_.x")
	return x
}

// Y prop
// js:"y"
func (*SVGSVGElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$_.y")
	return y
}

// PreserveAspectRatio prop
// js:"preserveAspectRatio"
func (*SVGSVGElement) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$_.preserveAspectRatio")
	return preserveAspectRatio
}

// ViewBox prop
// js:"viewBox"
func (*SVGSVGElement) ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect) {
	js.Rewrite("$_.viewBox")
	return viewBox
}

// ZoomAndPan prop
// js:"zoomAndPan"
func (*SVGSVGElement) ZoomAndPan() (zoomAndPan uint8) {
	js.Rewrite("$_.zoomAndPan")
	return zoomAndPan
}

// FarthestViewportElement prop
// js:"farthestViewportElement"
func (*SVGSVGElement) FarthestViewportElement() (farthestViewportElement SVGElement) {
	js.Rewrite("$_.farthestViewportElement")
	return farthestViewportElement
}

// NearestViewportElement prop
// js:"nearestViewportElement"
func (*SVGSVGElement) NearestViewportElement() (nearestViewportElement SVGElement) {
	js.Rewrite("$_.nearestViewportElement")
	return nearestViewportElement
}

// Transform prop
// js:"transform"
func (*SVGSVGElement) Transform() (transform *svganimatedtransformlist.SVGAnimatedTransformList) {
	js.Rewrite("$_.transform")
	return transform
}

// RequiredExtensions prop
// js:"requiredExtensions"
func (*SVGSVGElement) RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList) {
	js.Rewrite("$_.requiredExtensions")
	return requiredExtensions
}

// RequiredFeatures prop
// js:"requiredFeatures"
func (*SVGSVGElement) RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList) {
	js.Rewrite("$_.requiredFeatures")
	return requiredFeatures
}

// SystemLanguage prop
// js:"systemLanguage"
func (*SVGSVGElement) SystemLanguage() (systemLanguage *svgstringlist.SVGStringList) {
	js.Rewrite("$_.systemLanguage")
	return systemLanguage
}

// Dataset prop
// js:"dataset"
func (*SVGSVGElement) Dataset() (dataset *domstringmap.DOMStringMap) {
	js.Rewrite("$_.dataset")
	return dataset
}

// OwnerSVGElement prop
// js:"ownerSVGElement"
func (*SVGSVGElement) OwnerSVGElement() (ownerSVGElement *SVGSVGElement) {
	js.Rewrite("$_.ownerSVGElement")
	return ownerSVGElement
}

// ViewportElement prop
// js:"viewportElement"
func (*SVGSVGElement) ViewportElement() (viewportElement SVGElement) {
	js.Rewrite("$_.viewportElement")
	return viewportElement
}

// Xmlbase prop
// js:"xmlbase"
func (*SVGSVGElement) Xmlbase() (xmlbase string) {
	js.Rewrite("$_.xmlbase")
	return xmlbase
}

// SetXmlbase prop
// js:"xmlbase"
func (*SVGSVGElement) SetXmlbase(xmlbase string) {
	js.Rewrite("$_.xmlbase = $1", xmlbase)
}

// ClassList prop
// js:"classList"
func (*SVGSVGElement) ClassList() (classList domtokenlist.DOMTokenList) {
	js.Rewrite("$_.classList")
	return classList
}

// ClassName prop
// js:"className"
func (*SVGSVGElement) ClassName() (className string) {
	js.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*SVGSVGElement) SetClassName(className string) {
	js.Rewrite("$_.className = $1", className)
}

// ClientHeight prop
// js:"clientHeight"
func (*SVGSVGElement) ClientHeight() (clientHeight int) {
	js.Rewrite("$_.clientHeight")
	return clientHeight
}

// ClientLeft prop
// js:"clientLeft"
func (*SVGSVGElement) ClientLeft() (clientLeft int) {
	js.Rewrite("$_.clientLeft")
	return clientLeft
}

// ClientTop prop
// js:"clientTop"
func (*SVGSVGElement) ClientTop() (clientTop int) {
	js.Rewrite("$_.clientTop")
	return clientTop
}

// ClientWidth prop
// js:"clientWidth"
func (*SVGSVGElement) ClientWidth() (clientWidth int) {
	js.Rewrite("$_.clientWidth")
	return clientWidth
}

// ID prop
// js:"id"
func (*SVGSVGElement) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*SVGSVGElement) SetID(id string) {
	js.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*SVGSVGElement) InnerHTML() (innerHTML string) {
	js.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*SVGSVGElement) SetInnerHTML(innerHTML string) {
	js.Rewrite("$_.innerHTML = $1", innerHTML)
}

// MsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGSVGElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

// SetMsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGSVGElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

// MsRegionOverflow prop
// js:"msRegionOverflow"
func (*SVGSVGElement) MsRegionOverflow() (msRegionOverflow string) {
	js.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

// Onariarequest prop
// js:"onariarequest"
func (*SVGSVGElement) Onariarequest() (onariarequest func(Event)) {
	js.Rewrite("$_.onariarequest")
	return onariarequest
}

// SetOnariarequest prop
// js:"onariarequest"
func (*SVGSVGElement) SetOnariarequest(onariarequest func(Event)) {
	js.Rewrite("$_.onariarequest = $1", onariarequest)
}

// Oncommand prop
// js:"oncommand"
func (*SVGSVGElement) Oncommand() (oncommand func(Event)) {
	js.Rewrite("$_.oncommand")
	return oncommand
}

// SetOncommand prop
// js:"oncommand"
func (*SVGSVGElement) SetOncommand(oncommand func(Event)) {
	js.Rewrite("$_.oncommand = $1", oncommand)
}

// Ongotpointercapture prop
// js:"ongotpointercapture"
func (*SVGSVGElement) Ongotpointercapture() (ongotpointercapture func(*PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

// SetOngotpointercapture prop
// js:"ongotpointercapture"
func (*SVGSVGElement) SetOngotpointercapture(ongotpointercapture func(*PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

// Onlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGSVGElement) Onlostpointercapture() (onlostpointercapture func(*PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

// SetOnlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGSVGElement) SetOnlostpointercapture(onlostpointercapture func(*PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGSVGElement) Onmsgesturechange() (onmsgesturechange func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGSVGElement) SetOnmsgesturechange(onmsgesturechange func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGSVGElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGSVGElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*SVGSVGElement) Onmsgestureend() (onmsgestureend func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*SVGSVGElement) SetOnmsgestureend(onmsgestureend func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGSVGElement) Onmsgesturehold() (onmsgesturehold func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGSVGElement) SetOnmsgesturehold(onmsgesturehold func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGSVGElement) Onmsgesturestart() (onmsgesturestart func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGSVGElement) SetOnmsgesturestart(onmsgesturestart func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGSVGElement) Onmsgesturetap() (onmsgesturetap func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGSVGElement) SetOnmsgesturetap(onmsgesturetap func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGSVGElement) Onmsgotpointercapture() (onmsgotpointercapture func(*MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

// SetOnmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGSVGElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGSVGElement) Onmsinertiastart() (onmsinertiastart func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGSVGElement) SetOnmsinertiastart(onmsinertiastart func(*MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGSVGElement) Onmslostpointercapture() (onmslostpointercapture func(*MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

// SetOnmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGSVGElement) SetOnmslostpointercapture(onmslostpointercapture func(*MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*SVGSVGElement) Onmspointercancel() (onmspointercancel func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*SVGSVGElement) SetOnmspointercancel(onmspointercancel func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*SVGSVGElement) Onmspointerdown() (onmspointerdown func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*SVGSVGElement) SetOnmspointerdown(onmspointerdown func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*SVGSVGElement) Onmspointerenter() (onmspointerenter func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*SVGSVGElement) SetOnmspointerenter(onmspointerenter func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*SVGSVGElement) Onmspointerleave() (onmspointerleave func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*SVGSVGElement) SetOnmspointerleave(onmspointerleave func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*SVGSVGElement) Onmspointermove() (onmspointermove func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*SVGSVGElement) SetOnmspointermove(onmspointermove func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*SVGSVGElement) Onmspointerout() (onmspointerout func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*SVGSVGElement) SetOnmspointerout(onmspointerout func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*SVGSVGElement) Onmspointerover() (onmspointerover func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*SVGSVGElement) SetOnmspointerover(onmspointerover func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*SVGSVGElement) Onmspointerup() (onmspointerup func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*SVGSVGElement) SetOnmspointerup(onmspointerup func(*MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*SVGSVGElement) Ontouchcancel() (ontouchcancel func(*TouchEvent)) {
	js.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*SVGSVGElement) SetOntouchcancel(ontouchcancel func(*TouchEvent)) {
	js.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*SVGSVGElement) Ontouchend() (ontouchend func(*TouchEvent)) {
	js.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*SVGSVGElement) SetOntouchend(ontouchend func(*TouchEvent)) {
	js.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*SVGSVGElement) Ontouchmove() (ontouchmove func(*TouchEvent)) {
	js.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*SVGSVGElement) SetOntouchmove(ontouchmove func(*TouchEvent)) {
	js.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*SVGSVGElement) Ontouchstart() (ontouchstart func(*TouchEvent)) {
	js.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*SVGSVGElement) SetOntouchstart(ontouchstart func(*TouchEvent)) {
	js.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGSVGElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGSVGElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGSVGElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGSVGElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// OuterHTML prop
// js:"outerHTML"
func (*SVGSVGElement) OuterHTML() (outerHTML string) {
	js.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*SVGSVGElement) SetOuterHTML(outerHTML string) {
	js.Rewrite("$_.outerHTML = $1", outerHTML)
}

// Prefix prop
// js:"prefix"
func (*SVGSVGElement) Prefix() (prefix string) {
	js.Rewrite("$_.prefix")
	return prefix
}

// ScrollHeight prop
// js:"scrollHeight"
func (*SVGSVGElement) ScrollHeight() (scrollHeight int) {
	js.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*SVGSVGElement) ScrollLeft() (scrollLeft int) {
	js.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*SVGSVGElement) SetScrollLeft(scrollLeft int) {
	js.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*SVGSVGElement) ScrollTop() (scrollTop int) {
	js.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*SVGSVGElement) SetScrollTop(scrollTop int) {
	js.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ScrollWidth prop
// js:"scrollWidth"
func (*SVGSVGElement) ScrollWidth() (scrollWidth int) {
	js.Rewrite("$_.scrollWidth")
	return scrollWidth
}

// TagName prop
// js:"tagName"
func (*SVGSVGElement) TagName() (tagName string) {
	js.Rewrite("$_.tagName")
	return tagName
}

// Onpointercancel prop
// js:"onpointercancel"
func (*SVGSVGElement) Onpointercancel() (onpointercancel func(Event)) {
	js.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*SVGSVGElement) SetOnpointercancel(onpointercancel func(Event)) {
	js.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*SVGSVGElement) Onpointerdown() (onpointerdown func(Event)) {
	js.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*SVGSVGElement) SetOnpointerdown(onpointerdown func(Event)) {
	js.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*SVGSVGElement) Onpointerenter() (onpointerenter func(Event)) {
	js.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*SVGSVGElement) SetOnpointerenter(onpointerenter func(Event)) {
	js.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*SVGSVGElement) Onpointerleave() (onpointerleave func(Event)) {
	js.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*SVGSVGElement) SetOnpointerleave(onpointerleave func(Event)) {
	js.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*SVGSVGElement) Onpointermove() (onpointermove func(Event)) {
	js.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*SVGSVGElement) SetOnpointermove(onpointermove func(Event)) {
	js.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*SVGSVGElement) Onpointerout() (onpointerout func(Event)) {
	js.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*SVGSVGElement) SetOnpointerout(onpointerout func(Event)) {
	js.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*SVGSVGElement) Onpointerover() (onpointerover func(Event)) {
	js.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*SVGSVGElement) SetOnpointerover(onpointerover func(Event)) {
	js.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*SVGSVGElement) Onpointerup() (onpointerup func(Event)) {
	js.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*SVGSVGElement) SetOnpointerup(onpointerup func(Event)) {
	js.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*SVGSVGElement) Onwheel() (onwheel func(Event)) {
	js.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*SVGSVGElement) SetOnwheel(onwheel func(Event)) {
	js.Rewrite("$_.onwheel = $1", onwheel)
}

// ChildElementCount prop
// js:"childElementCount"
func (*SVGSVGElement) ChildElementCount() (childElementCount uint) {
	js.Rewrite("$_.childElementCount")
	return childElementCount
}

// FirstElementChild prop
// js:"firstElementChild"
func (*SVGSVGElement) FirstElementChild() (firstElementChild Element) {
	js.Rewrite("$_.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
// js:"lastElementChild"
func (*SVGSVGElement) LastElementChild() (lastElementChild Element) {
	js.Rewrite("$_.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
// js:"nextElementSibling"
func (*SVGSVGElement) NextElementSibling() (nextElementSibling Element) {
	js.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
// js:"previousElementSibling"
func (*SVGSVGElement) PreviousElementSibling() (previousElementSibling Element) {
	js.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

// Attributes prop
// js:"attributes"
func (*SVGSVGElement) Attributes() (attributes *NamedNodeMap) {
	js.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*SVGSVGElement) BaseURI() (baseURI string) {
	js.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*SVGSVGElement) ChildNodes() (childNodes *NodeList) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*SVGSVGElement) FirstChild() (firstChild Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*SVGSVGElement) LastChild() (lastChild Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*SVGSVGElement) LocalName() (localName string) {
	js.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*SVGSVGElement) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*SVGSVGElement) NextSibling() (nextSibling Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*SVGSVGElement) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*SVGSVGElement) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*SVGSVGElement) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*SVGSVGElement) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*SVGSVGElement) OwnerDocument() (ownerDocument Document) {
	js.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*SVGSVGElement) ParentElement() (parentElement HTMLElement) {
	js.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*SVGSVGElement) ParentNode() (parentNode Node) {
	js.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*SVGSVGElement) PreviousSibling() (previousSibling Node) {
	js.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*SVGSVGElement) TextContent() (textContent string) {
	js.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*SVGSVGElement) SetTextContent(textContent string) {
	js.Rewrite("$_.textContent = $1", textContent)
}
