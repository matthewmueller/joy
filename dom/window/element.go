package window

import (
	"github.com/matthewmueller/golly/dom/clientrect"
	"github.com/matthewmueller/golly/dom/clientrectlist"
	"github.com/matthewmueller/golly/dom/domtokenlist"
	"github.com/matthewmueller/golly/dom/mszoomtooptions"
)

// js:"Element,omit"
type Element interface {
	Node

	// GetAttribute
	// js:"getAttribute"
	GetAttribute(qualifiedName string) (s string)

	// GetAttributeNode
	// js:"getAttributeNode"
	GetAttributeNode(name string) (a *Attr)

	// GetAttributeNodeNS
	// js:"getAttributeNodeNS"
	GetAttributeNodeNS(namespaceURI string, localName string) (a *Attr)

	// GetAttributeNS
	// js:"getAttributeNS"
	GetAttributeNS(namespaceURI string, localName string) (s string)

	// GetBoundingClientRect
	// js:"getBoundingClientRect"
	GetBoundingClientRect() (c *clientrect.ClientRect)

	// GetClientRects
	// js:"getClientRects"
	GetClientRects() (c *clientrectlist.ClientRectList)

	// GetElementsByTagName
	// js:"getElementsByTagName"
	GetElementsByTagName(name string) (n *NodeList)

	// GetElementsByTagNameNS
	// js:"getElementsByTagNameNS"
	GetElementsByTagNameNS(namespaceURI string, localName string) (n *NodeList)

	// HasAttribute
	// js:"hasAttribute"
	HasAttribute(name string) (b bool)

	// HasAttributeNS
	// js:"hasAttributeNS"
	HasAttributeNS(namespaceURI string, localName string) (b bool)

	// MsGetRegionContent
	// js:"msGetRegionContent"
	MsGetRegionContent() (m *MSRangeCollection)

	// MsGetUntransformedBounds
	// js:"msGetUntransformedBounds"
	MsGetUntransformedBounds() (c *clientrect.ClientRect)

	// MsMatchesSelector
	// js:"msMatchesSelector"
	MsMatchesSelector(selectors string) (b bool)

	// MsReleasePointerCapture
	// js:"msReleasePointerCapture"
	MsReleasePointerCapture(pointerId int)

	// MsSetPointerCapture
	// js:"msSetPointerCapture"
	MsSetPointerCapture(pointerId int)

	// MsZoomTo
	// js:"msZoomTo"
	MsZoomTo(args *mszoomtooptions.MsZoomToOptions)

	// ReleasePointerCapture
	// js:"releasePointerCapture"
	ReleasePointerCapture(pointerId int)

	// RemoveAttribute
	// js:"removeAttribute"
	RemoveAttribute(qualifiedName string)

	// RemoveAttributeNode
	// js:"removeAttributeNode"
	RemoveAttributeNode(oldAttr *Attr) (a *Attr)

	// RemoveAttributeNS
	// js:"removeAttributeNS"
	RemoveAttributeNS(namespaceURI string, localName string)

	// RequestFullscreen
	// js:"requestFullscreen"
	RequestFullscreen()

	// RequestPointerLock
	// js:"requestPointerLock"
	RequestPointerLock()

	// SetAttribute
	// js:"setAttribute"
	SetAttribute(qualifiedName string, value string)

	// SetAttributeNode
	// js:"setAttributeNode"
	SetAttributeNode(newAttr *Attr) (a *Attr)

	// SetAttributeNodeNS
	// js:"setAttributeNodeNS"
	SetAttributeNodeNS(newAttr *Attr) (a *Attr)

	// SetAttributeNS
	// js:"setAttributeNS"
	SetAttributeNS(namespaceURI string, qualifiedName string, value string)

	// SetPointerCapture
	// js:"setPointerCapture"
	SetPointerCapture(pointerId int)

	// WebkitMatchesSelector
	// js:"webkitMatchesSelector"
	WebkitMatchesSelector(selectors string) (b bool)

	// WebkitRequestFullscreen
	// js:"webkitRequestFullscreen"
	WebkitRequestFullscreen()

	// WebkitRequestFullScreen
	// js:"webkitRequestFullScreen"
	WebkitRequestFullScreen()

	// ClassList prop
	// js:"classList"
	ClassList() (classList domtokenlist.DOMTokenList)

	// ClassName prop
	// js:"className"
	ClassName() (className string)

	// ClassName prop
	SetClassName(className string)

	// ClientHeight prop
	// js:"clientHeight"
	ClientHeight() (clientHeight int)

	// ClientLeft prop
	// js:"clientLeft"
	ClientLeft() (clientLeft int)

	// ClientTop prop
	// js:"clientTop"
	ClientTop() (clientTop int)

	// ClientWidth prop
	// js:"clientWidth"
	ClientWidth() (clientWidth int)

	// ID prop
	// js:"id"
	ID() (id string)

	// ID prop
	SetID(id string)

	// InnerHTML prop
	// js:"innerHTML"
	InnerHTML() (innerHTML string)

	// InnerHTML prop
	SetInnerHTML(innerHTML string)

	// MsContentZoomFactor prop
	// js:"msContentZoomFactor"
	MsContentZoomFactor() (msContentZoomFactor float32)

	// MsContentZoomFactor prop
	SetMsContentZoomFactor(msContentZoomFactor float32)

	// MsRegionOverflow prop
	// js:"msRegionOverflow"
	MsRegionOverflow() (msRegionOverflow string)

	// Onariarequest prop
	// js:"onariarequest"
	Onariarequest() (onariarequest func(Event))

	// Onariarequest prop
	SetOnariarequest(onariarequest func(Event))

	// Oncommand prop
	// js:"oncommand"
	Oncommand() (oncommand func(Event))

	// Oncommand prop
	SetOncommand(oncommand func(Event))

	// Ongotpointercapture prop
	// js:"ongotpointercapture"
	Ongotpointercapture() (ongotpointercapture func(*PointerEvent))

	// Ongotpointercapture prop
	SetOngotpointercapture(ongotpointercapture func(*PointerEvent))

	// Onlostpointercapture prop
	// js:"onlostpointercapture"
	Onlostpointercapture() (onlostpointercapture func(*PointerEvent))

	// Onlostpointercapture prop
	SetOnlostpointercapture(onlostpointercapture func(*PointerEvent))

	// Onmsgesturechange prop
	// js:"onmsgesturechange"
	Onmsgesturechange() (onmsgesturechange func(*MSGestureEvent))

	// Onmsgesturechange prop
	SetOnmsgesturechange(onmsgesturechange func(*MSGestureEvent))

	// Onmsgesturedoubletap prop
	// js:"onmsgesturedoubletap"
	Onmsgesturedoubletap() (onmsgesturedoubletap func(*MSGestureEvent))

	// Onmsgesturedoubletap prop
	SetOnmsgesturedoubletap(onmsgesturedoubletap func(*MSGestureEvent))

	// Onmsgestureend prop
	// js:"onmsgestureend"
	Onmsgestureend() (onmsgestureend func(*MSGestureEvent))

	// Onmsgestureend prop
	SetOnmsgestureend(onmsgestureend func(*MSGestureEvent))

	// Onmsgesturehold prop
	// js:"onmsgesturehold"
	Onmsgesturehold() (onmsgesturehold func(*MSGestureEvent))

	// Onmsgesturehold prop
	SetOnmsgesturehold(onmsgesturehold func(*MSGestureEvent))

	// Onmsgesturestart prop
	// js:"onmsgesturestart"
	Onmsgesturestart() (onmsgesturestart func(*MSGestureEvent))

	// Onmsgesturestart prop
	SetOnmsgesturestart(onmsgesturestart func(*MSGestureEvent))

	// Onmsgesturetap prop
	// js:"onmsgesturetap"
	Onmsgesturetap() (onmsgesturetap func(*MSGestureEvent))

	// Onmsgesturetap prop
	SetOnmsgesturetap(onmsgesturetap func(*MSGestureEvent))

	// Onmsgotpointercapture prop
	// js:"onmsgotpointercapture"
	Onmsgotpointercapture() (onmsgotpointercapture func(*MSPointerEvent))

	// Onmsgotpointercapture prop
	SetOnmsgotpointercapture(onmsgotpointercapture func(*MSPointerEvent))

	// Onmsinertiastart prop
	// js:"onmsinertiastart"
	Onmsinertiastart() (onmsinertiastart func(*MSGestureEvent))

	// Onmsinertiastart prop
	SetOnmsinertiastart(onmsinertiastart func(*MSGestureEvent))

	// Onmslostpointercapture prop
	// js:"onmslostpointercapture"
	Onmslostpointercapture() (onmslostpointercapture func(*MSPointerEvent))

	// Onmslostpointercapture prop
	SetOnmslostpointercapture(onmslostpointercapture func(*MSPointerEvent))

	// Onmspointercancel prop
	// js:"onmspointercancel"
	Onmspointercancel() (onmspointercancel func(*MSPointerEvent))

	// Onmspointercancel prop
	SetOnmspointercancel(onmspointercancel func(*MSPointerEvent))

	// Onmspointerdown prop
	// js:"onmspointerdown"
	Onmspointerdown() (onmspointerdown func(*MSPointerEvent))

	// Onmspointerdown prop
	SetOnmspointerdown(onmspointerdown func(*MSPointerEvent))

	// Onmspointerenter prop
	// js:"onmspointerenter"
	Onmspointerenter() (onmspointerenter func(*MSPointerEvent))

	// Onmspointerenter prop
	SetOnmspointerenter(onmspointerenter func(*MSPointerEvent))

	// Onmspointerleave prop
	// js:"onmspointerleave"
	Onmspointerleave() (onmspointerleave func(*MSPointerEvent))

	// Onmspointerleave prop
	SetOnmspointerleave(onmspointerleave func(*MSPointerEvent))

	// Onmspointermove prop
	// js:"onmspointermove"
	Onmspointermove() (onmspointermove func(*MSPointerEvent))

	// Onmspointermove prop
	SetOnmspointermove(onmspointermove func(*MSPointerEvent))

	// Onmspointerout prop
	// js:"onmspointerout"
	Onmspointerout() (onmspointerout func(*MSPointerEvent))

	// Onmspointerout prop
	SetOnmspointerout(onmspointerout func(*MSPointerEvent))

	// Onmspointerover prop
	// js:"onmspointerover"
	Onmspointerover() (onmspointerover func(*MSPointerEvent))

	// Onmspointerover prop
	SetOnmspointerover(onmspointerover func(*MSPointerEvent))

	// Onmspointerup prop
	// js:"onmspointerup"
	Onmspointerup() (onmspointerup func(*MSPointerEvent))

	// Onmspointerup prop
	SetOnmspointerup(onmspointerup func(*MSPointerEvent))

	// Ontouchcancel prop
	// js:"ontouchcancel"
	Ontouchcancel() (ontouchcancel func(*TouchEvent))

	// Ontouchcancel prop
	SetOntouchcancel(ontouchcancel func(*TouchEvent))

	// Ontouchend prop
	// js:"ontouchend"
	Ontouchend() (ontouchend func(*TouchEvent))

	// Ontouchend prop
	SetOntouchend(ontouchend func(*TouchEvent))

	// Ontouchmove prop
	// js:"ontouchmove"
	Ontouchmove() (ontouchmove func(*TouchEvent))

	// Ontouchmove prop
	SetOntouchmove(ontouchmove func(*TouchEvent))

	// Ontouchstart prop
	// js:"ontouchstart"
	Ontouchstart() (ontouchstart func(*TouchEvent))

	// Ontouchstart prop
	SetOntouchstart(ontouchstart func(*TouchEvent))

	// Onwebkitfullscreenchange prop
	// js:"onwebkitfullscreenchange"
	Onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event))

	// Onwebkitfullscreenchange prop
	SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(Event))

	// Onwebkitfullscreenerror prop
	// js:"onwebkitfullscreenerror"
	Onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event))

	// Onwebkitfullscreenerror prop
	SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(Event))

	// OuterHTML prop
	// js:"outerHTML"
	OuterHTML() (outerHTML string)

	// OuterHTML prop
	SetOuterHTML(outerHTML string)

	// Prefix prop
	// js:"prefix"
	Prefix() (prefix string)

	// ScrollHeight prop
	// js:"scrollHeight"
	ScrollHeight() (scrollHeight int)

	// ScrollLeft prop
	// js:"scrollLeft"
	ScrollLeft() (scrollLeft int)

	// ScrollLeft prop
	SetScrollLeft(scrollLeft int)

	// ScrollTop prop
	// js:"scrollTop"
	ScrollTop() (scrollTop int)

	// ScrollTop prop
	SetScrollTop(scrollTop int)

	// ScrollWidth prop
	// js:"scrollWidth"
	ScrollWidth() (scrollWidth int)

	// TagName prop
	// js:"tagName"
	TagName() (tagName string)
}
