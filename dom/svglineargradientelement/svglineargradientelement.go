package svglineargradientelement

import (
	"github.com/matthewmueller/joy/dom/childnode"
	"github.com/matthewmueller/joy/dom/clientrect"
	"github.com/matthewmueller/joy/dom/clientrectlist"
	"github.com/matthewmueller/joy/dom/domstringmap"
	"github.com/matthewmueller/joy/dom/domtokenlist"
	"github.com/matthewmueller/joy/dom/mszoomtooptions"
	"github.com/matthewmueller/joy/dom/svganimatedenumeration"
	"github.com/matthewmueller/joy/dom/svganimatedlength"
	"github.com/matthewmueller/joy/dom/svganimatedstring"
	"github.com/matthewmueller/joy/dom/svganimatedtransformlist"
	"github.com/matthewmueller/joy/dom/svggradientelement"
	"github.com/matthewmueller/joy/dom/svgunittypes"
	"github.com/matthewmueller/joy/dom/svgurireference"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ svggradientelement.SVGGradientElement = (*SVGLinearGradientElement)(nil)
var _ svgunittypes.SVGUnitTypes = (*SVGLinearGradientElement)(nil)
var _ svgurireference.SVGURIReference = (*SVGLinearGradientElement)(nil)
var _ window.SVGElement = (*SVGLinearGradientElement)(nil)
var _ window.Element = (*SVGLinearGradientElement)(nil)
var _ window.GlobalEventHandlers = (*SVGLinearGradientElement)(nil)
var _ window.ElementTraversal = (*SVGLinearGradientElement)(nil)
var _ window.NodeSelector = (*SVGLinearGradientElement)(nil)
var _ childnode.ChildNode = (*SVGLinearGradientElement)(nil)
var _ window.Node = (*SVGLinearGradientElement)(nil)
var _ window.EventTarget = (*SVGLinearGradientElement)(nil)

// SVGLinearGradientElement struct
// js:"SVGLinearGradientElement,omit"
type SVGLinearGradientElement struct {
}

// GetAttribute fn
// js:"getAttribute"
func (*SVGLinearGradientElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

// GetAttributeNode fn
// js:"getAttributeNode"
func (*SVGLinearGradientElement) GetAttributeNode(name string) (w *window.Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

// GetAttributeNodeNS fn
// js:"getAttributeNodeNS"
func (*SVGLinearGradientElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *window.Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

// GetAttributeNS fn
// js:"getAttributeNS"
func (*SVGLinearGradientElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*SVGLinearGradientElement) GetBoundingClientRect() (c *clientrect.ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*SVGLinearGradientElement) GetClientRects() (c *clientrectlist.ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

// GetElementsByTagName fn
// js:"getElementsByTagName"
func (*SVGLinearGradientElement) GetElementsByTagName(name string) (w *window.NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*SVGLinearGradientElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *window.NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

// HasAttribute fn
// js:"hasAttribute"
func (*SVGLinearGradientElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

// HasAttributeNS fn
// js:"hasAttributeNS"
func (*SVGLinearGradientElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// MsGetRegionContent fn
// js:"msGetRegionContent"
func (*SVGLinearGradientElement) MsGetRegionContent() (w *window.MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return w
}

// MsGetUntransformedBounds fn
// js:"msGetUntransformedBounds"
func (*SVGLinearGradientElement) MsGetUntransformedBounds() (c *clientrect.ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

// MsMatchesSelector fn
// js:"msMatchesSelector"
func (*SVGLinearGradientElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

// MsReleasePointerCapture fn
// js:"msReleasePointerCapture"
func (*SVGLinearGradientElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

// MsSetPointerCapture fn
// js:"msSetPointerCapture"
func (*SVGLinearGradientElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

// MsZoomTo fn
// js:"msZoomTo"
func (*SVGLinearGradientElement) MsZoomTo(args *mszoomtooptions.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

// ReleasePointerCapture fn
// js:"releasePointerCapture"
func (*SVGLinearGradientElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

// RemoveAttribute fn
// js:"removeAttribute"
func (*SVGLinearGradientElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

// RemoveAttributeNode fn
// js:"removeAttributeNode"
func (*SVGLinearGradientElement) RemoveAttributeNode(oldAttr *window.Attr) (w *window.Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

// RemoveAttributeNS fn
// js:"removeAttributeNS"
func (*SVGLinearGradientElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// RequestFullscreen fn
// js:"requestFullscreen"
func (*SVGLinearGradientElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

// RequestPointerLock fn
// js:"requestPointerLock"
func (*SVGLinearGradientElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

// SetAttribute fn
// js:"setAttribute"
func (*SVGLinearGradientElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

// SetAttributeNode fn
// js:"setAttributeNode"
func (*SVGLinearGradientElement) SetAttributeNode(newAttr *window.Attr) (w *window.Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

// SetAttributeNodeNS fn
// js:"setAttributeNodeNS"
func (*SVGLinearGradientElement) SetAttributeNodeNS(newAttr *window.Attr) (w *window.Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

// SetAttributeNS fn
// js:"setAttributeNS"
func (*SVGLinearGradientElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// SetPointerCapture fn
// js:"setPointerCapture"
func (*SVGLinearGradientElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

// WebkitMatchesSelector fn
// js:"webkitMatchesSelector"
func (*SVGLinearGradientElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

// WebkitRequestFullscreen fn
// js:"webkitRequestFullscreen"
func (*SVGLinearGradientElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

// WebkitRequestFullScreen fn
// js:"webkitRequestFullScreen"
func (*SVGLinearGradientElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

// QuerySelector fn
// js:"querySelector"
func (*SVGLinearGradientElement) QuerySelector(selectors string) (w window.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*SVGLinearGradientElement) QuerySelectorAll(selectors string) (w *window.NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

// AppendChild fn
// js:"appendChild"
func (*SVGLinearGradientElement) AppendChild(newChild window.Node) (w window.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

// CloneNode fn
// js:"cloneNode"
func (*SVGLinearGradientElement) CloneNode(deep *bool) (w window.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*SVGLinearGradientElement) CompareDocumentPosition(other window.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*SVGLinearGradientElement) Contains(child window.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*SVGLinearGradientElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*SVGLinearGradientElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*SVGLinearGradientElement) InsertBefore(newChild window.Node, refChild *window.Node) (w window.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*SVGLinearGradientElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*SVGLinearGradientElement) IsEqualNode(arg window.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*SVGLinearGradientElement) IsSameNode(other window.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*SVGLinearGradientElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*SVGLinearGradientElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*SVGLinearGradientElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*SVGLinearGradientElement) RemoveChild(oldChild window.Node) (w window.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

// ReplaceChild fn
// js:"replaceChild"
func (*SVGLinearGradientElement) ReplaceChild(newChild window.Node, oldChild window.Node) (w window.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*SVGLinearGradientElement) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SVGLinearGradientElement) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SVGLinearGradientElement) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// X1 prop
// js:"x1"
func (*SVGLinearGradientElement) X1() (x1 *svganimatedlength.SVGAnimatedLength) {
	macro.Rewrite("$_.x1")
	return x1
}

// X2 prop
// js:"x2"
func (*SVGLinearGradientElement) X2() (x2 *svganimatedlength.SVGAnimatedLength) {
	macro.Rewrite("$_.x2")
	return x2
}

// Y1 prop
// js:"y1"
func (*SVGLinearGradientElement) Y1() (y1 *svganimatedlength.SVGAnimatedLength) {
	macro.Rewrite("$_.y1")
	return y1
}

// Y2 prop
// js:"y2"
func (*SVGLinearGradientElement) Y2() (y2 *svganimatedlength.SVGAnimatedLength) {
	macro.Rewrite("$_.y2")
	return y2
}

// GradientTransform prop
// js:"gradientTransform"
func (*SVGLinearGradientElement) GradientTransform() (gradientTransform *svganimatedtransformlist.SVGAnimatedTransformList) {
	macro.Rewrite("$_.gradientTransform")
	return gradientTransform
}

// GradientUnits prop
// js:"gradientUnits"
func (*SVGLinearGradientElement) GradientUnits() (gradientUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	macro.Rewrite("$_.gradientUnits")
	return gradientUnits
}

// SpreadMethod prop
// js:"spreadMethod"
func (*SVGLinearGradientElement) SpreadMethod() (spreadMethod *svganimatedenumeration.SVGAnimatedEnumeration) {
	macro.Rewrite("$_.spreadMethod")
	return spreadMethod
}

// Href prop
// js:"href"
func (*SVGLinearGradientElement) Href() (href *svganimatedstring.SVGAnimatedString) {
	macro.Rewrite("$_.href")
	return href
}

// Dataset prop
// js:"dataset"
func (*SVGLinearGradientElement) Dataset() (dataset *domstringmap.DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

// OwnerSVGElement prop
// js:"ownerSVGElement"
func (*SVGLinearGradientElement) OwnerSVGElement() (ownerSVGElement *window.SVGSVGElement) {
	macro.Rewrite("$_.ownerSVGElement")
	return ownerSVGElement
}

// ViewportElement prop
// js:"viewportElement"
func (*SVGLinearGradientElement) ViewportElement() (viewportElement window.SVGElement) {
	macro.Rewrite("$_.viewportElement")
	return viewportElement
}

// Xmlbase prop
// js:"xmlbase"
func (*SVGLinearGradientElement) Xmlbase() (xmlbase string) {
	macro.Rewrite("$_.xmlbase")
	return xmlbase
}

// SetXmlbase prop
// js:"xmlbase"
func (*SVGLinearGradientElement) SetXmlbase(xmlbase string) {
	macro.Rewrite("$_.xmlbase = $1", xmlbase)
}

// ClassList prop
// js:"classList"
func (*SVGLinearGradientElement) ClassList() (classList domtokenlist.DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

// ClassName prop
// js:"className"
func (*SVGLinearGradientElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*SVGLinearGradientElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

// ClientHeight prop
// js:"clientHeight"
func (*SVGLinearGradientElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

// ClientLeft prop
// js:"clientLeft"
func (*SVGLinearGradientElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

// ClientTop prop
// js:"clientTop"
func (*SVGLinearGradientElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

// ClientWidth prop
// js:"clientWidth"
func (*SVGLinearGradientElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

// ID prop
// js:"id"
func (*SVGLinearGradientElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*SVGLinearGradientElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*SVGLinearGradientElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*SVGLinearGradientElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

// MsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGLinearGradientElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

// SetMsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGLinearGradientElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

// MsRegionOverflow prop
// js:"msRegionOverflow"
func (*SVGLinearGradientElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

// Onariarequest prop
// js:"onariarequest"
func (*SVGLinearGradientElement) Onariarequest() (onariarequest func(window.Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

// SetOnariarequest prop
// js:"onariarequest"
func (*SVGLinearGradientElement) SetOnariarequest(onariarequest func(window.Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

// Oncommand prop
// js:"oncommand"
func (*SVGLinearGradientElement) Oncommand() (oncommand func(window.Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

// SetOncommand prop
// js:"oncommand"
func (*SVGLinearGradientElement) SetOncommand(oncommand func(window.Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

// Ongotpointercapture prop
// js:"ongotpointercapture"
func (*SVGLinearGradientElement) Ongotpointercapture() (ongotpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

// SetOngotpointercapture prop
// js:"ongotpointercapture"
func (*SVGLinearGradientElement) SetOngotpointercapture(ongotpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

// Onlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGLinearGradientElement) Onlostpointercapture() (onlostpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

// SetOnlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGLinearGradientElement) SetOnlostpointercapture(onlostpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGLinearGradientElement) Onmsgesturechange() (onmsgesturechange func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGLinearGradientElement) SetOnmsgesturechange(onmsgesturechange func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGLinearGradientElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGLinearGradientElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*SVGLinearGradientElement) Onmsgestureend() (onmsgestureend func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*SVGLinearGradientElement) SetOnmsgestureend(onmsgestureend func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGLinearGradientElement) Onmsgesturehold() (onmsgesturehold func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGLinearGradientElement) SetOnmsgesturehold(onmsgesturehold func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGLinearGradientElement) Onmsgesturestart() (onmsgesturestart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGLinearGradientElement) SetOnmsgesturestart(onmsgesturestart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGLinearGradientElement) Onmsgesturetap() (onmsgesturetap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGLinearGradientElement) SetOnmsgesturetap(onmsgesturetap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGLinearGradientElement) Onmsgotpointercapture() (onmsgotpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

// SetOnmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGLinearGradientElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGLinearGradientElement) Onmsinertiastart() (onmsinertiastart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGLinearGradientElement) SetOnmsinertiastart(onmsinertiastart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGLinearGradientElement) Onmslostpointercapture() (onmslostpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

// SetOnmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGLinearGradientElement) SetOnmslostpointercapture(onmslostpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*SVGLinearGradientElement) Onmspointercancel() (onmspointercancel func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*SVGLinearGradientElement) SetOnmspointercancel(onmspointercancel func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*SVGLinearGradientElement) Onmspointerdown() (onmspointerdown func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*SVGLinearGradientElement) SetOnmspointerdown(onmspointerdown func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*SVGLinearGradientElement) Onmspointerenter() (onmspointerenter func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*SVGLinearGradientElement) SetOnmspointerenter(onmspointerenter func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*SVGLinearGradientElement) Onmspointerleave() (onmspointerleave func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*SVGLinearGradientElement) SetOnmspointerleave(onmspointerleave func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*SVGLinearGradientElement) Onmspointermove() (onmspointermove func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*SVGLinearGradientElement) SetOnmspointermove(onmspointermove func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*SVGLinearGradientElement) Onmspointerout() (onmspointerout func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*SVGLinearGradientElement) SetOnmspointerout(onmspointerout func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*SVGLinearGradientElement) Onmspointerover() (onmspointerover func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*SVGLinearGradientElement) SetOnmspointerover(onmspointerover func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*SVGLinearGradientElement) Onmspointerup() (onmspointerup func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*SVGLinearGradientElement) SetOnmspointerup(onmspointerup func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*SVGLinearGradientElement) Ontouchcancel() (ontouchcancel func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*SVGLinearGradientElement) SetOntouchcancel(ontouchcancel func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*SVGLinearGradientElement) Ontouchend() (ontouchend func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*SVGLinearGradientElement) SetOntouchend(ontouchend func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*SVGLinearGradientElement) Ontouchmove() (ontouchmove func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*SVGLinearGradientElement) SetOntouchmove(ontouchmove func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*SVGLinearGradientElement) Ontouchstart() (ontouchstart func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*SVGLinearGradientElement) SetOntouchstart(ontouchstart func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGLinearGradientElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGLinearGradientElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGLinearGradientElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGLinearGradientElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// OuterHTML prop
// js:"outerHTML"
func (*SVGLinearGradientElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*SVGLinearGradientElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

// Prefix prop
// js:"prefix"
func (*SVGLinearGradientElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

// ScrollHeight prop
// js:"scrollHeight"
func (*SVGLinearGradientElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*SVGLinearGradientElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*SVGLinearGradientElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*SVGLinearGradientElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*SVGLinearGradientElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ScrollWidth prop
// js:"scrollWidth"
func (*SVGLinearGradientElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

// TagName prop
// js:"tagName"
func (*SVGLinearGradientElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

// Onpointercancel prop
// js:"onpointercancel"
func (*SVGLinearGradientElement) Onpointercancel() (onpointercancel func(window.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*SVGLinearGradientElement) SetOnpointercancel(onpointercancel func(window.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*SVGLinearGradientElement) Onpointerdown() (onpointerdown func(window.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*SVGLinearGradientElement) SetOnpointerdown(onpointerdown func(window.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*SVGLinearGradientElement) Onpointerenter() (onpointerenter func(window.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*SVGLinearGradientElement) SetOnpointerenter(onpointerenter func(window.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*SVGLinearGradientElement) Onpointerleave() (onpointerleave func(window.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*SVGLinearGradientElement) SetOnpointerleave(onpointerleave func(window.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*SVGLinearGradientElement) Onpointermove() (onpointermove func(window.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*SVGLinearGradientElement) SetOnpointermove(onpointermove func(window.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*SVGLinearGradientElement) Onpointerout() (onpointerout func(window.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*SVGLinearGradientElement) SetOnpointerout(onpointerout func(window.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*SVGLinearGradientElement) Onpointerover() (onpointerover func(window.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*SVGLinearGradientElement) SetOnpointerover(onpointerover func(window.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*SVGLinearGradientElement) Onpointerup() (onpointerup func(window.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*SVGLinearGradientElement) SetOnpointerup(onpointerup func(window.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*SVGLinearGradientElement) Onwheel() (onwheel func(window.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*SVGLinearGradientElement) SetOnwheel(onwheel func(window.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

// ChildElementCount prop
// js:"childElementCount"
func (*SVGLinearGradientElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

// FirstElementChild prop
// js:"firstElementChild"
func (*SVGLinearGradientElement) FirstElementChild() (firstElementChild window.Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
// js:"lastElementChild"
func (*SVGLinearGradientElement) LastElementChild() (lastElementChild window.Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
// js:"nextElementSibling"
func (*SVGLinearGradientElement) NextElementSibling() (nextElementSibling window.Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
// js:"previousElementSibling"
func (*SVGLinearGradientElement) PreviousElementSibling() (previousElementSibling window.Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

// Attributes prop
// js:"attributes"
func (*SVGLinearGradientElement) Attributes() (attributes *window.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*SVGLinearGradientElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*SVGLinearGradientElement) ChildNodes() (childNodes *window.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*SVGLinearGradientElement) FirstChild() (firstChild window.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*SVGLinearGradientElement) LastChild() (lastChild window.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*SVGLinearGradientElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*SVGLinearGradientElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*SVGLinearGradientElement) NextSibling() (nextSibling window.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*SVGLinearGradientElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*SVGLinearGradientElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*SVGLinearGradientElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*SVGLinearGradientElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*SVGLinearGradientElement) OwnerDocument() (ownerDocument window.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*SVGLinearGradientElement) ParentElement() (parentElement window.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*SVGLinearGradientElement) ParentNode() (parentNode window.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*SVGLinearGradientElement) PreviousSibling() (previousSibling window.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*SVGLinearGradientElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*SVGLinearGradientElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
