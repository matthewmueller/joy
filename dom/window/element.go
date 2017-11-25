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

	// QuerySelector
	// js:"querySelector"
	QuerySelector(selectors string) (e Element)

	// QuerySelectorAll
	// js:"querySelectorAll"
	QuerySelectorAll(selectors string) (n *NodeList)

	// Remove
	// js:"remove"
	Remove()

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

	// Onpointercancel prop
	// js:"onpointercancel"
	Onpointercancel() (onpointercancel func(Event))

	// Onpointercancel prop
	// js:"setonpointercancel"
	SetOnpointercancel(onpointercancel func(Event))

	// Onpointerdown prop
	// js:"onpointerdown"
	Onpointerdown() (onpointerdown func(Event))

	// Onpointerdown prop
	// js:"setonpointerdown"
	SetOnpointerdown(onpointerdown func(Event))

	// Onpointerenter prop
	// js:"onpointerenter"
	Onpointerenter() (onpointerenter func(Event))

	// Onpointerenter prop
	// js:"setonpointerenter"
	SetOnpointerenter(onpointerenter func(Event))

	// Onpointerleave prop
	// js:"onpointerleave"
	Onpointerleave() (onpointerleave func(Event))

	// Onpointerleave prop
	// js:"setonpointerleave"
	SetOnpointerleave(onpointerleave func(Event))

	// Onpointermove prop
	// js:"onpointermove"
	Onpointermove() (onpointermove func(Event))

	// Onpointermove prop
	// js:"setonpointermove"
	SetOnpointermove(onpointermove func(Event))

	// Onpointerout prop
	// js:"onpointerout"
	Onpointerout() (onpointerout func(Event))

	// Onpointerout prop
	// js:"setonpointerout"
	SetOnpointerout(onpointerout func(Event))

	// Onpointerover prop
	// js:"onpointerover"
	Onpointerover() (onpointerover func(Event))

	// Onpointerover prop
	// js:"setonpointerover"
	SetOnpointerover(onpointerover func(Event))

	// Onpointerup prop
	// js:"onpointerup"
	Onpointerup() (onpointerup func(Event))

	// Onpointerup prop
	// js:"setonpointerup"
	SetOnpointerup(onpointerup func(Event))

	// Onwheel prop
	// js:"onwheel"
	Onwheel() (onwheel func(Event))

	// Onwheel prop
	// js:"setonwheel"
	SetOnwheel(onwheel func(Event))

	// ChildElementCount prop
	// js:"childElementCount"
	ChildElementCount() (childElementCount uint)

	// FirstElementChild prop
	// js:"firstElementChild"
	FirstElementChild() (firstElementChild Element)

	// LastElementChild prop
	// js:"lastElementChild"
	LastElementChild() (lastElementChild Element)

	// NextElementSibling prop
	// js:"nextElementSibling"
	NextElementSibling() (nextElementSibling Element)

	// PreviousElementSibling prop
	// js:"previousElementSibling"
	PreviousElementSibling() (previousElementSibling Element)

	// ClassList prop
	// js:"classList"
	ClassList() (classList domtokenlist.DOMTokenList)

	// ClassName prop
	// js:"className"
	ClassName() (className string)

	// ClassName prop
	// js:"setclassName"
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
	// js:"setid"
	SetID(id string)

	// InnerHTML prop
	// js:"innerHTML"
	InnerHTML() (innerHTML string)

	// InnerHTML prop
	// js:"setinnerHTML"
	SetInnerHTML(innerHTML string)

	// MsContentZoomFactor prop
	// js:"msContentZoomFactor"
	MsContentZoomFactor() (msContentZoomFactor float32)

	// MsContentZoomFactor prop
	// js:"setmsContentZoomFactor"
	SetMsContentZoomFactor(msContentZoomFactor float32)

	// MsRegionOverflow prop
	// js:"msRegionOverflow"
	MsRegionOverflow() (msRegionOverflow string)

	// Onariarequest prop
	// js:"onariarequest"
	Onariarequest() (onariarequest func(Event))

	// Onariarequest prop
	// js:"setonariarequest"
	SetOnariarequest(onariarequest func(Event))

	// Oncommand prop
	// js:"oncommand"
	Oncommand() (oncommand func(Event))

	// Oncommand prop
	// js:"setoncommand"
	SetOncommand(oncommand func(Event))

	// Ongotpointercapture prop
	// js:"ongotpointercapture"
	Ongotpointercapture() (ongotpointercapture func(*PointerEvent))

	// Ongotpointercapture prop
	// js:"setongotpointercapture"
	SetOngotpointercapture(ongotpointercapture func(*PointerEvent))

	// Onlostpointercapture prop
	// js:"onlostpointercapture"
	Onlostpointercapture() (onlostpointercapture func(*PointerEvent))

	// Onlostpointercapture prop
	// js:"setonlostpointercapture"
	SetOnlostpointercapture(onlostpointercapture func(*PointerEvent))

	// Onmsgesturechange prop
	// js:"onmsgesturechange"
	Onmsgesturechange() (onmsgesturechange func(*MSGestureEvent))

	// Onmsgesturechange prop
	// js:"setonmsgesturechange"
	SetOnmsgesturechange(onmsgesturechange func(*MSGestureEvent))

	// Onmsgesturedoubletap prop
	// js:"onmsgesturedoubletap"
	Onmsgesturedoubletap() (onmsgesturedoubletap func(*MSGestureEvent))

	// Onmsgesturedoubletap prop
	// js:"setonmsgesturedoubletap"
	SetOnmsgesturedoubletap(onmsgesturedoubletap func(*MSGestureEvent))

	// Onmsgestureend prop
	// js:"onmsgestureend"
	Onmsgestureend() (onmsgestureend func(*MSGestureEvent))

	// Onmsgestureend prop
	// js:"setonmsgestureend"
	SetOnmsgestureend(onmsgestureend func(*MSGestureEvent))

	// Onmsgesturehold prop
	// js:"onmsgesturehold"
	Onmsgesturehold() (onmsgesturehold func(*MSGestureEvent))

	// Onmsgesturehold prop
	// js:"setonmsgesturehold"
	SetOnmsgesturehold(onmsgesturehold func(*MSGestureEvent))

	// Onmsgesturestart prop
	// js:"onmsgesturestart"
	Onmsgesturestart() (onmsgesturestart func(*MSGestureEvent))

	// Onmsgesturestart prop
	// js:"setonmsgesturestart"
	SetOnmsgesturestart(onmsgesturestart func(*MSGestureEvent))

	// Onmsgesturetap prop
	// js:"onmsgesturetap"
	Onmsgesturetap() (onmsgesturetap func(*MSGestureEvent))

	// Onmsgesturetap prop
	// js:"setonmsgesturetap"
	SetOnmsgesturetap(onmsgesturetap func(*MSGestureEvent))

	// Onmsgotpointercapture prop
	// js:"onmsgotpointercapture"
	Onmsgotpointercapture() (onmsgotpointercapture func(*MSPointerEvent))

	// Onmsgotpointercapture prop
	// js:"setonmsgotpointercapture"
	SetOnmsgotpointercapture(onmsgotpointercapture func(*MSPointerEvent))

	// Onmsinertiastart prop
	// js:"onmsinertiastart"
	Onmsinertiastart() (onmsinertiastart func(*MSGestureEvent))

	// Onmsinertiastart prop
	// js:"setonmsinertiastart"
	SetOnmsinertiastart(onmsinertiastart func(*MSGestureEvent))

	// Onmslostpointercapture prop
	// js:"onmslostpointercapture"
	Onmslostpointercapture() (onmslostpointercapture func(*MSPointerEvent))

	// Onmslostpointercapture prop
	// js:"setonmslostpointercapture"
	SetOnmslostpointercapture(onmslostpointercapture func(*MSPointerEvent))

	// Onmspointercancel prop
	// js:"onmspointercancel"
	Onmspointercancel() (onmspointercancel func(*MSPointerEvent))

	// Onmspointercancel prop
	// js:"setonmspointercancel"
	SetOnmspointercancel(onmspointercancel func(*MSPointerEvent))

	// Onmspointerdown prop
	// js:"onmspointerdown"
	Onmspointerdown() (onmspointerdown func(*MSPointerEvent))

	// Onmspointerdown prop
	// js:"setonmspointerdown"
	SetOnmspointerdown(onmspointerdown func(*MSPointerEvent))

	// Onmspointerenter prop
	// js:"onmspointerenter"
	Onmspointerenter() (onmspointerenter func(*MSPointerEvent))

	// Onmspointerenter prop
	// js:"setonmspointerenter"
	SetOnmspointerenter(onmspointerenter func(*MSPointerEvent))

	// Onmspointerleave prop
	// js:"onmspointerleave"
	Onmspointerleave() (onmspointerleave func(*MSPointerEvent))

	// Onmspointerleave prop
	// js:"setonmspointerleave"
	SetOnmspointerleave(onmspointerleave func(*MSPointerEvent))

	// Onmspointermove prop
	// js:"onmspointermove"
	Onmspointermove() (onmspointermove func(*MSPointerEvent))

	// Onmspointermove prop
	// js:"setonmspointermove"
	SetOnmspointermove(onmspointermove func(*MSPointerEvent))

	// Onmspointerout prop
	// js:"onmspointerout"
	Onmspointerout() (onmspointerout func(*MSPointerEvent))

	// Onmspointerout prop
	// js:"setonmspointerout"
	SetOnmspointerout(onmspointerout func(*MSPointerEvent))

	// Onmspointerover prop
	// js:"onmspointerover"
	Onmspointerover() (onmspointerover func(*MSPointerEvent))

	// Onmspointerover prop
	// js:"setonmspointerover"
	SetOnmspointerover(onmspointerover func(*MSPointerEvent))

	// Onmspointerup prop
	// js:"onmspointerup"
	Onmspointerup() (onmspointerup func(*MSPointerEvent))

	// Onmspointerup prop
	// js:"setonmspointerup"
	SetOnmspointerup(onmspointerup func(*MSPointerEvent))

	// Ontouchcancel prop
	// js:"ontouchcancel"
	Ontouchcancel() (ontouchcancel func(*TouchEvent))

	// Ontouchcancel prop
	// js:"setontouchcancel"
	SetOntouchcancel(ontouchcancel func(*TouchEvent))

	// Ontouchend prop
	// js:"ontouchend"
	Ontouchend() (ontouchend func(*TouchEvent))

	// Ontouchend prop
	// js:"setontouchend"
	SetOntouchend(ontouchend func(*TouchEvent))

	// Ontouchmove prop
	// js:"ontouchmove"
	Ontouchmove() (ontouchmove func(*TouchEvent))

	// Ontouchmove prop
	// js:"setontouchmove"
	SetOntouchmove(ontouchmove func(*TouchEvent))

	// Ontouchstart prop
	// js:"ontouchstart"
	Ontouchstart() (ontouchstart func(*TouchEvent))

	// Ontouchstart prop
	// js:"setontouchstart"
	SetOntouchstart(ontouchstart func(*TouchEvent))

	// Onwebkitfullscreenchange prop
	// js:"onwebkitfullscreenchange"
	Onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event))

	// Onwebkitfullscreenchange prop
	// js:"setonwebkitfullscreenchange"
	SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(Event))

	// Onwebkitfullscreenerror prop
	// js:"onwebkitfullscreenerror"
	Onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event))

	// Onwebkitfullscreenerror prop
	// js:"setonwebkitfullscreenerror"
	SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(Event))

	// OuterHTML prop
	// js:"outerHTML"
	OuterHTML() (outerHTML string)

	// OuterHTML prop
	// js:"setouterHTML"
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
	// js:"setscrollLeft"
	SetScrollLeft(scrollLeft int)

	// ScrollTop prop
	// js:"scrollTop"
	ScrollTop() (scrollTop int)

	// ScrollTop prop
	// js:"setscrollTop"
	SetScrollTop(scrollTop int)

	// ScrollWidth prop
	// js:"scrollWidth"
	ScrollWidth() (scrollWidth int)

	// TagName prop
	// js:"tagName"
	TagName() (tagName string)
}
