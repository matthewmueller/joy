package svgfefuncrelement

import (
	"github.com/matthewmueller/golly/dom/childnode"
	"github.com/matthewmueller/golly/dom/clientrect"
	"github.com/matthewmueller/golly/dom/clientrectlist"
	"github.com/matthewmueller/golly/dom/domstringmap"
	"github.com/matthewmueller/golly/dom/domtokenlist"
	"github.com/matthewmueller/golly/dom/mszoomtooptions"
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatednumber"
	"github.com/matthewmueller/golly/dom/svganimatednumberlist"
	"github.com/matthewmueller/golly/dom/svgcomponenttransferfunctionelement"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ svgcomponenttransferfunctionelement.SVGComponentTransferFunctionElement = (*SVGFEFuncRElement)(nil)
var _ window.SVGElement = (*SVGFEFuncRElement)(nil)
var _ window.Element = (*SVGFEFuncRElement)(nil)
var _ window.GlobalEventHandlers = (*SVGFEFuncRElement)(nil)
var _ window.ElementTraversal = (*SVGFEFuncRElement)(nil)
var _ window.NodeSelector = (*SVGFEFuncRElement)(nil)
var _ childnode.ChildNode = (*SVGFEFuncRElement)(nil)
var _ window.Node = (*SVGFEFuncRElement)(nil)
var _ window.EventTarget = (*SVGFEFuncRElement)(nil)

// SVGFEFuncRElement struct
// js:"SVGFEFuncRElement,omit"
type SVGFEFuncRElement struct {
}

// GetAttribute fn
// js:"getAttribute"
func (*SVGFEFuncRElement) GetAttribute(qualifiedName string) (s string) {
	js.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

// GetAttributeNode fn
// js:"getAttributeNode"
func (*SVGFEFuncRElement) GetAttributeNode(name string) (w *window.Attr) {
	js.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

// GetAttributeNodeNS fn
// js:"getAttributeNodeNS"
func (*SVGFEFuncRElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *window.Attr) {
	js.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

// GetAttributeNS fn
// js:"getAttributeNS"
func (*SVGFEFuncRElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	js.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*SVGFEFuncRElement) GetBoundingClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*SVGFEFuncRElement) GetClientRects() (c *clientrectlist.ClientRectList) {
	js.Rewrite("$_.getClientRects()")
	return c
}

// GetElementsByTagName fn
// js:"getElementsByTagName"
func (*SVGFEFuncRElement) GetElementsByTagName(name string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*SVGFEFuncRElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

// HasAttribute fn
// js:"hasAttribute"
func (*SVGFEFuncRElement) HasAttribute(name string) (b bool) {
	js.Rewrite("$_.hasAttribute($1)", name)
	return b
}

// HasAttributeNS fn
// js:"hasAttributeNS"
func (*SVGFEFuncRElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	js.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// MsGetRegionContent fn
// js:"msGetRegionContent"
func (*SVGFEFuncRElement) MsGetRegionContent() (w *window.MSRangeCollection) {
	js.Rewrite("$_.msGetRegionContent()")
	return w
}

// MsGetUntransformedBounds fn
// js:"msGetUntransformedBounds"
func (*SVGFEFuncRElement) MsGetUntransformedBounds() (c *clientrect.ClientRect) {
	js.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

// MsMatchesSelector fn
// js:"msMatchesSelector"
func (*SVGFEFuncRElement) MsMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

// MsReleasePointerCapture fn
// js:"msReleasePointerCapture"
func (*SVGFEFuncRElement) MsReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

// MsSetPointerCapture fn
// js:"msSetPointerCapture"
func (*SVGFEFuncRElement) MsSetPointerCapture(pointerId int) {
	js.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

// MsZoomTo fn
// js:"msZoomTo"
func (*SVGFEFuncRElement) MsZoomTo(args *mszoomtooptions.MsZoomToOptions) {
	js.Rewrite("$_.msZoomTo($1)", args)
}

// ReleasePointerCapture fn
// js:"releasePointerCapture"
func (*SVGFEFuncRElement) ReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

// RemoveAttribute fn
// js:"removeAttribute"
func (*SVGFEFuncRElement) RemoveAttribute(qualifiedName string) {
	js.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

// RemoveAttributeNode fn
// js:"removeAttributeNode"
func (*SVGFEFuncRElement) RemoveAttributeNode(oldAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

// RemoveAttributeNS fn
// js:"removeAttributeNS"
func (*SVGFEFuncRElement) RemoveAttributeNS(namespaceURI string, localName string) {
	js.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// RequestFullscreen fn
// js:"requestFullscreen"
func (*SVGFEFuncRElement) RequestFullscreen() {
	js.Rewrite("$_.requestFullscreen()")
}

// RequestPointerLock fn
// js:"requestPointerLock"
func (*SVGFEFuncRElement) RequestPointerLock() {
	js.Rewrite("$_.requestPointerLock()")
}

// SetAttribute fn
// js:"setAttribute"
func (*SVGFEFuncRElement) SetAttribute(qualifiedName string, value string) {
	js.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

// SetAttributeNode fn
// js:"setAttributeNode"
func (*SVGFEFuncRElement) SetAttributeNode(newAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

// SetAttributeNodeNS fn
// js:"setAttributeNodeNS"
func (*SVGFEFuncRElement) SetAttributeNodeNS(newAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

// SetAttributeNS fn
// js:"setAttributeNS"
func (*SVGFEFuncRElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	js.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// SetPointerCapture fn
// js:"setPointerCapture"
func (*SVGFEFuncRElement) SetPointerCapture(pointerId int) {
	js.Rewrite("$_.setPointerCapture($1)", pointerId)
}

// WebkitMatchesSelector fn
// js:"webkitMatchesSelector"
func (*SVGFEFuncRElement) WebkitMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

// WebkitRequestFullscreen fn
// js:"webkitRequestFullscreen"
func (*SVGFEFuncRElement) WebkitRequestFullscreen() {
	js.Rewrite("$_.webkitRequestFullscreen()")
}

// WebkitRequestFullScreen fn
// js:"webkitRequestFullScreen"
func (*SVGFEFuncRElement) WebkitRequestFullScreen() {
	js.Rewrite("$_.webkitRequestFullScreen()")
}

// QuerySelector fn
// js:"querySelector"
func (*SVGFEFuncRElement) QuerySelector(selectors string) (w window.Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return w
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*SVGFEFuncRElement) QuerySelectorAll(selectors string) (w *window.NodeList) {
	js.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

// AppendChild fn
// js:"appendChild"
func (*SVGFEFuncRElement) AppendChild(newChild window.Node) (w window.Node) {
	js.Rewrite("$_.appendChild($1)", newChild)
	return w
}

// CloneNode fn
// js:"cloneNode"
func (*SVGFEFuncRElement) CloneNode(deep *bool) (w window.Node) {
	js.Rewrite("$_.cloneNode($1)", deep)
	return w
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*SVGFEFuncRElement) CompareDocumentPosition(other window.Node) (u uint8) {
	js.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*SVGFEFuncRElement) Contains(child window.Node) (b bool) {
	js.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*SVGFEFuncRElement) HasAttributes() (b bool) {
	js.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*SVGFEFuncRElement) HasChildNodes() (b bool) {
	js.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*SVGFEFuncRElement) InsertBefore(newChild window.Node, refChild *window.Node) (w window.Node) {
	js.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*SVGFEFuncRElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	js.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*SVGFEFuncRElement) IsEqualNode(arg window.Node) (b bool) {
	js.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*SVGFEFuncRElement) IsSameNode(other window.Node) (b bool) {
	js.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*SVGFEFuncRElement) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*SVGFEFuncRElement) LookupPrefix(namespaceURI string) (s string) {
	js.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*SVGFEFuncRElement) Normalize() {
	js.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*SVGFEFuncRElement) RemoveChild(oldChild window.Node) (w window.Node) {
	js.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

// ReplaceChild fn
// js:"replaceChild"
func (*SVGFEFuncRElement) ReplaceChild(newChild window.Node, oldChild window.Node) (w window.Node) {
	js.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*SVGFEFuncRElement) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SVGFEFuncRElement) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SVGFEFuncRElement) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Amplitude prop
// js:"amplitude"
func (*SVGFEFuncRElement) Amplitude() (amplitude *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$_.amplitude")
	return amplitude
}

// Exponent prop
// js:"exponent"
func (*SVGFEFuncRElement) Exponent() (exponent *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$_.exponent")
	return exponent
}

// Intercept prop
// js:"intercept"
func (*SVGFEFuncRElement) Intercept() (intercept *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$_.intercept")
	return intercept
}

// Offset prop
// js:"offset"
func (*SVGFEFuncRElement) Offset() (offset *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$_.offset")
	return offset
}

// Slope prop
// js:"slope"
func (*SVGFEFuncRElement) Slope() (slope *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$_.slope")
	return slope
}

// TableValues prop
// js:"tableValues"
func (*SVGFEFuncRElement) TableValues() (tableValues *svganimatednumberlist.SVGAnimatedNumberList) {
	js.Rewrite("$_.tableValues")
	return tableValues
}

// Type prop
// js:"type"
func (*SVGFEFuncRElement) Type() (kind *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$_.type")
	return kind
}

// Dataset prop
// js:"dataset"
func (*SVGFEFuncRElement) Dataset() (dataset *domstringmap.DOMStringMap) {
	js.Rewrite("$_.dataset")
	return dataset
}

// OwnerSVGElement prop
// js:"ownerSVGElement"
func (*SVGFEFuncRElement) OwnerSVGElement() (ownerSVGElement *window.SVGSVGElement) {
	js.Rewrite("$_.ownerSVGElement")
	return ownerSVGElement
}

// ViewportElement prop
// js:"viewportElement"
func (*SVGFEFuncRElement) ViewportElement() (viewportElement window.SVGElement) {
	js.Rewrite("$_.viewportElement")
	return viewportElement
}

// Xmlbase prop
// js:"xmlbase"
func (*SVGFEFuncRElement) Xmlbase() (xmlbase string) {
	js.Rewrite("$_.xmlbase")
	return xmlbase
}

// SetXmlbase prop
// js:"xmlbase"
func (*SVGFEFuncRElement) SetXmlbase(xmlbase string) {
	js.Rewrite("$_.xmlbase = $1", xmlbase)
}

// ClassList prop
// js:"classList"
func (*SVGFEFuncRElement) ClassList() (classList domtokenlist.DOMTokenList) {
	js.Rewrite("$_.classList")
	return classList
}

// ClassName prop
// js:"className"
func (*SVGFEFuncRElement) ClassName() (className string) {
	js.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*SVGFEFuncRElement) SetClassName(className string) {
	js.Rewrite("$_.className = $1", className)
}

// ClientHeight prop
// js:"clientHeight"
func (*SVGFEFuncRElement) ClientHeight() (clientHeight int) {
	js.Rewrite("$_.clientHeight")
	return clientHeight
}

// ClientLeft prop
// js:"clientLeft"
func (*SVGFEFuncRElement) ClientLeft() (clientLeft int) {
	js.Rewrite("$_.clientLeft")
	return clientLeft
}

// ClientTop prop
// js:"clientTop"
func (*SVGFEFuncRElement) ClientTop() (clientTop int) {
	js.Rewrite("$_.clientTop")
	return clientTop
}

// ClientWidth prop
// js:"clientWidth"
func (*SVGFEFuncRElement) ClientWidth() (clientWidth int) {
	js.Rewrite("$_.clientWidth")
	return clientWidth
}

// ID prop
// js:"id"
func (*SVGFEFuncRElement) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*SVGFEFuncRElement) SetID(id string) {
	js.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*SVGFEFuncRElement) InnerHTML() (innerHTML string) {
	js.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*SVGFEFuncRElement) SetInnerHTML(innerHTML string) {
	js.Rewrite("$_.innerHTML = $1", innerHTML)
}

// MsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGFEFuncRElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

// SetMsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*SVGFEFuncRElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

// MsRegionOverflow prop
// js:"msRegionOverflow"
func (*SVGFEFuncRElement) MsRegionOverflow() (msRegionOverflow string) {
	js.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

// Onariarequest prop
// js:"onariarequest"
func (*SVGFEFuncRElement) Onariarequest() (onariarequest func(window.Event)) {
	js.Rewrite("$_.onariarequest")
	return onariarequest
}

// SetOnariarequest prop
// js:"onariarequest"
func (*SVGFEFuncRElement) SetOnariarequest(onariarequest func(window.Event)) {
	js.Rewrite("$_.onariarequest = $1", onariarequest)
}

// Oncommand prop
// js:"oncommand"
func (*SVGFEFuncRElement) Oncommand() (oncommand func(window.Event)) {
	js.Rewrite("$_.oncommand")
	return oncommand
}

// SetOncommand prop
// js:"oncommand"
func (*SVGFEFuncRElement) SetOncommand(oncommand func(window.Event)) {
	js.Rewrite("$_.oncommand = $1", oncommand)
}

// Ongotpointercapture prop
// js:"ongotpointercapture"
func (*SVGFEFuncRElement) Ongotpointercapture() (ongotpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

// SetOngotpointercapture prop
// js:"ongotpointercapture"
func (*SVGFEFuncRElement) SetOngotpointercapture(ongotpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

// Onlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGFEFuncRElement) Onlostpointercapture() (onlostpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

// SetOnlostpointercapture prop
// js:"onlostpointercapture"
func (*SVGFEFuncRElement) SetOnlostpointercapture(onlostpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGFEFuncRElement) Onmsgesturechange() (onmsgesturechange func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*SVGFEFuncRElement) SetOnmsgesturechange(onmsgesturechange func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGFEFuncRElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*SVGFEFuncRElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*SVGFEFuncRElement) Onmsgestureend() (onmsgestureend func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*SVGFEFuncRElement) SetOnmsgestureend(onmsgestureend func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGFEFuncRElement) Onmsgesturehold() (onmsgesturehold func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*SVGFEFuncRElement) SetOnmsgesturehold(onmsgesturehold func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGFEFuncRElement) Onmsgesturestart() (onmsgesturestart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*SVGFEFuncRElement) SetOnmsgesturestart(onmsgesturestart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGFEFuncRElement) Onmsgesturetap() (onmsgesturetap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*SVGFEFuncRElement) SetOnmsgesturetap(onmsgesturetap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGFEFuncRElement) Onmsgotpointercapture() (onmsgotpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

// SetOnmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*SVGFEFuncRElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGFEFuncRElement) Onmsinertiastart() (onmsinertiastart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*SVGFEFuncRElement) SetOnmsinertiastart(onmsinertiastart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGFEFuncRElement) Onmslostpointercapture() (onmslostpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

// SetOnmslostpointercapture prop
// js:"onmslostpointercapture"
func (*SVGFEFuncRElement) SetOnmslostpointercapture(onmslostpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*SVGFEFuncRElement) Onmspointercancel() (onmspointercancel func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*SVGFEFuncRElement) SetOnmspointercancel(onmspointercancel func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*SVGFEFuncRElement) Onmspointerdown() (onmspointerdown func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*SVGFEFuncRElement) SetOnmspointerdown(onmspointerdown func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*SVGFEFuncRElement) Onmspointerenter() (onmspointerenter func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*SVGFEFuncRElement) SetOnmspointerenter(onmspointerenter func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*SVGFEFuncRElement) Onmspointerleave() (onmspointerleave func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*SVGFEFuncRElement) SetOnmspointerleave(onmspointerleave func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*SVGFEFuncRElement) Onmspointermove() (onmspointermove func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*SVGFEFuncRElement) SetOnmspointermove(onmspointermove func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*SVGFEFuncRElement) Onmspointerout() (onmspointerout func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*SVGFEFuncRElement) SetOnmspointerout(onmspointerout func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*SVGFEFuncRElement) Onmspointerover() (onmspointerover func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*SVGFEFuncRElement) SetOnmspointerover(onmspointerover func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*SVGFEFuncRElement) Onmspointerup() (onmspointerup func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*SVGFEFuncRElement) SetOnmspointerup(onmspointerup func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*SVGFEFuncRElement) Ontouchcancel() (ontouchcancel func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*SVGFEFuncRElement) SetOntouchcancel(ontouchcancel func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*SVGFEFuncRElement) Ontouchend() (ontouchend func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*SVGFEFuncRElement) SetOntouchend(ontouchend func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*SVGFEFuncRElement) Ontouchmove() (ontouchmove func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*SVGFEFuncRElement) SetOntouchmove(ontouchmove func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*SVGFEFuncRElement) Ontouchstart() (ontouchstart func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*SVGFEFuncRElement) SetOntouchstart(ontouchstart func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGFEFuncRElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*SVGFEFuncRElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGFEFuncRElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*SVGFEFuncRElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// OuterHTML prop
// js:"outerHTML"
func (*SVGFEFuncRElement) OuterHTML() (outerHTML string) {
	js.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*SVGFEFuncRElement) SetOuterHTML(outerHTML string) {
	js.Rewrite("$_.outerHTML = $1", outerHTML)
}

// Prefix prop
// js:"prefix"
func (*SVGFEFuncRElement) Prefix() (prefix string) {
	js.Rewrite("$_.prefix")
	return prefix
}

// ScrollHeight prop
// js:"scrollHeight"
func (*SVGFEFuncRElement) ScrollHeight() (scrollHeight int) {
	js.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*SVGFEFuncRElement) ScrollLeft() (scrollLeft int) {
	js.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*SVGFEFuncRElement) SetScrollLeft(scrollLeft int) {
	js.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*SVGFEFuncRElement) ScrollTop() (scrollTop int) {
	js.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*SVGFEFuncRElement) SetScrollTop(scrollTop int) {
	js.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ScrollWidth prop
// js:"scrollWidth"
func (*SVGFEFuncRElement) ScrollWidth() (scrollWidth int) {
	js.Rewrite("$_.scrollWidth")
	return scrollWidth
}

// TagName prop
// js:"tagName"
func (*SVGFEFuncRElement) TagName() (tagName string) {
	js.Rewrite("$_.tagName")
	return tagName
}

// Onpointercancel prop
// js:"onpointercancel"
func (*SVGFEFuncRElement) Onpointercancel() (onpointercancel func(window.Event)) {
	js.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*SVGFEFuncRElement) SetOnpointercancel(onpointercancel func(window.Event)) {
	js.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*SVGFEFuncRElement) Onpointerdown() (onpointerdown func(window.Event)) {
	js.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*SVGFEFuncRElement) SetOnpointerdown(onpointerdown func(window.Event)) {
	js.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*SVGFEFuncRElement) Onpointerenter() (onpointerenter func(window.Event)) {
	js.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*SVGFEFuncRElement) SetOnpointerenter(onpointerenter func(window.Event)) {
	js.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*SVGFEFuncRElement) Onpointerleave() (onpointerleave func(window.Event)) {
	js.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*SVGFEFuncRElement) SetOnpointerleave(onpointerleave func(window.Event)) {
	js.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*SVGFEFuncRElement) Onpointermove() (onpointermove func(window.Event)) {
	js.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*SVGFEFuncRElement) SetOnpointermove(onpointermove func(window.Event)) {
	js.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*SVGFEFuncRElement) Onpointerout() (onpointerout func(window.Event)) {
	js.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*SVGFEFuncRElement) SetOnpointerout(onpointerout func(window.Event)) {
	js.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*SVGFEFuncRElement) Onpointerover() (onpointerover func(window.Event)) {
	js.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*SVGFEFuncRElement) SetOnpointerover(onpointerover func(window.Event)) {
	js.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*SVGFEFuncRElement) Onpointerup() (onpointerup func(window.Event)) {
	js.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*SVGFEFuncRElement) SetOnpointerup(onpointerup func(window.Event)) {
	js.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*SVGFEFuncRElement) Onwheel() (onwheel func(window.Event)) {
	js.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*SVGFEFuncRElement) SetOnwheel(onwheel func(window.Event)) {
	js.Rewrite("$_.onwheel = $1", onwheel)
}

// ChildElementCount prop
// js:"childElementCount"
func (*SVGFEFuncRElement) ChildElementCount() (childElementCount uint) {
	js.Rewrite("$_.childElementCount")
	return childElementCount
}

// FirstElementChild prop
// js:"firstElementChild"
func (*SVGFEFuncRElement) FirstElementChild() (firstElementChild window.Element) {
	js.Rewrite("$_.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
// js:"lastElementChild"
func (*SVGFEFuncRElement) LastElementChild() (lastElementChild window.Element) {
	js.Rewrite("$_.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
// js:"nextElementSibling"
func (*SVGFEFuncRElement) NextElementSibling() (nextElementSibling window.Element) {
	js.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
// js:"previousElementSibling"
func (*SVGFEFuncRElement) PreviousElementSibling() (previousElementSibling window.Element) {
	js.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

// Attributes prop
// js:"attributes"
func (*SVGFEFuncRElement) Attributes() (attributes *window.NamedNodeMap) {
	js.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*SVGFEFuncRElement) BaseURI() (baseURI string) {
	js.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*SVGFEFuncRElement) ChildNodes() (childNodes *window.NodeList) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*SVGFEFuncRElement) FirstChild() (firstChild window.Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*SVGFEFuncRElement) LastChild() (lastChild window.Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*SVGFEFuncRElement) LocalName() (localName string) {
	js.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*SVGFEFuncRElement) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*SVGFEFuncRElement) NextSibling() (nextSibling window.Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*SVGFEFuncRElement) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*SVGFEFuncRElement) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*SVGFEFuncRElement) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*SVGFEFuncRElement) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*SVGFEFuncRElement) OwnerDocument() (ownerDocument window.Document) {
	js.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*SVGFEFuncRElement) ParentElement() (parentElement window.HTMLElement) {
	js.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*SVGFEFuncRElement) ParentNode() (parentNode window.Node) {
	js.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*SVGFEFuncRElement) PreviousSibling() (previousSibling window.Node) {
	js.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*SVGFEFuncRElement) TextContent() (textContent string) {
	js.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*SVGFEFuncRElement) SetTextContent(textContent string) {
	js.Rewrite("$_.textContent = $1", textContent)
}
