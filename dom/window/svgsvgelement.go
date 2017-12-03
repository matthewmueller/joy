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
	"github.com/matthewmueller/joy/macro"
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
	macro.Rewrite("$_.checkEnclosure($1, $2)", element, rect)
	return b
}

// CheckIntersection fn
// js:"checkIntersection"
func (*SVGSVGElement) CheckIntersection(element SVGElement, rect *svgrect.SVGRect) (b bool) {
	macro.Rewrite("$_.checkIntersection($1, $2)", element, rect)
	return b
}

// CreateSVGAngle fn
// js:"createSVGAngle"
func (*SVGSVGElement) CreateSVGAngle() (s *svgangle.SVGAngle) {
	macro.Rewrite("$_.createSVGAngle()")
	return s
}

// CreateSVGLength fn
// js:"createSVGLength"
func (*SVGSVGElement) CreateSVGLength() (s *svglength.SVGLength) {
	macro.Rewrite("$_.createSVGLength()")
	return s
}

// CreateSVGMatrix fn
// js:"createSVGMatrix"
func (*SVGSVGElement) CreateSVGMatrix() (s *svgmatrix.SVGMatrix) {
	macro.Rewrite("$_.createSVGMatrix()")
	return s
}

// CreateSVGNumber fn
// js:"createSVGNumber"
func (*SVGSVGElement) CreateSVGNumber() (s *svgnumber.SVGNumber) {
	macro.Rewrite("$_.createSVGNumber()")
	return s
}

// CreateSVGPoint fn
// js:"createSVGPoint"
func (*SVGSVGElement) CreateSVGPoint() (s *svgpoint.SVGPoint) {
	macro.Rewrite("$_.createSVGPoint()")
	return s
}

// CreateSVGRect fn
// js:"createSVGRect"
func (*SVGSVGElement) CreateSVGRect() (s *svgrect.SVGRect) {
	macro.Rewrite("$_.createSVGRect()")
	return s
}

// CreateSVGTransform fn
// js:"createSVGTransform"
func (*SVGSVGElement) CreateSVGTransform() (s *svgtransform.SVGTransform) {
	macro.Rewrite("$_.createSVGTransform()")
	return s
}

// CreateSVGTransformFromMatrix fn
// js:"createSVGTransformFromMatrix"
func (*SVGSVGElement) CreateSVGTransformFromMatrix(matrix *svgmatrix.SVGMatrix) (s *svgtransform.SVGTransform) {
	macro.Rewrite("$_.createSVGTransformFromMatrix($1)", matrix)
	return s
}

// DeselectAll fn
// js:"deselectAll"
func (*SVGSVGElement) DeselectAll() {
	macro.Rewrite("$_.deselectAll()")
}

// ForceRedraw fn
// js:"forceRedraw"
func (*SVGSVGElement) ForceRedraw() {
	macro.Rewrite("$_.forceRedraw()")
}

// GetComputedStyle fn
// js:"getComputedStyle"
func (*SVGSVGElement) GetComputedStyle(elt Element, pseudoElt *string) (c *CSSStyleDeclaration) {
	macro.Rewrite("$_.getComputedStyle($1, $2)", elt, pseudoElt)
	return c
}

// GetCurrentTime fn
// js:"getCurrentTime"
func (*SVGSVGElement) GetCurrentTime() (f float32) {
	macro.Rewrite("$_.getCurrentTime()")
	return f
}

// GetElementByID fn
// js:"getElementById"
func (*SVGSVGElement) GetElementByID(elementId string) (e Element) {
	macro.Rewrite("$_.getElementById($1)", elementId)
	return e
}

// GetEnclosureList fn
// js:"getEnclosureList"
func (*SVGSVGElement) GetEnclosureList(rect *svgrect.SVGRect, referenceElement SVGElement) (n *NodeList) {
	macro.Rewrite("$_.getEnclosureList($1, $2)", rect, referenceElement)
	return n
}

