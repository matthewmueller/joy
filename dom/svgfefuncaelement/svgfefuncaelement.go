package svgfefuncaelement

import (
	"github.com/matthewmueller/joy/dom/childnode"
	"github.com/matthewmueller/joy/dom/clientrect"
	"github.com/matthewmueller/joy/dom/clientrectlist"
	"github.com/matthewmueller/joy/dom/domstringmap"
	"github.com/matthewmueller/joy/dom/domtokenlist"
	"github.com/matthewmueller/joy/dom/mszoomtooptions"
	"github.com/matthewmueller/joy/dom/svganimatedenumeration"
	"github.com/matthewmueller/joy/dom/svganimatednumber"
	"github.com/matthewmueller/joy/dom/svganimatednumberlist"
	"github.com/matthewmueller/joy/dom/svgcomponenttransferfunctionelement"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ svgcomponenttransferfunctionelement.SVGComponentTransferFunctionElement = (*SVGFEFuncAElement)(nil)
var _ window.SVGElement = (*SVGFEFuncAElement)(nil)
var _ window.Element = (*SVGFEFuncAElement)(nil)
var _ window.GlobalEventHandlers = (*SVGFEFuncAElement)(nil)
var _ window.ElementTraversal = (*SVGFEFuncAElement)(nil)
var _ window.NodeSelector = (*SVGFEFuncAElement)(nil)
var _ childnode.ChildNode = (*SVGFEFuncAElement)(nil)
var _ window.Node = (*SVGFEFuncAElement)(nil)
var _ window.EventTarget = (*SVGFEFuncAElement)(nil)

// SVGFEFuncAElement struct
// js:"SVGFEFuncAElement,omit"
type SVGFEFuncAElement struct {
}

// GetAttribute fn
// js:"getAttribute"
func (*SVGFEFuncAElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

// GetAttributeNode fn
// js:"getAttributeNode"
func (*SVGFEFuncAElement) GetAttributeNode(name string) (w *window.Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

// GetAttributeNodeNS fn
// js:"getAttributeNodeNS"
func (*SVGFEFuncAElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *window.Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

// GetAttributeNS fn
// js:"getAttributeNS"
func (*SVGFEFuncAElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*SVGFEFuncAElement) GetBoundingClientRect() (c *clientrect.ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*SVGFEFuncAElement) GetClientRects() (c *clientrectlist.ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

// GetElementsByTagName fn
// js:"getElementsByTagName"
func (*SVGFEFuncAElement) GetElementsByTagName(name string) (w *window.NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*SVGFEFuncAElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *window.NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

// HasAttribute fn
// js:"hasAttribute"
func (*SVGFEFuncAElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

// HasAttributeNS fn
// js:"hasAttributeNS"
func (*SVGFEFuncAElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// MsGetRegionContent fn
// js:"msGetRegionContent"
func (*SVGFEFuncAElement) MsGetRegionContent() (w *window.MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return w
}

// MsGetUntransformedBounds fn
// js:"msGetUntransformedBounds"
func (*SVGFEFuncAElement) MsGetUntransformedBounds() (c *clientrect.ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

// MsMatchesSelector fn
// js:"msMatchesSelector"
func (*SVGFEFuncAElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

// MsReleasePointerCapture fn
// js:"msReleasePointerCapture"
func (*SVGFEFuncAElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

// MsSetPointerCapture fn
// js:"msSetPointerCapture"
func (*SVGFEFuncAElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

// MsZoomTo fn
// js:"msZoomTo"
func (*SVGFEFuncAElement) MsZoomTo(args *mszoomtooptions.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

// ReleasePointerCapture fn
// js:"releasePointerCapture"
func (*SVGFEFuncAElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

// RemoveAttribute fn
// js:"removeAttribute"
func (*SVGFEFuncAElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

// RemoveAttributeNode fn
// js:"removeAttributeNode"
func (*SVGFEFuncAElement) RemoveAttributeNode(oldAttr *window.Attr) (w *window.Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

// RemoveAttributeNS fn
// js:"removeAttributeNS"
func (*SVGFEFuncAElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// RequestFullscreen fn
// js:"requestFullscreen"
func (*SVGFEFuncAElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

// RequestPointerLock fn
// js:"requestPointerLock"
func (*SVGFEFuncAElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

// SetAttribute fn
// js:"setAttribute"
func (*SVGFEFuncAElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

// SetAttributeNode fn
// js:"setAttributeNode"
func (*SVGFEFuncAElement) SetAttributeNode(newAttr *window.Attr) (w *window.Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

// SetAttributeNodeNS fn
// js:"setAttributeNodeNS"
func (*SVGFEFuncAElement) SetAttributeNodeNS(newAttr *window.Attr) (w *window.Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

// SetAttributeNS fn
// js:"setAttributeNS"
func (*SVGFEFuncAElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// SetPointerCapture fn
// js:"setPointerCapture"
func (*SVGFEFuncAElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

// WebkitMatchesSelector fn
// js:"webkitMatchesSelector"
func (*SVGFEFuncAElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

// WebkitRequestFullscreen fn
// js:"webkitRequestFullscreen"
func (*SVGFEFuncAElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

// WebkitRequestFullScreen fn
// js:"webkitRequestFullScreen"
func (*SVGFEFuncAElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

// QuerySelector fn
// js:"querySelector"
func (*SVGFEFuncAElement) QuerySelector(selectors string) (w window.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*SVGFEFuncAElement) QuerySelectorAll(selectors string) (w *window.NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

// AppendChild fn
// js:"appendChild"
func (*SVGFEFuncAElement) AppendChild(newChild window.Node) (w window.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

// CloneNode fn
// js:"cloneNode"
func (*SVGFEFuncAElement) CloneNode(deep *bool) (w window.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*SVGFEFuncAElement) CompareDocumentPosition(other window.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*SVGFEFuncAElement) Contains(child window.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*SVGFEFuncAElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*SVGFEFuncAElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*SVGFEFuncAElement) InsertBefore(newChild window.Node, refChild *window.Node) (w window.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*SVGFEFuncAElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*SVGFEFuncAElement) IsEqualNode(arg window.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*SVGFEFuncAElement) IsSameNode(other window.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*SVGFEFuncAElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*SVGFEFuncAElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*SVGFEFuncAElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*SVGFEFuncAElement) RemoveChild(oldChild window.Node) (w window.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

// ReplaceChild fn
// js:"replaceChild"
func (*SVGFEFuncAElement) ReplaceChild(newChild window.Node, oldChild window.Node) (w window.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*SVGFEFuncAElement) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SVGFEFuncAElement) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SVGFEFuncAElement) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Amplitude prop
// js:"amplitude"
func (*SVGFEFuncAElement) Amplitude() (amplitude *svganimatednumber.SVGAnimatedNumber) {
	macro.Rewrite("$_.amplitude")
	return amplitude
}

// Exponent prop
// js:"exponent"
func (*SVGFEFuncAElement) Exponent() (exponent *svganimatednumber.SVGAnimatedNumber) {
	macro.Rewrite("$_.exponent")
	return exponent
}

// Intercept prop
// js:"intercept"
func (*SVGFEFuncAElement) Intercept() (intercept *svganimatednumber.SVGAnimatedNumber) {
	macro.Rewrite("$_.intercept")
	return intercept
}

// Offset prop
// js:"offset"
func (*SVGFEFuncAElement) Offset() (offset *svganimatednumber.SVGAnimatedNumber) {
	macro.Rewrite("$_.offset")
	return offset
}

// Slope prop
// js:"slope"
func (*SVGFEFuncAElement) Slope() (slope *svganimatednumber.SVGAnimatedNumber) {
	macro.Rewrite("$_.slope")
	return slope
}

// TableValues prop
// js:"tableValues"
func (*SVGFEFuncAElement) TableValues() (tableValues *svganimatednumberlist.SVGAnimatedNumberList) {
	macro.Rewrite("$_.tableValues")
	return tableValues
}

// Type prop
// js:"type"
func (*SVGFEFuncAElement) Type() (kind *svganimatedenumeration.SVGAnimatedEnumeration) {
	macro.Rewrite("$_.type")
	return kind
}

// Dataset prop
// js:"dataset"
func (*SVGFEFuncAElement) Dataset() (dataset *domstringmap.DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

// OwnerSVGElement prop
// js:"ownerSVGElement"
func (*SVGFEFuncAElement) OwnerSVGElement() (ownerSVGElement *window.SVGSVGElement) {
	macro.Rewrite("$_.ownerSVGElement")
	return ownerSVGElement
}

// ViewportElement prop
// js:"viewportElement"
func (*SVGFEFuncAElement) ViewportElement() (viewportElement window.SVGElement) {
	macro.Rewrite("$_.viewportElement")
	return viewportElement
}

// Xmlbase prop
// js:"xmlbase"
func (*SVGFEFuncAElement) Xmlbase() (xmlbase string) {
	macro.Rewrite("$_.xmlbase")
	return xmlbase
}

// SetXmlbase prop
// js:"xmlbase"
func (*SVGFEFuncAElement) SetXmlbase(xmlbase string) {
	macro.Rewrite("$_.xmlbase = $1", xmlbase)
}

// ClassList prop
// js:"classList"
func (*SVGFEFuncAElement) ClassList() (classList domtokenlist.DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

// ClassName prop
// js:"className"
func (*SVGFEFuncAElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*SVGFEFuncAElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

// ClientHeight prop
// js:"clientHeight"
func (*SVGFEFuncAElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

// ClientLeft prop
// js:"clientLeft"
func (*SVGFEFuncAElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

// ClientTop prop
// js:"clientTop"
func (*SVGFEFuncAElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

// ClientWidth prop
// js:"clientWidth"
func (*SVGFEFuncAElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

// ID prop
// js:"id"
func (*SVGFEFuncAElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*SVGFEFuncAElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*SVGFEFuncAElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*SVGFEFuncAElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

// MsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGFEFuncAElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

// SetMsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGFEFuncAElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

// MsRegionOverflow prop
// js:"msRegionOverflow"
func (*SVGFEFuncAElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

// Onariarequest prop
// js:"onariarequest"
func (*SVGFEFuncAElement) Onariarequest() (onariarequest func(window.Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

// SetOnariarequest prop
// js:"onariarequest"
func (*SVGFEFuncAElement) SetOnariarequest(onariarequest func(window.Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

// Oncommand prop
// js:"oncommand"
func (*SVGFEFuncAElement) Oncommand() (oncommand func(window.Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

// SetOncommand prop
// js:"oncommand"
func (*SVGFEFuncAElement) SetOncommand(oncommand func(window.Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

// Ongotpointercapture prop
// js:"ongotpointercapture"
func (*SVGFEFuncAElement) Ongotpointercapture() (ongotpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

// SetOngotpointercapture prop
// js:"ongotpointercapture"
func (*SVGFEFuncAElement) SetOngotpointercapture(ongotpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

// Onlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGFEFuncAElement) Onlostpointercapture() (onlostpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

// SetOnlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGFEFuncAElement) SetOnlostpointercapture(onlostpointercapture func(*window.PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGFEFuncAElement) Onmsgesturechange() (onmsgesturechange func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGFEFuncAElement) SetOnmsgesturechange(onmsgesturechange func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGFEFuncAElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGFEFuncAElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*SVGFEFuncAElement) Onmsgestureend() (onmsgestureend func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*SVGFEFuncAElement) SetOnmsgestureend(onmsgestureend func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGFEFuncAElement) Onmsgesturehold() (onmsgesturehold func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGFEFuncAElement) SetOnmsgesturehold(onmsgesturehold func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGFEFuncAElement) Onmsgesturestart() (onmsgesturestart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGFEFuncAElement) SetOnmsgesturestart(onmsgesturestart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGFEFuncAElement) Onmsgesturetap() (onmsgesturetap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGFEFuncAElement) SetOnmsgesturetap(onmsgesturetap func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGFEFuncAElement) Onmsgotpointercapture() (onmsgotpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

// SetOnmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGFEFuncAElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGFEFuncAElement) Onmsinertiastart() (onmsinertiastart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGFEFuncAElement) SetOnmsinertiastart(onmsinertiastart func(*window.MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGFEFuncAElement) Onmslostpointercapture() (onmslostpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

// SetOnmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGFEFuncAElement) SetOnmslostpointercapture(onmslostpointercapture func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*SVGFEFuncAElement) Onmspointercancel() (onmspointercancel func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*SVGFEFuncAElement) SetOnmspointercancel(onmspointercancel func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*SVGFEFuncAElement) Onmspointerdown() (onmspointerdown func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*SVGFEFuncAElement) SetOnmspointerdown(onmspointerdown func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*SVGFEFuncAElement) Onmspointerenter() (onmspointerenter func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*SVGFEFuncAElement) SetOnmspointerenter(onmspointerenter func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*SVGFEFuncAElement) Onmspointerleave() (onmspointerleave func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*SVGFEFuncAElement) SetOnmspointerleave(onmspointerleave func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*SVGFEFuncAElement) Onmspointermove() (onmspointermove func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*SVGFEFuncAElement) SetOnmspointermove(onmspointermove func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*SVGFEFuncAElement) Onmspointerout() (onmspointerout func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*SVGFEFuncAElement) SetOnmspointerout(onmspointerout func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*SVGFEFuncAElement) Onmspointerover() (onmspointerover func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*SVGFEFuncAElement) SetOnmspointerover(onmspointerover func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*SVGFEFuncAElement) Onmspointerup() (onmspointerup func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*SVGFEFuncAElement) SetOnmspointerup(onmspointerup func(*window.MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*SVGFEFuncAElement) Ontouchcancel() (ontouchcancel func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*SVGFEFuncAElement) SetOntouchcancel(ontouchcancel func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*SVGFEFuncAElement) Ontouchend() (ontouchend func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*SVGFEFuncAElement) SetOntouchend(ontouchend func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*SVGFEFuncAElement) Ontouchmove() (ontouchmove func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*SVGFEFuncAElement) SetOntouchmove(ontouchmove func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*SVGFEFuncAElement) Ontouchstart() (ontouchstart func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*SVGFEFuncAElement) SetOntouchstart(ontouchstart func(*window.TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGFEFuncAElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGFEFuncAElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGFEFuncAElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGFEFuncAElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// OuterHTML prop
// js:"outerHTML"
func (*SVGFEFuncAElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*SVGFEFuncAElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

// Prefix prop
// js:"prefix"
func (*SVGFEFuncAElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

// ScrollHeight prop
// js:"scrollHeight"
func (*SVGFEFuncAElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*SVGFEFuncAElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*SVGFEFuncAElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*SVGFEFuncAElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*SVGFEFuncAElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ScrollWidth prop
// js:"scrollWidth"
func (*SVGFEFuncAElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

// TagName prop
// js:"tagName"
func (*SVGFEFuncAElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

// Onpointercancel prop
// js:"onpointercancel"
func (*SVGFEFuncAElement) Onpointercancel() (onpointercancel func(window.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*SVGFEFuncAElement) SetOnpointercancel(onpointercancel func(window.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*SVGFEFuncAElement) Onpointerdown() (onpointerdown func(window.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*SVGFEFuncAElement) SetOnpointerdown(onpointerdown func(window.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*SVGFEFuncAElement) Onpointerenter() (onpointerenter func(window.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*SVGFEFuncAElement) SetOnpointerenter(onpointerenter func(window.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*SVGFEFuncAElement) Onpointerleave() (onpointerleave func(window.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*SVGFEFuncAElement) SetOnpointerleave(onpointerleave func(window.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*SVGFEFuncAElement) Onpointermove() (onpointermove func(window.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*SVGFEFuncAElement) SetOnpointermove(onpointermove func(window.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*SVGFEFuncAElement) Onpointerout() (onpointerout func(window.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*SVGFEFuncAElement) SetOnpointerout(onpointerout func(window.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*SVGFEFuncAElement) Onpointerover() (onpointerover func(window.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*SVGFEFuncAElement) SetOnpointerover(onpointerover func(window.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*SVGFEFuncAElement) Onpointerup() (onpointerup func(window.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*SVGFEFuncAElement) SetOnpointerup(onpointerup func(window.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*SVGFEFuncAElement) Onwheel() (onwheel func(window.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*SVGFEFuncAElement) SetOnwheel(onwheel func(window.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

// ChildElementCount prop
// js:"childElementCount"
func (*SVGFEFuncAElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

// FirstElementChild prop
// js:"firstElementChild"
func (*SVGFEFuncAElement) FirstElementChild() (firstElementChild window.Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
// js:"lastElementChild"
func (*SVGFEFuncAElement) LastElementChild() (lastElementChild window.Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
// js:"nextElementSibling"
func (*SVGFEFuncAElement) NextElementSibling() (nextElementSibling window.Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
// js:"previousElementSibling"
func (*SVGFEFuncAElement) PreviousElementSibling() (previousElementSibling window.Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

// Attributes prop
// js:"attributes"
func (*SVGFEFuncAElement) Attributes() (attributes *window.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*SVGFEFuncAElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*SVGFEFuncAElement) ChildNodes() (childNodes *window.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*SVGFEFuncAElement) FirstChild() (firstChild window.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*SVGFEFuncAElement) LastChild() (lastChild window.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*SVGFEFuncAElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*SVGFEFuncAElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*SVGFEFuncAElement) NextSibling() (nextSibling window.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*SVGFEFuncAElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*SVGFEFuncAElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*SVGFEFuncAElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*SVGFEFuncAElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*SVGFEFuncAElement) OwnerDocument() (ownerDocument window.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*SVGFEFuncAElement) ParentElement() (parentElement window.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*SVGFEFuncAElement) ParentNode() (parentNode window.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*SVGFEFuncAElement) PreviousSibling() (previousSibling window.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*SVGFEFuncAElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*SVGFEFuncAElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
