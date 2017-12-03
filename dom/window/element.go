package window

import (
	"github.com/matthewmueller/joy/dom/clientrect"
	"github.com/matthewmueller/joy/dom/clientrectlist"
	"github.com/matthewmueller/joy/dom/domtokenlist"
	"github.com/matthewmueller/joy/dom/mszoomtooptions"
)

// Element interface
// js:"Element"
type Element interface {
	Node

	// QuerySelector
	// js:"querySelector"
	// jsrewrite:"$_.querySelector($1)"
	QuerySelector(selectors string) (e Element)

	// QuerySelectorAll
	// js:"querySelectorAll"
	// jsrewrite:"$_.querySelectorAll($1)"
	QuerySelectorAll(selectors string) (n *NodeList)

	// GetAttribute
	// js:"getAttribute"
	// jsrewrite:"$_.getAttribute($1)"
	GetAttribute(qualifiedName string) (s string)

	// GetAttributeNode
	// js:"getAttributeNode"
	// jsrewrite:"$_.getAttributeNode($1)"
	GetAttributeNode(name string) (a *Attr)

	// GetAttributeNodeNS
	// js:"getAttributeNodeNS"
	// jsrewrite:"$_.getAttributeNodeNS($1, $2)"
	GetAttributeNodeNS(namespaceURI string, localName string) (a *Attr)

	// GetAttributeNS
	// js:"getAttributeNS"
	// jsrewrite:"$_.getAttributeNS($1, $2)"
	GetAttributeNS(namespaceURI string, localName string) (s string)

	// GetBoundingClientRect
	// js:"getBoundingClientRect"
	// jsrewrite:"$_.getBoundingClientRect()"
	GetBoundingClientRect() (c *clientrect.ClientRect)

	// GetClientRects
	// js:"getClientRects"
	// jsrewrite:"$_.getClientRects()"
	GetClientRects() (c *clientrectlist.ClientRectList)

	// GetElementsByTagName
	// js:"getElementsByTagName"
	// jsrewrite:"$_.getElementsByTagName($1)"
	GetElementsByTagName(name string) (n *NodeList)

	// GetElementsByTagNameNS
	// js:"getElementsByTagNameNS"
	// jsrewrite:"$_.getElementsByTagNameNS($1, $2)"
	GetElementsByTagNameNS(namespaceURI string, localName string) (n *NodeList)

	// HasAttribute
	// js:"hasAttribute"
	// jsrewrite:"$_.hasAttribute($1)"
	HasAttribute(name string) (b bool)

	// HasAttributeNS
	// js:"hasAttributeNS"
	// jsrewrite:"$_.hasAttributeNS($1, $2)"
	HasAttributeNS(namespaceURI string, localName string) (b bool)

	// MsGetRegionContent
	// js:"msGetRegionContent"
	// jsrewrite:"$_.msGetRegionContent()"
	MsGetRegionContent() (m *MSRangeCollection)

	// MsGetUntransformedBounds
	// js:"msGetUntransformedBounds"
	// jsrewrite:"$_.msGetUntransformedBounds()"
	MsGetUntransformedBounds() (c *clientrect.ClientRect)

	// MsMatchesSelector
	// js:"msMatchesSelector"
	// jsrewrite:"$_.msMatchesSelector($1)"
	MsMatchesSelector(selectors string) (b bool)

	// MsReleasePointerCapture
	// js:"msReleasePointerCapture"
	// jsrewrite:"$_.msReleasePointerCapture($1)"
	MsReleasePointerCapture(pointerId int)

	// MsSetPointerCapture
	// js:"msSetPointerCapture"
	// jsrewrite:"$_.msSetPointerCapture($1)"
	MsSetPointerCapture(pointerId int)

	// MsZoomTo
	// js:"msZoomTo"
	// jsrewrite:"$_.msZoomTo($1)"
	MsZoomTo(args *mszoomtooptions.MsZoomToOptions)

	// ReleasePointerCapture
	// js:"releasePointerCapture"
	// jsrewrite:"$_.releasePointerCapture($1)"
	ReleasePointerCapture(pointerId int)

	// RemoveAttribute
	// js:"removeAttribute"
	// jsrewrite:"$_.removeAttribute($1)"
	RemoveAttribute(qualifiedName string)

	// RemoveAttributeNode
	// js:"removeAttributeNode"
	// jsrewrite:"$_.removeAttributeNode($1)"
	RemoveAttributeNode(oldAttr *Attr) (a *Attr)

	// RemoveAttributeNS
	// js:"removeAttributeNS"
	// jsrewrite:"$_.removeAttributeNS($1, $2)"
	RemoveAttributeNS(namespaceURI string, localName string)

	// RequestFullscreen
	// js:"requestFullscreen"
	// jsrewrite:"$_.requestFullscreen()"
	RequestFullscreen()

	// RequestPointerLock
	// js:"requestPointerLock"
	// jsrewrite:"$_.requestPointerLock()"
	RequestPointerLock()

	// SetAttribute
	// js:"setAttribute"
	// jsrewrite:"$_.setAttribute($1, $2)"
	SetAttribute(qualifiedName string, value string)

	// SetAttributeNode
	// js:"setAttributeNode"
	// jsrewrite:"$_.setAttributeNode($1)"
	SetAttributeNode(newAttr *Attr) (a *Attr)

	// SetAttributeNodeNS
	// js:"setAttributeNodeNS"
	// jsrewrite:"$_.setAttributeNodeNS($1)"
	SetAttributeNodeNS(newAttr *Attr) (a *Attr)

	// SetAttributeNS
	// js:"setAttributeNS"
	// jsrewrite:"$_.setAttributeNS($1, $2, $3)"
	SetAttributeNS(namespaceURI string, qualifiedName string, value string)

	// SetPointerCapture
	// js:"setPointerCapture"
	// jsrewrite:"$_.setPointerCapture($1)"
	SetPointerCapture(pointerId int)

	// WebkitMatchesSelector
	// js:"webkitMatchesSelector"
	// jsrewrite:"$_.webkitMatchesSelector($1)"
	WebkitMatchesSelector(selectors string) (b bool)

	// WebkitRequestFullscreen
	// js:"webkitRequestFullscreen"
	// jsrewrite:"$_.webkitRequestFullscreen()"
	WebkitRequestFullscreen()

	// WebkitRequestFullScreen
	// js:"webkitRequestFullScreen"
	// jsrewrite:"$_.webkitRequestFullScreen()"
	WebkitRequestFullScreen()

	// Onpointercancel prop
	// js:"onpointercancel"
	// jsrewrite:"$_.onpointercancel"
	Onpointercancel() (onpointercancel func(Event))

	// SetOnpointercancel prop
	// js:"onpointercancel"
	// jsrewrite:"$_.onpointercancel = $1"
	SetOnpointercancel(onpointercancel func(Event))

	// Onpointerdown prop
	// js:"onpointerdown"
	// jsrewrite:"$_.onpointerdown"
	Onpointerdown() (onpointerdown func(Event))

	// SetOnpointerdown prop
	// js:"onpointerdown"
	// jsrewrite:"$_.onpointerdown = $1"
	SetOnpointerdown(onpointerdown func(Event))

	// Onpointerenter prop
	// js:"onpointerenter"
	// jsrewrite:"$_.onpointerenter"
	Onpointerenter() (onpointerenter func(Event))

	// SetOnpointerenter prop
	// js:"onpointerenter"
	// jsrewrite:"$_.onpointerenter = $1"
	SetOnpointerenter(onpointerenter func(Event))

	// Onpointerleave prop
	// js:"onpointerleave"
	// jsrewrite:"$_.onpointerleave"
	Onpointerleave() (onpointerleave func(Event))

	// SetOnpointerleave prop
	// js:"onpointerleave"
	// jsrewrite:"$_.onpointerleave = $1"
	SetOnpointerleave(onpointerleave func(Event))

	// Onpointermove prop
	// js:"onpointermove"
	// jsrewrite:"$_.onpointermove"
	Onpointermove() (onpointermove func(Event))

	// SetOnpointermove prop
	// js:"onpointermove"
	// jsrewrite:"$_.onpointermove = $1"
	SetOnpointermove(onpointermove func(Event))

	// Onpointerout prop
	// js:"onpointerout"
	// jsrewrite:"$_.onpointerout"
	Onpointerout() (onpointerout func(Event))

	// SetOnpointerout prop
	// js:"onpointerout"
	// jsrewrite:"$_.onpointerout = $1"
	SetOnpointerout(onpointerout func(Event))

	// Onpointerover prop
	// js:"onpointerover"
	// jsrewrite:"$_.onpointerover"
	Onpointerover() (onpointerover func(Event))

	// SetOnpointerover prop
	// js:"onpointerover"
	// jsrewrite:"$_.onpointerover = $1"
	SetOnpointerover(onpointerover func(Event))

	// Onpointerup prop
	// js:"onpointerup"
	// jsrewrite:"$_.onpointerup"
	Onpointerup() (onpointerup func(Event))

	// SetOnpointerup prop
	// js:"onpointerup"
	// jsrewrite:"$_.onpointerup = $1"
	SetOnpointerup(onpointerup func(Event))

	// Onwheel prop
	// js:"onwheel"
	// jsrewrite:"$_.onwheel"
	Onwheel() (onwheel func(Event))

	// SetOnwheel prop
	// js:"onwheel"
	// jsrewrite:"$_.onwheel = $1"
	SetOnwheel(onwheel func(Event))

	// ChildElementCount prop
	// js:"childElementCount"
	// jsrewrite:"$_.childElementCount"
	ChildElementCount() (childElementCount uint)

	// FirstElementChild prop
	// js:"firstElementChild"
	// jsrewrite:"$_.firstElementChild"
	FirstElementChild() (firstElementChild Element)

	// LastElementChild prop
	// js:"lastElementChild"
	// jsrewrite:"$_.lastElementChild"
	LastElementChild() (lastElementChild Element)

	// NextElementSibling prop
	// js:"nextElementSibling"
	// jsrewrite:"$_.nextElementSibling"
	NextElementSibling() (nextElementSibling Element)

	// PreviousElementSibling prop
	// js:"previousElementSibling"
	// jsrewrite:"$_.previousElementSibling"
	PreviousElementSibling() (previousElementSibling Element)

	// ClassList prop
	// js:"classList"
	// jsrewrite:"$_.classList"
	ClassList() (classList domtokenlist.DOMTokenList)

	// ClassName prop
	// js:"className"
	// jsrewrite:"$_.className"
	ClassName() (className string)

	// SetClassName prop
	// js:"className"
	// jsrewrite:"$_.className = $1"
	SetClassName(className string)

	// ClientHeight prop
	// js:"clientHeight"
	// jsrewrite:"$_.clientHeight"
	ClientHeight() (clientHeight int)

	// ClientLeft prop
	// js:"clientLeft"
	// jsrewrite:"$_.clientLeft"
	ClientLeft() (clientLeft int)

	// ClientTop prop
	// js:"clientTop"
	// jsrewrite:"$_.clientTop"
	ClientTop() (clientTop int)

	// ClientWidth prop
	// js:"clientWidth"
	// jsrewrite:"$_.clientWidth"
	ClientWidth() (clientWidth int)

	// ID prop
	// js:"id"
	// jsrewrite:"$_.id"
	ID() (id string)

	// SetID prop
	// js:"id"
	// jsrewrite:"$_.id = $1"
	SetID(id string)

	// InnerHTML prop
	// js:"innerHTML"
	// jsrewrite:"$_.innerHTML"
	InnerHTML() (innerHTML string)

	// SetInnerHTML prop
	// js:"innerHTML"
	// jsrewrite:"$_.innerHTML = $1"
	SetInnerHTML(innerHTML string)

	// MsContentZoomFactor prop
	// js:"msContentZoomFactor"
	// jsrewrite:"$_.msContentZoomFactor"
	MsContentZoomFactor() (msContentZoomFactor float32)

	// SetMsContentZoomFactor prop
	// js:"msContentZoomFactor"
	// jsrewrite:"$_.msContentZoomFactor = $1"
	SetMsContentZoomFactor(msContentZoomFactor float32)

	// MsRegionOverflow prop
	// js:"msRegionOverflow"
	// jsrewrite:"$_.msRegionOverflow"
	MsRegionOverflow() (msRegionOverflow string)

	// Onariarequest prop
	// js:"onariarequest"
	// jsrewrite:"$_.onariarequest"
	Onariarequest() (onariarequest func(Event))

	// SetOnariarequest prop
	// js:"onariarequest"
	// jsrewrite:"$_.onariarequest = $1"
	SetOnariarequest(onariarequest func(Event))

	// Oncommand prop
	// js:"oncommand"
	// jsrewrite:"$_.oncommand"
	Oncommand() (oncommand func(Event))

	// SetOncommand prop
	// js:"oncommand"
	// jsrewrite:"$_.oncommand = $1"
	SetOncommand(oncommand func(Event))

	// Ongotpointercapture prop
	// js:"ongotpointercapture"
	// jsrewrite:"$_.ongotpointercapture"
	Ongotpointercapture() (ongotpointercapture func(*PointerEvent))

	// SetOngotpointercapture prop
	// js:"ongotpointercapture"
	// jsrewrite:"$_.ongotpointercapture = $1"
	SetOngotpointercapture(ongotpointercapture func(*PointerEvent))

	// Onlostpointercapture prop
	// js:"onlostpointercapture"
	// jsrewrite:"$_.onlostpointercapture"
	Onlostpointercapture() (onlostpointercapture func(*PointerEvent))

	// SetOnlostpointercapture prop
	// js:"onlostpointercapture"
	// jsrewrite:"$_.onlostpointercapture = $1"
	SetOnlostpointercapture(onlostpointercapture func(*PointerEvent))

	// Onmsgesturechange prop
	// js:"onmsgesturechange"
	// jsrewrite:"$_.onmsgesturechange"
	Onmsgesturechange() (onmsgesturechange func(*MSGestureEvent))

	// SetOnmsgesturechange prop
	// js:"onmsgesturechange"
	// jsrewrite:"$_.onmsgesturechange = $1"
	SetOnmsgesturechange(onmsgesturechange func(*MSGestureEvent))

	// Onmsgesturedoubletap prop
	// js:"onmsgesturedoubletap"
	// jsrewrite:"$_.onmsgesturedoubletap"
	Onmsgesturedoubletap() (onmsgesturedoubletap func(*MSGestureEvent))

	// SetOnmsgesturedoubletap prop
	// js:"onmsgesturedoubletap"
	// jsrewrite:"$_.onmsgesturedoubletap = $1"
	SetOnmsgesturedoubletap(onmsgesturedoubletap func(*MSGestureEvent))

	// Onmsgestureend prop
	// js:"onmsgestureend"
	// jsrewrite:"$_.onmsgestureend"
	Onmsgestureend() (onmsgestureend func(*MSGestureEvent))

	// SetOnmsgestureend prop
	// js:"onmsgestureend"
	// jsrewrite:"$_.onmsgestureend = $1"
	SetOnmsgestureend(onmsgestureend func(*MSGestureEvent))

	// Onmsgesturehold prop
	// js:"onmsgesturehold"
	// jsrewrite:"$_.onmsgesturehold"
	Onmsgesturehold() (onmsgesturehold func(*MSGestureEvent))

	// SetOnmsgesturehold prop
	// js:"onmsgesturehold"
	// jsrewrite:"$_.onmsgesturehold = $1"
	SetOnmsgesturehold(onmsgesturehold func(*MSGestureEvent))

	// Onmsgesturestart prop
	// js:"onmsgesturestart"
	// jsrewrite:"$_.onmsgesturestart"
	Onmsgesturestart() (onmsgesturestart func(*MSGestureEvent))

	// SetOnmsgesturestart prop
	// js:"onmsgesturestart"
	// jsrewrite:"$_.onmsgesturestart = $1"
	SetOnmsgesturestart(onmsgesturestart func(*MSGestureEvent))

	// Onmsgesturetap prop
	// js:"onmsgesturetap"
	// jsrewrite:"$_.onmsgesturetap"
	Onmsgesturetap() (onmsgesturetap func(*MSGestureEvent))

	// SetOnmsgesturetap prop
	// js:"onmsgesturetap"
	// jsrewrite:"$_.onmsgesturetap = $1"
	SetOnmsgesturetap(onmsgesturetap func(*MSGestureEvent))

	// Onmsgotpointercapture prop
	// js:"onmsgotpointercapture"
	// jsrewrite:"$_.onmsgotpointercapture"
	Onmsgotpointercapture() (onmsgotpointercapture func(*MSPointerEvent))

	// SetOnmsgotpointercapture prop
	// js:"onmsgotpointercapture"
	// jsrewrite:"$_.onmsgotpointercapture = $1"
	SetOnmsgotpointercapture(onmsgotpointercapture func(*MSPointerEvent))

	// Onmsinertiastart prop
	// js:"onmsinertiastart"
	// jsrewrite:"$_.onmsinertiastart"
	Onmsinertiastart() (onmsinertiastart func(*MSGestureEvent))

	// SetOnmsinertiastart prop
	// js:"onmsinertiastart"
	// jsrewrite:"$_.onmsinertiastart = $1"
	SetOnmsinertiastart(onmsinertiastart func(*MSGestureEvent))

	// Onmslostpointercapture prop
	// js:"onmslostpointercapture"
	// jsrewrite:"$_.onmslostpointercapture"
	Onmslostpointercapture() (onmslostpointercapture func(*MSPointerEvent))

	// SetOnmslostpointercapture prop
	// js:"onmslostpointercapture"
	// jsrewrite:"$_.onmslostpointercapture = $1"
	SetOnmslostpointercapture(onmslostpointercapture func(*MSPointerEvent))

	// Onmspointercancel prop
	// js:"onmspointercancel"
	// jsrewrite:"$_.onmspointercancel"
	Onmspointercancel() (onmspointercancel func(*MSPointerEvent))

	// SetOnmspointercancel prop
	// js:"onmspointercancel"
	// jsrewrite:"$_.onmspointercancel = $1"
	SetOnmspointercancel(onmspointercancel func(*MSPointerEvent))

	// Onmspointerdown prop
	// js:"onmspointerdown"
	// jsrewrite:"$_.onmspointerdown"
	Onmspointerdown() (onmspointerdown func(*MSPointerEvent))

	// SetOnmspointerdown prop
	// js:"onmspointerdown"
	// jsrewrite:"$_.onmspointerdown = $1"
	SetOnmspointerdown(onmspointerdown func(*MSPointerEvent))

	// Onmspointerenter prop
	// js:"onmspointerenter"
	// jsrewrite:"$_.onmspointerenter"
	Onmspointerenter() (onmspointerenter func(*MSPointerEvent))

	// SetOnmspointerenter prop
	// js:"onmspointerenter"
	// jsrewrite:"$_.onmspointerenter = $1"
	SetOnmspointerenter(onmspointerenter func(*MSPointerEvent))

	// Onmspointerleave prop
	// js:"onmspointerleave"
	// jsrewrite:"$_.onmspointerleave"
	Onmspointerleave() (onmspointerleave func(*MSPointerEvent))

	// SetOnmspointerleave prop
	// js:"onmspointerleave"
	// jsrewrite:"$_.onmspointerleave = $1"
	SetOnmspointerleave(onmspointerleave func(*MSPointerEvent))

	// Onmspointermove prop
	// js:"onmspointermove"
	// jsrewrite:"$_.onmspointermove"
	Onmspointermove() (onmspointermove func(*MSPointerEvent))

	// SetOnmspointermove prop
	// js:"onmspointermove"
	// jsrewrite:"$_.onmspointermove = $1"
	SetOnmspointermove(onmspointermove func(*MSPointerEvent))

	// Onmspointerout prop
	// js:"onmspointerout"
	// jsrewrite:"$_.onmspointerout"
	Onmspointerout() (onmspointerout func(*MSPointerEvent))

	// SetOnmspointerout prop
	// js:"onmspointerout"
	// jsrewrite:"$_.onmspointerout = $1"
	SetOnmspointerout(onmspointerout func(*MSPointerEvent))

	// Onmspointerover prop
	// js:"onmspointerover"
	// jsrewrite:"$_.onmspointerover"
	Onmspointerover() (onmspointerover func(*MSPointerEvent))

	// SetOnmspointerover prop
	// js:"onmspointerover"
	// jsrewrite:"$_.onmspointerover = $1"
	SetOnmspointerover(onmspointerover func(*MSPointerEvent))

	// Onmspointerup prop
	// js:"onmspointerup"
	// jsrewrite:"$_.onmspointerup"
	Onmspointerup() (onmspointerup func(*MSPointerEvent))

	// SetOnmspointerup prop
	// js:"onmspointerup"
	// jsrewrite:"$_.onmspointerup = $1"
	SetOnmspointerup(onmspointerup func(*MSPointerEvent))

	// Ontouchcancel prop
	// js:"ontouchcancel"
	// jsrewrite:"$_.ontouchcancel"
	Ontouchcancel() (ontouchcancel func(*TouchEvent))

	// SetOntouchcancel prop
	// js:"ontouchcancel"
	// jsrewrite:"$_.ontouchcancel = $1"
	SetOntouchcancel(ontouchcancel func(*TouchEvent))

	// Ontouchend prop
	// js:"ontouchend"
	// jsrewrite:"$_.ontouchend"
	Ontouchend() (ontouchend func(*TouchEvent))

	// SetOntouchend prop
	// js:"ontouchend"
	// jsrewrite:"$_.ontouchend = $1"
	SetOntouchend(ontouchend func(*TouchEvent))

	// Ontouchmove prop
	// js:"ontouchmove"
	// jsrewrite:"$_.ontouchmove"
	Ontouchmove() (ontouchmove func(*TouchEvent))

	// SetOntouchmove prop
	// js:"ontouchmove"
	// jsrewrite:"$_.ontouchmove = $1"
	SetOntouchmove(ontouchmove func(*TouchEvent))

	// Ontouchstart prop
	// js:"ontouchstart"
	// jsrewrite:"$_.ontouchstart"
	Ontouchstart() (ontouchstart func(*TouchEvent))

	// SetOntouchstart prop
	// js:"ontouchstart"
	// jsrewrite:"$_.ontouchstart = $1"
	SetOntouchstart(ontouchstart func(*TouchEvent))

	// Onwebkitfullscreenchange prop
	// js:"onwebkitfullscreenchange"
	// jsrewrite:"$_.onwebkitfullscreenchange"
	Onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event))

	// SetOnwebkitfullscreenchange prop
	// js:"onwebkitfullscreenchange"
	// jsrewrite:"$_.onwebkitfullscreenchange = $1"
	SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(Event))

	// Onwebkitfullscreenerror prop
	// js:"onwebkitfullscreenerror"
	// jsrewrite:"$_.onwebkitfullscreenerror"
	Onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event))

	// SetOnwebkitfullscreenerror prop
	// js:"onwebkitfullscreenerror"
	// jsrewrite:"$_.onwebkitfullscreenerror = $1"
	SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(Event))

	// OuterHTML prop
	// js:"outerHTML"
	// jsrewrite:"$_.outerHTML"
	OuterHTML() (outerHTML string)

	// SetOuterHTML prop
	// js:"outerHTML"
	// jsrewrite:"$_.outerHTML = $1"
	SetOuterHTML(outerHTML string)

	// Prefix prop
	// js:"prefix"
	// jsrewrite:"$_.prefix"
	Prefix() (prefix string)

	// ScrollHeight prop
	// js:"scrollHeight"
	// jsrewrite:"$_.scrollHeight"
	ScrollHeight() (scrollHeight int)

	// ScrollLeft prop
	// js:"scrollLeft"
	// jsrewrite:"$_.scrollLeft"
	ScrollLeft() (scrollLeft int)

	// SetScrollLeft prop
	// js:"scrollLeft"
	// jsrewrite:"$_.scrollLeft = $1"
	SetScrollLeft(scrollLeft int)

	// ScrollTop prop
	// js:"scrollTop"
	// jsrewrite:"$_.scrollTop"
	ScrollTop() (scrollTop int)

	// SetScrollTop prop
	// js:"scrollTop"
	// jsrewrite:"$_.scrollTop = $1"
	SetScrollTop(scrollTop int)

	// ScrollWidth prop
	// js:"scrollWidth"
	// jsrewrite:"$_.scrollWidth"
	ScrollWidth() (scrollWidth int)

	// TagName prop
	// js:"tagName"
	// jsrewrite:"$_.tagName"
	TagName() (tagName string)
}