// GetIntersectionList fn
// js:"getIntersectionList"
func (*SVGSVGElement) GetIntersectionList(rect *svgrect.SVGRect, referenceElement SVGElement) (n *NodeList) {
	macro.Rewrite("$_.getIntersectionList($1, $2)", rect, referenceElement)
	return n
}

// PauseAnimations fn
// js:"pauseAnimations"
func (*SVGSVGElement) PauseAnimations() {
	macro.Rewrite("$_.pauseAnimations()")
}

// SetCurrentTime fn
// js:"setCurrentTime"
func (*SVGSVGElement) SetCurrentTime(seconds float32) {
	macro.Rewrite("$_.setCurrentTime($1)", seconds)
}

// SuspendRedraw fn
// js:"suspendRedraw"
func (*SVGSVGElement) SuspendRedraw(maxWaitMilliseconds uint) (u uint) {
	macro.Rewrite("$_.suspendRedraw($1)", maxWaitMilliseconds)
	return u
}

// UnpauseAnimations fn
// js:"unpauseAnimations"
func (*SVGSVGElement) UnpauseAnimations() {
	macro.Rewrite("$_.unpauseAnimations()")
}

// UnsuspendRedraw fn
// js:"unsuspendRedraw"
func (*SVGSVGElement) UnsuspendRedraw(suspendHandleID uint) {
	macro.Rewrite("$_.unsuspendRedraw($1)", suspendHandleID)
}

// UnsuspendRedrawAll fn
// js:"unsuspendRedrawAll"
func (*SVGSVGElement) UnsuspendRedrawAll() {
	macro.Rewrite("$_.unsuspendRedrawAll()")
}

// CreateEvent fn
// js:"createEvent"
func (*SVGSVGElement) CreateEvent(eventInterface string) (e Event) {
	macro.Rewrite("$_.createEvent($1)", eventInterface)
	return e
}

// GetBBox fn
// js:"getBBox"
func (*SVGSVGElement) GetBBox() (s *svgrect.SVGRect) {
	macro.Rewrite("$_.getBBox()")
	return s
}

// GetCTM fn
// js:"getCTM"
func (*SVGSVGElement) GetCTM() (s *svgmatrix.SVGMatrix) {
	macro.Rewrite("$_.getCTM()")
	return s
}

// GetScreenCTM fn
// js:"getScreenCTM"
func (*SVGSVGElement) GetScreenCTM() (s *svgmatrix.SVGMatrix) {
	macro.Rewrite("$_.getScreenCTM()")
	return s
}

// GetTransformToElement fn
// js:"getTransformToElement"
func (*SVGSVGElement) GetTransformToElement(element SVGElement) (s *svgmatrix.SVGMatrix) {
	macro.Rewrite("$_.getTransformToElement($1)", element)
	return s
}

// HasExtension fn
// js:"hasExtension"
func (*SVGSVGElement) HasExtension(extension string) (b bool) {
	macro.Rewrite("$_.hasExtension($1)", extension)
	return b
}

// GetAttribute fn
// js:"getAttribute"
func (*SVGSVGElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

// GetAttributeNode fn
// js:"getAttributeNode"
func (*SVGSVGElement) GetAttributeNode(name string) (a *Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return a
}

// GetAttributeNodeNS fn
// js:"getAttributeNodeNS"
func (*SVGSVGElement) GetAttributeNodeNS(namespaceURI string, localName string) (a *Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return a
}

// GetAttributeNS fn
// js:"getAttributeNS"
func (*SVGSVGElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*SVGSVGElement) GetBoundingClientRect() (c *clientrect.ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*SVGSVGElement) GetClientRects() (c *clientrectlist.ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

// GetElementsByTagName fn
// js:"getElementsByTagName"
func (*SVGSVGElement) GetElementsByTagName(name string) (n *NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return n
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*SVGSVGElement) GetElementsByTagNameNS(namespaceURI string, localName string) (n *NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return n
}

// HasAttribute fn
// js:"hasAttribute"
func (*SVGSVGElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

// HasAttributeNS fn
// js:"hasAttributeNS"
func (*SVGSVGElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// MsGetRegionContent fn
// js:"msGetRegionContent"
func (*SVGSVGElement) MsGetRegionContent() (m *MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return m
}

// MsGetUntransformedBounds fn
// js:"msGetUntransformedBounds"
func (*SVGSVGElement) MsGetUntransformedBounds() (c *clientrect.ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

// MsMatchesSelector fn
// js:"msMatchesSelector"
func (*SVGSVGElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

// MsReleasePointerCapture fn
// js:"msReleasePointerCapture"
func (*SVGSVGElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

// MsSetPointerCapture fn
// js:"msSetPointerCapture"
func (*SVGSVGElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

// MsZoomTo fn
// js:"msZoomTo"
func (*SVGSVGElement) MsZoomTo(args *mszoomtooptions.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

// ReleasePointerCapture fn
// js:"releasePointerCapture"
func (*SVGSVGElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

// RemoveAttribute fn
// js:"removeAttribute"
func (*SVGSVGElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

// RemoveAttributeNode fn
// js:"removeAttributeNode"
func (*SVGSVGElement) RemoveAttributeNode(oldAttr *Attr) (a *Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return a
}

// RemoveAttributeNS fn
// js:"removeAttributeNS"
func (*SVGSVGElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// RequestFullscreen fn
// js:"requestFullscreen"
func (*SVGSVGElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

// RequestPointerLock fn
// js:"requestPointerLock"
func (*SVGSVGElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

// SetAttribute fn
// js:"setAttribute"
func (*SVGSVGElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

// SetAttributeNode fn
// js:"setAttributeNode"
func (*SVGSVGElement) SetAttributeNode(newAttr *Attr) (a *Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return a
}

// SetAttributeNodeNS fn
// js:"setAttributeNodeNS"
func (*SVGSVGElement) SetAttributeNodeNS(newAttr *Attr) (a *Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return a
}

// SetAttributeNS fn
// js:"setAttributeNS"
func (*SVGSVGElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// SetPointerCapture fn
// js:"setPointerCapture"
func (*SVGSVGElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

// WebkitMatchesSelector fn
// js:"webkitMatchesSelector"
func (*SVGSVGElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

// WebkitRequestFullscreen fn
// js:"webkitRequestFullscreen"
func (*SVGSVGElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

// WebkitRequestFullScreen fn
// js:"webkitRequestFullScreen"
func (*SVGSVGElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

// QuerySelector fn
// js:"querySelector"
func (*SVGSVGElement) QuerySelector(selectors string) (e Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return e
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*SVGSVGElement) QuerySelectorAll(selectors string) (n *NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return n
}

// AppendChild fn
// js:"appendChild"
func (*SVGSVGElement) AppendChild(newChild Node) (n Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return n
}

// CloneNode fn
// js:"cloneNode"
func (*SVGSVGElement) CloneNode(deep *bool) (n Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return n
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*SVGSVGElement) CompareDocumentPosition(other Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*SVGSVGElement) Contains(child Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*SVGSVGElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*SVGSVGElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*SVGSVGElement) InsertBefore(newChild Node, refChild *Node) (n Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*SVGSVGElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*SVGSVGElement) IsEqualNode(arg Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*SVGSVGElement) IsSameNode(other Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*SVGSVGElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*SVGSVGElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*SVGSVGElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*SVGSVGElement) RemoveChild(oldChild Node) (n Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

// ReplaceChild fn
// js:"replaceChild"
func (*SVGSVGElement) ReplaceChild(newChild Node, oldChild Node) (n Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

// AddEventListener fn
// js:"addEventListener"
func (*SVGSVGElement) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SVGSVGElement) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SVGSVGElement) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// ContentScriptType prop
// js:"contentScriptType"
func (*SVGSVGElement) ContentScriptType() (contentScriptType string) {
	macro.Rewrite("$_.contentScriptType")
	return contentScriptType
}

// SetContentScriptType prop
// js:"contentScriptType"
func (*SVGSVGElement) SetContentScriptType(contentScriptType string) {
	macro.Rewrite("$_.contentScriptType = $1", contentScriptType)
}

// ContentStyleType prop
// js:"contentStyleType"
func (*SVGSVGElement) ContentStyleType() (contentStyleType string) {
	macro.Rewrite("$_.contentStyleType")
	return contentStyleType
}

// SetContentStyleType prop
// js:"contentStyleType"
func (*SVGSVGElement) SetContentStyleType(contentStyleType string) {
	macro.Rewrite("$_.contentStyleType = $1", contentStyleType)
}

// CurrentScale prop
// js:"currentScale"
func (*SVGSVGElement) CurrentScale() (currentScale float32) {
	macro.Rewrite("$_.currentScale")
	return currentScale
}

// SetCurrentScale prop
// js:"currentScale"
func (*SVGSVGElement) SetCurrentScale(currentScale float32) {
	macro.Rewrite("$_.currentScale = $1", currentScale)
}

// CurrentTranslate prop
// js:"currentTranslate"
func (*SVGSVGElement) CurrentTranslate() (currentTranslate *svgpoint.SVGPoint) {
	macro.Rewrite("$_.currentTranslate")
	return currentTranslate
}

// Height prop
// js:"height"
func (*SVGSVGElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	macro.Rewrite("$_.height")
	return height
}

// Onabort prop
// js:"onabort"
func (*SVGSVGElement) Onabort() (onabort func(Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*SVGSVGElement) SetOnabort(onabort func(Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

// Onerror prop
// js:"onerror"
func (*SVGSVGElement) Onerror() (onerror func(Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*SVGSVGElement) SetOnerror(onerror func(Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Onresize prop
// js:"onresize"
func (*SVGSVGElement) Onresize() (onresize func(Event)) {
	macro.Rewrite("$_.onresize")
	return onresize
}

// SetOnresize prop
// js:"onresize"
func (*SVGSVGElement) SetOnresize(onresize func(Event)) {
	macro.Rewrite("$_.onresize = $1", onresize)
}

// Onscroll prop
// js:"onscroll"
func (*SVGSVGElement) Onscroll() (onscroll func(UIEvent)) {
	macro.Rewrite("$_.onscroll")
	return onscroll
}

// SetOnscroll prop
// js:"onscroll"
func (*SVGSVGElement) SetOnscroll(onscroll func(UIEvent)) {
	macro.Rewrite("$_.onscroll = $1", onscroll)
}

// Onunload prop
// js:"onunload"
func (*SVGSVGElement) Onunload() (onunload func(Event)) {
	macro.Rewrite("$_.onunload")
	return onunload
}

// SetOnunload prop
// js:"onunload"
func (*SVGSVGElement) SetOnunload(onunload func(Event)) {
	macro.Rewrite("$_.onunload = $1", onunload)
}

// Onzoom prop
// js:"onzoom"
func (*SVGSVGElement) Onzoom() (onzoom func(*SVGZoomEvent)) {
	macro.Rewrite("$_.onzoom")
	return onzoom
}

// SetOnzoom prop
// js:"onzoom"
func (*SVGSVGElement) SetOnzoom(onzoom func(*SVGZoomEvent)) {
	macro.Rewrite("$_.onzoom = $1", onzoom)
}

// PixelUnitToMillimeterX prop
// js:"pixelUnitToMillimeterX"
func (*SVGSVGElement) PixelUnitToMillimeterX() (pixelUnitToMillimeterX float32) {
	macro.Rewrite("$_.pixelUnitToMillimeterX")
	return pixelUnitToMillimeterX
}

// PixelUnitToMillimeterY prop
// js:"pixelUnitToMillimeterY"
func (*SVGSVGElement) PixelUnitToMillimeterY() (pixelUnitToMillimeterY float32) {
	macro.Rewrite("$_.pixelUnitToMillimeterY")
	return pixelUnitToMillimeterY
}

// ScreenPixelToMillimeterX prop
// js:"screenPixelToMillimeterX"
func (*SVGSVGElement) ScreenPixelToMillimeterX() (screenPixelToMillimeterX float32) {
	macro.Rewrite("$_.screenPixelToMillimeterX")
	return screenPixelToMillimeterX
}

// ScreenPixelToMillimeterY prop
// js:"screenPixelToMillimeterY"
func (*SVGSVGElement) ScreenPixelToMillimeterY() (screenPixelToMillimeterY float32) {
	macro.Rewrite("$_.screenPixelToMillimeterY")
	return screenPixelToMillimeterY
}

// Viewport prop
// js:"viewport"
func (*SVGSVGElement) Viewport() (viewport *svgrect.SVGRect) {
	macro.Rewrite("$_.viewport")
	return viewport
}

// Width prop
// js:"width"
func (*SVGSVGElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	macro.Rewrite("$_.width")
	return width
}

// X prop
// js:"x"
func (*SVGSVGElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	macro.Rewrite("$_.x")
	return x
}

// Y prop
// js:"y"
func (*SVGSVGElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	macro.Rewrite("$_.y")
	return y
}

// PreserveAspectRatio prop
// js:"preserveAspectRatio"
func (*SVGSVGElement) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	macro.Rewrite("$_.preserveAspectRatio")
	return preserveAspectRatio
}

// ViewBox prop
// js:"viewBox"
func (*SVGSVGElement) ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect) {
	macro.Rewrite("$_.viewBox")
	return viewBox
}

// ZoomAndPan prop
// js:"zoomAndPan"
func (*SVGSVGElement) ZoomAndPan() (zoomAndPan uint8) {
	macro.Rewrite("$_.zoomAndPan")
	return zoomAndPan
}

// FarthestViewportElement prop
// js:"farthestViewportElement"
func (*SVGSVGElement) FarthestViewportElement() (farthestViewportElement SVGElement) {
	macro.Rewrite("$_.farthestViewportElement")
	return farthestViewportElement
}

// NearestViewportElement prop
// js:"nearestViewportElement"
func (*SVGSVGElement) NearestViewportElement() (nearestViewportElement SVGElement) {
	macro.Rewrite("$_.nearestViewportElement")
	return nearestViewportElement
}

// Transform prop
// js:"transform"
func (*SVGSVGElement) Transform() (transform *svganimatedtransformlist.SVGAnimatedTransformList) {
	macro.Rewrite("$_.transform")
	return transform
}

// RequiredExtensions prop
// js:"requiredExtensions"
func (*SVGSVGElement) RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList) {
	macro.Rewrite("$_.requiredExtensions")
	return requiredExtensions
}

// RequiredFeatures prop
// js:"requiredFeatures"
func (*SVGSVGElement) RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList) {
	macro.Rewrite("$_.requiredFeatures")
	return requiredFeatures
}

// SystemLanguage prop
// js:"systemLanguage"
func (*SVGSVGElement) SystemLanguage() (systemLanguage *svgstringlist.SVGStringList) {
	macro.Rewrite("$_.systemLanguage")
	return systemLanguage
}

// Dataset prop
// js:"dataset"
func (*SVGSVGElement) Dataset() (dataset *domstringmap.DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

// OwnerSVGElement prop
// js:"ownerSVGElement"
func (*SVGSVGElement) OwnerSVGElement() (ownerSVGElement *SVGSVGElement) {
	macro.Rewrite("$_.ownerSVGElement")
	return ownerSVGElement
}

// ViewportElement prop
// js:"viewportElement"
func (*SVGSVGElement) ViewportElement() (viewportElement SVGElement) {
	macro.Rewrite("$_.viewportElement")
	return viewportElement
}

// Xmlbase prop
// js:"xmlbase"
func (*SVGSVGElement) Xmlbase() (xmlbase string) {
	macro.Rewrite("$_.xmlbase")
	return xmlbase
}

// SetXmlbase prop
// js:"xmlbase"
func (*SVGSVGElement) SetXmlbase(xmlbase string) {
	macro.Rewrite("$_.xmlbase = $1", xmlbase)
}

// ClassList prop
// js:"classList"
func (*SVGSVGElement) ClassList() (classList domtokenlist.DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

// ClassName prop
// js:"className"
func (*SVGSVGElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*SVGSVGElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

// ClientHeight prop
// js:"clientHeight"
func (*SVGSVGElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

// ClientLeft prop
// js:"clientLeft"
func (*SVGSVGElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

// ClientTop prop
// js:"clientTop"
func (*SVGSVGElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

// ClientWidth prop
// js:"clientWidth"
func (*SVGSVGElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

// ID prop
// js:"id"
func (*SVGSVGElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*SVGSVGElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*SVGSVGElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*SVGSVGElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

// MsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGSVGElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

// SetMsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGSVGElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

// MsRegionOverflow prop
// js:"msRegionOverflow"
func (*SVGSVGElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

// Onariarequest prop
// js:"onariarequest"
func (*SVGSVGElement) Onariarequest() (onariarequest func(Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

// SetOnariarequest prop
// js:"onariarequest"
func (*SVGSVGElement) SetOnariarequest(onariarequest func(Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

// Oncommand prop
// js:"oncommand"
func (*SVGSVGElement) Oncommand() (oncommand func(Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

// SetOncommand prop
// js:"oncommand"
func (*SVGSVGElement) SetOncommand(oncommand func(Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

// Ongotpointercapture prop
// js:"ongotpointercapture"
func (*SVGSVGElement) Ongotpointercapture() (ongotpointercapture func(*PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

// SetOngotpointercapture prop
// js:"ongotpointercapture"
func (*SVGSVGElement) SetOngotpointercapture(ongotpointercapture func(*PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

// Onlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGSVGElement) Onlostpointercapture() (onlostpointercapture func(*PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

// SetOnlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGSVGElement) SetOnlostpointercapture(onlostpointercapture func(*PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGSVGElement) Onmsgesturechange() (onmsgesturechange func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGSVGElement) SetOnmsgesturechange(onmsgesturechange func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGSVGElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGSVGElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*SVGSVGElement) Onmsgestureend() (onmsgestureend func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*SVGSVGElement) SetOnmsgestureend(onmsgestureend func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGSVGElement) Onmsgesturehold() (onmsgesturehold func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGSVGElement) SetOnmsgesturehold(onmsgesturehold func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGSVGElement) Onmsgesturestart() (onmsgesturestart func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGSVGElement) SetOnmsgesturestart(onmsgesturestart func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGSVGElement) Onmsgesturetap() (onmsgesturetap func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGSVGElement) SetOnmsgesturetap(onmsgesturetap func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGSVGElement) Onmsgotpointercapture() (onmsgotpointercapture func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

// SetOnmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGSVGElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGSVGElement) Onmsinertiastart() (onmsinertiastart func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGSVGElement) SetOnmsinertiastart(onmsinertiastart func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGSVGElement) Onmslostpointercapture() (onmslostpointercapture func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

// SetOnmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGSVGElement) SetOnmslostpointercapture(onmslostpointercapture func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*SVGSVGElement) Onmspointercancel() (onmspointercancel func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*SVGSVGElement) SetOnmspointercancel(onmspointercancel func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*SVGSVGElement) Onmspointerdown() (onmspointerdown func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*SVGSVGElement) SetOnmspointerdown(onmspointerdown func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*SVGSVGElement) Onmspointerenter() (onmspointerenter func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*SVGSVGElement) SetOnmspointerenter(onmspointerenter func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*SVGSVGElement) Onmspointerleave() (onmspointerleave func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*SVGSVGElement) SetOnmspointerleave(onmspointerleave func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*SVGSVGElement) Onmspointermove() (onmspointermove func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*SVGSVGElement) SetOnmspointermove(onmspointermove func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*SVGSVGElement) Onmspointerout() (onmspointerout func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*SVGSVGElement) SetOnmspointerout(onmspointerout func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*SVGSVGElement) Onmspointerover() (onmspointerover func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*SVGSVGElement) SetOnmspointerover(onmspointerover func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*SVGSVGElement) Onmspointerup() (onmspointerup func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*SVGSVGElement) SetOnmspointerup(onmspointerup func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*SVGSVGElement) Ontouchcancel() (ontouchcancel func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*SVGSVGElement) SetOntouchcancel(ontouchcancel func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*SVGSVGElement) Ontouchend() (ontouchend func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*SVGSVGElement) SetOntouchend(ontouchend func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*SVGSVGElement) Ontouchmove() (ontouchmove func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*SVGSVGElement) SetOntouchmove(ontouchmove func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*SVGSVGElement) Ontouchstart() (ontouchstart func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*SVGSVGElement) SetOntouchstart(ontouchstart func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGSVGElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGSVGElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGSVGElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGSVGElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// OuterHTML prop
// js:"outerHTML"
func (*SVGSVGElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*SVGSVGElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

// Prefix prop
// js:"prefix"
func (*SVGSVGElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

// ScrollHeight prop
// js:"scrollHeight"
func (*SVGSVGElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*SVGSVGElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*SVGSVGElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*SVGSVGElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*SVGSVGElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ScrollWidth prop
// js:"scrollWidth"
func (*SVGSVGElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

// TagName prop
// js:"tagName"
func (*SVGSVGElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

// Onpointercancel prop
// js:"onpointercancel"
func (*SVGSVGElement) Onpointercancel() (onpointercancel func(Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*SVGSVGElement) SetOnpointercancel(onpointercancel func(Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*SVGSVGElement) Onpointerdown() (onpointerdown func(Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*SVGSVGElement) SetOnpointerdown(onpointerdown func(Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*SVGSVGElement) Onpointerenter() (onpointerenter func(Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*SVGSVGElement) SetOnpointerenter(onpointerenter func(Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*SVGSVGElement) Onpointerleave() (onpointerleave func(Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*SVGSVGElement) SetOnpointerleave(onpointerleave func(Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*SVGSVGElement) Onpointermove() (onpointermove func(Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*SVGSVGElement) SetOnpointermove(onpointermove func(Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*SVGSVGElement) Onpointerout() (onpointerout func(Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*SVGSVGElement) SetOnpointerout(onpointerout func(Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*SVGSVGElement) Onpointerover() (onpointerover func(Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*SVGSVGElement) SetOnpointerover(onpointerover func(Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*SVGSVGElement) Onpointerup() (onpointerup func(Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*SVGSVGElement) SetOnpointerup(onpointerup func(Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*SVGSVGElement) Onwheel() (onwheel func(Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*SVGSVGElement) SetOnwheel(onwheel func(Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

// ChildElementCount prop
// js:"childElementCount"
func (*SVGSVGElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

// FirstElementChild prop
// js:"firstElementChild"
func (*SVGSVGElement) FirstElementChild() (firstElementChild Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
// js:"lastElementChild"
func (*SVGSVGElement) LastElementChild() (lastElementChild Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
// js:"nextElementSibling"
func (*SVGSVGElement) NextElementSibling() (nextElementSibling Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
// js:"previousElementSibling"
func (*SVGSVGElement) PreviousElementSibling() (previousElementSibling Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

// Attributes prop
// js:"attributes"
func (*SVGSVGElement) Attributes() (attributes *NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*SVGSVGElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*SVGSVGElement) ChildNodes() (childNodes *NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*SVGSVGElement) FirstChild() (firstChild Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*SVGSVGElement) LastChild() (lastChild Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*SVGSVGElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*SVGSVGElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*SVGSVGElement) NextSibling() (nextSibling Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*SVGSVGElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*SVGSVGElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*SVGSVGElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*SVGSVGElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*SVGSVGElement) OwnerDocument() (ownerDocument Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*SVGSVGElement) ParentElement() (parentElement HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*SVGSVGElement) ParentNode() (parentNode Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*SVGSVGElement) PreviousSibling() (previousSibling Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*SVGSVGElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*SVGSVGElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
