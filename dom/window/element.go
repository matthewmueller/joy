package window

import (
	"github.com/matthewmueller/golly/dom/clientrect"
	"github.com/matthewmueller/golly/dom/clientrectlist"
	"github.com/matthewmueller/golly/dom/domtokenlist"
	"github.com/matthewmueller/golly/dom/mszoomtooptions"
	"github.com/matthewmueller/golly/js"
)

// js:"Element,omit"
type Element interface {
	Node

	// QuerySelector
	// js:"querySelector,rewrite=queryselector"
	QuerySelector(selectors string) (e Element)

	// QuerySelectorAll
	// js:"querySelectorAll,rewrite=queryselectorall"
	QuerySelectorAll(selectors string) (n *NodeList)

	// // Remove
	// // js:"remove,rewrite=remove"
	// Remove()

	// // GetAttribute
	// // js:"getAttribute,rewrite=getattribute"
	// GetAttribute(qualifiedName string) (s string)

	// // GetAttributeNode
	// // js:"getAttributeNode,rewrite=getattributenode"
	// GetAttributeNode(name string) (a *Attr)

	// // GetAttributeNodeNS
	// // js:"getAttributeNodeNS,rewrite=getattributenodens"
	// GetAttributeNodeNS(namespaceURI string, localName string) (a *Attr)

	// // GetAttributeNS
	// // js:"getAttributeNS,rewrite=getattributens"
	// GetAttributeNS(namespaceURI string, localName string) (s string)

	// // GetBoundingClientRect
	// // js:"getBoundingClientRect,rewrite=getboundingclientrect"
	// GetBoundingClientRect() (c *clientrect.ClientRect)

	// // GetClientRects
	// // js:"getClientRects,rewrite=getclientrects"
	// GetClientRects() (c *clientrectlist.ClientRectList)

	// // GetElementsByTagName
	// // js:"getElementsByTagName,rewrite=getelementsbytagname"
	// GetElementsByTagName(name string) (n *NodeList)

	// // GetElementsByTagNameNS
	// // js:"getElementsByTagNameNS,rewrite=getelementsbytagnamens"
	// GetElementsByTagNameNS(namespaceURI string, localName string) (n *NodeList)

	// // HasAttribute
	// // js:"hasAttribute,rewrite=hasattribute"
	// HasAttribute(name string) (b bool)

	// // HasAttributeNS
	// // js:"hasAttributeNS,rewrite=hasattributens"
	// HasAttributeNS(namespaceURI string, localName string) (b bool)

	// // MsGetRegionContent
	// // js:"msGetRegionContent,rewrite=msgetregioncontent"
	// MsGetRegionContent() (m *MSRangeCollection)

	// // MsGetUntransformedBounds
	// // js:"msGetUntransformedBounds,rewrite=msgetuntransformedbounds"
	// MsGetUntransformedBounds() (c *clientrect.ClientRect)

	// // MsMatchesSelector
	// // js:"msMatchesSelector,rewrite=msmatchesselector"
	// MsMatchesSelector(selectors string) (b bool)

	// // MsReleasePointerCapture
	// // js:"msReleasePointerCapture,rewrite=msreleasepointercapture"
	// MsReleasePointerCapture(pointerId int)

	// // MsSetPointerCapture
	// // js:"msSetPointerCapture,rewrite=mssetpointercapture"
	// MsSetPointerCapture(pointerId int)

	// // MsZoomTo
	// // js:"msZoomTo,rewrite=mszoomto"
	// MsZoomTo(args *mszoomtooptions.MsZoomToOptions)

	// // ReleasePointerCapture
	// // js:"releasePointerCapture,rewrite=releasepointercapture"
	// ReleasePointerCapture(pointerId int)

	// // RemoveAttribute
	// // js:"removeAttribute,rewrite=removeattribute"
	// RemoveAttribute(qualifiedName string)

	// // RemoveAttributeNode
	// // js:"removeAttributeNode,rewrite=removeattributenode"
	// RemoveAttributeNode(oldAttr *Attr) (a *Attr)

	// // RemoveAttributeNS
	// // js:"removeAttributeNS,rewrite=removeattributens"
	// RemoveAttributeNS(namespaceURI string, localName string)

	// // RequestFullscreen
	// // js:"requestFullscreen,rewrite=requestfullscreen"
	// RequestFullscreen()

	// // RequestPointerLock
	// // js:"requestPointerLock,rewrite=requestpointerlock"
	// RequestPointerLock()

	// // SetAttribute
	// // js:"setAttribute,rewrite=setattribute"
	// SetAttribute(qualifiedName string, value string)

	// // SetAttributeNode
	// // js:"setAttributeNode,rewrite=setattributenode"
	// SetAttributeNode(newAttr *Attr) (a *Attr)

	// // SetAttributeNodeNS
	// // js:"setAttributeNodeNS,rewrite=setattributenodens"
	// SetAttributeNodeNS(newAttr *Attr) (a *Attr)

	// // SetAttributeNS
	// // js:"setAttributeNS,rewrite=setattributens"
	// SetAttributeNS(namespaceURI string, qualifiedName string, value string)

	// // SetPointerCapture
	// // js:"setPointerCapture,rewrite=setpointercapture"
	// SetPointerCapture(pointerId int)

	// // WebkitMatchesSelector
	// // js:"webkitMatchesSelector,rewrite=webkitmatchesselector"
	// WebkitMatchesSelector(selectors string) (b bool)

	// // WebkitRequestFullscreen
	// // js:"webkitRequestFullscreen,rewrite=webkitrequestfullscreen"
	// WebkitRequestFullscreen()

	// // WebkitRequestFullScreen
	// // js:"webkitRequestFullScreen,rewrite=webkitrequestfullscreen"
	// WebkitRequestFullScreen()

	// // Onpointercancel prop
	// // js:"onpointercancel,rewrite=onpointercancel"
	// Onpointercancel() (onpointercancel func(Event))

	// // Onpointercancel prop
	// // js:"setonpointercancel,rewrite=setonpointercancel"
	// SetOnpointercancel(onpointercancel func(Event))

	// // Onpointerdown prop
	// // js:"onpointerdown,rewrite=onpointerdown"
	// Onpointerdown() (onpointerdown func(Event))

	// // Onpointerdown prop
	// // js:"setonpointerdown,rewrite=setonpointerdown"
	// SetOnpointerdown(onpointerdown func(Event))

	// // Onpointerenter prop
	// // js:"onpointerenter,rewrite=onpointerenter"
	// Onpointerenter() (onpointerenter func(Event))

	// // Onpointerenter prop
	// // js:"setonpointerenter,rewrite=setonpointerenter"
	// SetOnpointerenter(onpointerenter func(Event))

	// // Onpointerleave prop
	// // js:"onpointerleave,rewrite=onpointerleave"
	// Onpointerleave() (onpointerleave func(Event))

	// // Onpointerleave prop
	// // js:"setonpointerleave,rewrite=setonpointerleave"
	// SetOnpointerleave(onpointerleave func(Event))

	// // Onpointermove prop
	// // js:"onpointermove,rewrite=onpointermove"
	// Onpointermove() (onpointermove func(Event))

	// // Onpointermove prop
	// // js:"setonpointermove,rewrite=setonpointermove"
	// SetOnpointermove(onpointermove func(Event))

	// // Onpointerout prop
	// // js:"onpointerout,rewrite=onpointerout"
	// Onpointerout() (onpointerout func(Event))

	// // Onpointerout prop
	// // js:"setonpointerout,rewrite=setonpointerout"
	// SetOnpointerout(onpointerout func(Event))

	// // Onpointerover prop
	// // js:"onpointerover,rewrite=onpointerover"
	// Onpointerover() (onpointerover func(Event))

	// // Onpointerover prop
	// // js:"setonpointerover,rewrite=setonpointerover"
	// SetOnpointerover(onpointerover func(Event))

	// // Onpointerup prop
	// // js:"onpointerup,rewrite=onpointerup"
	// Onpointerup() (onpointerup func(Event))

	// // Onpointerup prop
	// // js:"setonpointerup,rewrite=setonpointerup"
	// SetOnpointerup(onpointerup func(Event))

	// // Onwheel prop
	// // js:"onwheel,rewrite=onwheel"
	// Onwheel() (onwheel func(Event))

	// // Onwheel prop
	// // js:"setonwheel,rewrite=setonwheel"
	// SetOnwheel(onwheel func(Event))

	// // ChildElementCount prop
	// // js:"childElementCount,rewrite=childelementcount"
	// ChildElementCount() (childElementCount uint)

	// // FirstElementChild prop
	// // js:"firstElementChild,rewrite=firstelementchild"
	// FirstElementChild() (firstElementChild Element)

	// // LastElementChild prop
	// // js:"lastElementChild,rewrite=lastelementchild"
	// LastElementChild() (lastElementChild Element)

	// // NextElementSibling prop
	// // js:"nextElementSibling,rewrite=nextelementsibling"
	// NextElementSibling() (nextElementSibling Element)

	// // PreviousElementSibling prop
	// // js:"previousElementSibling,rewrite=previouselementsibling"
	// PreviousElementSibling() (previousElementSibling Element)

	// // ClassList prop
	// // js:"classList,rewrite=classlist"
	// ClassList() (classList domtokenlist.DOMTokenList)

	// // ClassName prop
	// // js:"className,rewrite=classname"
	// ClassName() (className string)

	// // ClassName prop
	// // js:"setclassName,rewrite=setclassname"
	// SetClassName(className string)

	// // ClientHeight prop
	// // js:"clientHeight,rewrite=clientheight"
	// ClientHeight() (clientHeight int)

	// // ClientLeft prop
	// // js:"clientLeft,rewrite=clientleft"
	// ClientLeft() (clientLeft int)

	// // ClientTop prop
	// // js:"clientTop,rewrite=clienttop"
	// ClientTop() (clientTop int)

	// // ClientWidth prop
	// // js:"clientWidth,rewrite=clientwidth"
	// ClientWidth() (clientWidth int)

	// // ID prop
	// // js:"id,rewrite=id"
	// ID() (id string)

	// // ID prop
	// // js:"setid,rewrite=setid"
	// SetID(id string)

	// // InnerHTML prop
	// // js:"innerHTML,rewrite=innerhtml"
	// InnerHTML() (innerHTML string)

	// // InnerHTML prop
	// // js:"setinnerHTML,rewrite=setinnerhtml"
	// SetInnerHTML(innerHTML string)

	// // MsContentZoomFactor prop
	// // js:"msContentZoomFactor,rewrite=mscontentzoomfactor"
	// MsContentZoomFactor() (msContentZoomFactor float32)

	// // MsContentZoomFactor prop
	// // js:"setmsContentZoomFactor,rewrite=setmscontentzoomfactor"
	// SetMsContentZoomFactor(msContentZoomFactor float32)

	// // MsRegionOverflow prop
	// // js:"msRegionOverflow,rewrite=msregionoverflow"
	// MsRegionOverflow() (msRegionOverflow string)

	// // Onariarequest prop
	// // js:"onariarequest,rewrite=onariarequest"
	// Onariarequest() (onariarequest func(Event))

	// // Onariarequest prop
	// // js:"setonariarequest,rewrite=setonariarequest"
	// SetOnariarequest(onariarequest func(Event))

	// // Oncommand prop
	// // js:"oncommand,rewrite=oncommand"
	// Oncommand() (oncommand func(Event))

	// // Oncommand prop
	// // js:"setoncommand,rewrite=setoncommand"
	// SetOncommand(oncommand func(Event))

	// // Ongotpointercapture prop
	// // js:"ongotpointercapture,rewrite=ongotpointercapture"
	// Ongotpointercapture() (ongotpointercapture func(*PointerEvent))

	// // Ongotpointercapture prop
	// // js:"setongotpointercapture,rewrite=setongotpointercapture"
	// SetOngotpointercapture(ongotpointercapture func(*PointerEvent))

	// // Onlostpointercapture prop
	// // js:"onlostpointercapture,rewrite=onlostpointercapture"
	// Onlostpointercapture() (onlostpointercapture func(*PointerEvent))

	// // Onlostpointercapture prop
	// // js:"setonlostpointercapture,rewrite=setonlostpointercapture"
	// SetOnlostpointercapture(onlostpointercapture func(*PointerEvent))

	// // Onmsgesturechange prop
	// // js:"onmsgesturechange,rewrite=onmsgesturechange"
	// Onmsgesturechange() (onmsgesturechange func(*MSGestureEvent))

	// // Onmsgesturechange prop
	// // js:"setonmsgesturechange,rewrite=setonmsgesturechange"
	// SetOnmsgesturechange(onmsgesturechange func(*MSGestureEvent))

	// // Onmsgesturedoubletap prop
	// // js:"onmsgesturedoubletap,rewrite=onmsgesturedoubletap"
	// Onmsgesturedoubletap() (onmsgesturedoubletap func(*MSGestureEvent))

	// // Onmsgesturedoubletap prop
	// // js:"setonmsgesturedoubletap,rewrite=setonmsgesturedoubletap"
	// SetOnmsgesturedoubletap(onmsgesturedoubletap func(*MSGestureEvent))

	// // Onmsgestureend prop
	// // js:"onmsgestureend,rewrite=onmsgestureend"
	// Onmsgestureend() (onmsgestureend func(*MSGestureEvent))

	// // Onmsgestureend prop
	// // js:"setonmsgestureend,rewrite=setonmsgestureend"
	// SetOnmsgestureend(onmsgestureend func(*MSGestureEvent))

	// // Onmsgesturehold prop
	// // js:"onmsgesturehold,rewrite=onmsgesturehold"
	// Onmsgesturehold() (onmsgesturehold func(*MSGestureEvent))

	// // Onmsgesturehold prop
	// // js:"setonmsgesturehold,rewrite=setonmsgesturehold"
	// SetOnmsgesturehold(onmsgesturehold func(*MSGestureEvent))

	// // Onmsgesturestart prop
	// // js:"onmsgesturestart,rewrite=onmsgesturestart"
	// Onmsgesturestart() (onmsgesturestart func(*MSGestureEvent))

	// // Onmsgesturestart prop
	// // js:"setonmsgesturestart,rewrite=setonmsgesturestart"
	// SetOnmsgesturestart(onmsgesturestart func(*MSGestureEvent))

	// // Onmsgesturetap prop
	// // js:"onmsgesturetap,rewrite=onmsgesturetap"
	// Onmsgesturetap() (onmsgesturetap func(*MSGestureEvent))

	// // Onmsgesturetap prop
	// // js:"setonmsgesturetap,rewrite=setonmsgesturetap"
	// SetOnmsgesturetap(onmsgesturetap func(*MSGestureEvent))

	// // Onmsgotpointercapture prop
	// // js:"onmsgotpointercapture,rewrite=onmsgotpointercapture"
	// Onmsgotpointercapture() (onmsgotpointercapture func(*MSPointerEvent))

	// // Onmsgotpointercapture prop
	// // js:"setonmsgotpointercapture,rewrite=setonmsgotpointercapture"
	// SetOnmsgotpointercapture(onmsgotpointercapture func(*MSPointerEvent))

	// // Onmsinertiastart prop
	// // js:"onmsinertiastart,rewrite=onmsinertiastart"
	// Onmsinertiastart() (onmsinertiastart func(*MSGestureEvent))

	// // Onmsinertiastart prop
	// // js:"setonmsinertiastart,rewrite=setonmsinertiastart"
	// SetOnmsinertiastart(onmsinertiastart func(*MSGestureEvent))

	// // Onmslostpointercapture prop
	// // js:"onmslostpointercapture,rewrite=onmslostpointercapture"
	// Onmslostpointercapture() (onmslostpointercapture func(*MSPointerEvent))

	// // Onmslostpointercapture prop
	// // js:"setonmslostpointercapture,rewrite=setonmslostpointercapture"
	// SetOnmslostpointercapture(onmslostpointercapture func(*MSPointerEvent))

	// // Onmspointercancel prop
	// // js:"onmspointercancel,rewrite=onmspointercancel"
	// Onmspointercancel() (onmspointercancel func(*MSPointerEvent))

	// // Onmspointercancel prop
	// // js:"setonmspointercancel,rewrite=setonmspointercancel"
	// SetOnmspointercancel(onmspointercancel func(*MSPointerEvent))

	// // Onmspointerdown prop
	// // js:"onmspointerdown,rewrite=onmspointerdown"
	// Onmspointerdown() (onmspointerdown func(*MSPointerEvent))

	// // Onmspointerdown prop
	// // js:"setonmspointerdown,rewrite=setonmspointerdown"
	// SetOnmspointerdown(onmspointerdown func(*MSPointerEvent))

	// // Onmspointerenter prop
	// // js:"onmspointerenter,rewrite=onmspointerenter"
	// Onmspointerenter() (onmspointerenter func(*MSPointerEvent))

	// // Onmspointerenter prop
	// // js:"setonmspointerenter,rewrite=setonmspointerenter"
	// SetOnmspointerenter(onmspointerenter func(*MSPointerEvent))

	// // Onmspointerleave prop
	// // js:"onmspointerleave,rewrite=onmspointerleave"
	// Onmspointerleave() (onmspointerleave func(*MSPointerEvent))

	// // Onmspointerleave prop
	// // js:"setonmspointerleave,rewrite=setonmspointerleave"
	// SetOnmspointerleave(onmspointerleave func(*MSPointerEvent))

	// // Onmspointermove prop
	// // js:"onmspointermove,rewrite=onmspointermove"
	// Onmspointermove() (onmspointermove func(*MSPointerEvent))

	// // Onmspointermove prop
	// // js:"setonmspointermove,rewrite=setonmspointermove"
	// SetOnmspointermove(onmspointermove func(*MSPointerEvent))

	// // Onmspointerout prop
	// // js:"onmspointerout,rewrite=onmspointerout"
	// Onmspointerout() (onmspointerout func(*MSPointerEvent))

	// // Onmspointerout prop
	// // js:"setonmspointerout,rewrite=setonmspointerout"
	// SetOnmspointerout(onmspointerout func(*MSPointerEvent))

	// // Onmspointerover prop
	// // js:"onmspointerover,rewrite=onmspointerover"
	// Onmspointerover() (onmspointerover func(*MSPointerEvent))

	// // Onmspointerover prop
	// // js:"setonmspointerover,rewrite=setonmspointerover"
	// SetOnmspointerover(onmspointerover func(*MSPointerEvent))

	// // Onmspointerup prop
	// // js:"onmspointerup,rewrite=onmspointerup"
	// Onmspointerup() (onmspointerup func(*MSPointerEvent))

	// // Onmspointerup prop
	// // js:"setonmspointerup,rewrite=setonmspointerup"
	// SetOnmspointerup(onmspointerup func(*MSPointerEvent))

	// // Ontouchcancel prop
	// // js:"ontouchcancel,rewrite=ontouchcancel"
	// Ontouchcancel() (ontouchcancel func(*TouchEvent))

	// // Ontouchcancel prop
	// // js:"setontouchcancel,rewrite=setontouchcancel"
	// SetOntouchcancel(ontouchcancel func(*TouchEvent))

	// // Ontouchend prop
	// // js:"ontouchend,rewrite=ontouchend"
	// Ontouchend() (ontouchend func(*TouchEvent))

	// // Ontouchend prop
	// // js:"setontouchend,rewrite=setontouchend"
	// SetOntouchend(ontouchend func(*TouchEvent))

	// // Ontouchmove prop
	// // js:"ontouchmove,rewrite=ontouchmove"
	// Ontouchmove() (ontouchmove func(*TouchEvent))

	// // Ontouchmove prop
	// // js:"setontouchmove,rewrite=setontouchmove"
	// SetOntouchmove(ontouchmove func(*TouchEvent))

	// // Ontouchstart prop
	// // js:"ontouchstart,rewrite=ontouchstart"
	// Ontouchstart() (ontouchstart func(*TouchEvent))

	// // Ontouchstart prop
	// // js:"setontouchstart,rewrite=setontouchstart"
	// SetOntouchstart(ontouchstart func(*TouchEvent))

	// // Onwebkitfullscreenchange prop
	// // js:"onwebkitfullscreenchange,rewrite=onwebkitfullscreenchange"
	// Onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event))

	// // Onwebkitfullscreenchange prop
	// // js:"setonwebkitfullscreenchange,rewrite=setonwebkitfullscreenchange"
	// SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(Event))

	// // Onwebkitfullscreenerror prop
	// // js:"onwebkitfullscreenerror,rewrite=onwebkitfullscreenerror"
	// Onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event))

	// // Onwebkitfullscreenerror prop
	// // js:"setonwebkitfullscreenerror,rewrite=setonwebkitfullscreenerror"
	// SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(Event))

	// // OuterHTML prop
	// // js:"outerHTML,rewrite=outerhtml"
	// OuterHTML() (outerHTML string)

	// // OuterHTML prop
	// // js:"setouterHTML,rewrite=setouterhtml"
	// SetOuterHTML(outerHTML string)

	// // Prefix prop
	// // js:"prefix,rewrite=prefix"
	// Prefix() (prefix string)

	// // ScrollHeight prop
	// // js:"scrollHeight,rewrite=scrollheight"
	// ScrollHeight() (scrollHeight int)

	// // ScrollLeft prop
	// // js:"scrollLeft,rewrite=scrollleft"
	// ScrollLeft() (scrollLeft int)

	// // ScrollLeft prop
	// // js:"setscrollLeft,rewrite=setscrollleft"
	// SetScrollLeft(scrollLeft int)

	// // ScrollTop prop
	// // js:"scrollTop,rewrite=scrolltop"
	// ScrollTop() (scrollTop int)

	// // ScrollTop prop
	// // js:"setscrollTop,rewrite=setscrolltop"
	// SetScrollTop(scrollTop int)

	// // ScrollWidth prop
	// // js:"scrollWidth,rewrite=scrollwidth"
	// ScrollWidth() (scrollWidth int)

	// // TagName prop
	// // js:"tagName,rewrite=tagname"
	// TagName() (tagName string)
}

// getattribute fn
func getattribute(qualifiedName string) (s string) {
	js.Rewrite("$<.getAttribute($1)", qualifiedName)
	return s
}

// getattributenode fn
func getattributenode(name string) (a *Attr) {
	js.Rewrite("$<.getAttributeNode($1)", name)
	return a
}

// getattributenodens fn
func getattributenodens(namespaceURI string, localName string) (a *Attr) {
	js.Rewrite("$<.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return a
}

// getattributens fn
func getattributens(namespaceURI string, localName string) (s string) {
	js.Rewrite("$<.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// getboundingclientrect fn
func getboundingclientrect() (c *clientrect.ClientRect) {
	js.Rewrite("$<.getBoundingClientRect()")
	return c
}

// getclientrects fn
func getclientrects() (c *clientrectlist.ClientRectList) {
	js.Rewrite("$<.getClientRects()")
	return c
}

// // getelementsbytagname fn
// func getelementsbytagname(name string) (n *NodeList) {
// 	js.Rewrite("$<.getElementsByTagName($1)", name)
// 	return n
// }

// // getelementsbytagnamens fn
// func getelementsbytagnamens(namespaceURI string, localName string) (n *NodeList) {
// 	js.Rewrite("$<.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
// 	return n
// }

// hasattribute fn
func hasattribute(name string) (b bool) {
	js.Rewrite("$<.hasAttribute($1)", name)
	return b
}

// hasattributens fn
func hasattributens(namespaceURI string, localName string) (b bool) {
	js.Rewrite("$<.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// msgetregioncontent fn
func msgetregioncontent() (m *MSRangeCollection) {
	js.Rewrite("$<.msGetRegionContent()")
	return m
}

// msgetuntransformedbounds fn
func msgetuntransformedbounds() (c *clientrect.ClientRect) {
	js.Rewrite("$<.msGetUntransformedBounds()")
	return c
}

// msmatchesselector fn
func msmatchesselector(selectors string) (b bool) {
	js.Rewrite("$<.msMatchesSelector($1)", selectors)
	return b
}

// msreleasepointercapture fn
func msreleasepointercapture(pointerId int) {
	js.Rewrite("$<.msReleasePointerCapture($1)", pointerId)
}

// mssetpointercapture fn
func mssetpointercapture(pointerId int) {
	js.Rewrite("$<.msSetPointerCapture($1)", pointerId)
}

// mszoomto fn
func mszoomto(args *mszoomtooptions.MsZoomToOptions) {
	js.Rewrite("$<.msZoomTo($1)", args)
}

// releasepointercapture fn
func releasepointercapture(pointerId int) {
	js.Rewrite("$<.releasePointerCapture($1)", pointerId)
}

// removeattribute fn
func removeattribute(qualifiedName string) {
	js.Rewrite("$<.removeAttribute($1)", qualifiedName)
}

// removeattributenode fn
func removeattributenode(oldAttr *Attr) (a *Attr) {
	js.Rewrite("$<.removeAttributeNode($1)", oldAttr)
	return a
}

// removeattributens fn
func removeattributens(namespaceURI string, localName string) {
	js.Rewrite("$<.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// requestfullscreen fn
func requestfullscreen() {
	js.Rewrite("$<.requestFullscreen()")
}

// requestpointerlock fn
func requestpointerlock() {
	js.Rewrite("$<.requestPointerLock()")
}

// setattribute fn
func setattribute(qualifiedName string, value string) {
	js.Rewrite("$<.setAttribute($1, $2)", qualifiedName, value)
}

// setattributenode fn
func setattributenode(newAttr *Attr) (a *Attr) {
	js.Rewrite("$<.setAttributeNode($1)", newAttr)
	return a
}

// setattributenodens fn
func setattributenodens(newAttr *Attr) (a *Attr) {
	js.Rewrite("$<.setAttributeNodeNS($1)", newAttr)
	return a
}

// setattributens fn
func setattributens(namespaceURI string, qualifiedName string, value string) {
	js.Rewrite("$<.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// setpointercapture fn
func setpointercapture(pointerId int) {
	js.Rewrite("$<.setPointerCapture($1)", pointerId)
}

// webkitmatchesselector fn
func webkitmatchesselector(selectors string) (b bool) {
	js.Rewrite("$<.webkitMatchesSelector($1)", selectors)
	return b
}

// // webkitrequestfullscreen fn
// func webkitrequestfullscreen() {
// 	js.Rewrite("$<.webkitRequestFullscreen()")
// }

// // webkitrequestfullscreen fn
// func webkitrequestfullscreen() {
// 	js.Rewrite("$<.webkitRequestFullScreen()")
// }

// classlist prop
func classlist() (classList domtokenlist.DOMTokenList) {
	js.Rewrite("$<.classList")
	return classList
}

// classname prop
func classname() (className string) {
	js.Rewrite("$<.className")
	return className
}

// setclassname prop
func setclassname(className string) {
	js.Rewrite("$<.className = className")
}

// clientheight prop
func clientheight() (clientHeight int) {
	js.Rewrite("$<.clientHeight")
	return clientHeight
}

// clientleft prop
func clientleft() (clientLeft int) {
	js.Rewrite("$<.clientLeft")
	return clientLeft
}

// clienttop prop
func clienttop() (clientTop int) {
	js.Rewrite("$<.clientTop")
	return clientTop
}

// clientwidth prop
func clientwidth() (clientWidth int) {
	js.Rewrite("$<.clientWidth")
	return clientWidth
}

// id prop
func id() (id string) {
	js.Rewrite("$<.id")
	return id
}

// setid prop
func setid(id string) {
	js.Rewrite("$<.id = id")
}

// innerhtml prop
func innerhtml() (innerHTML string) {
	js.Rewrite("$<.innerHTML")
	return innerHTML
}

// setinnerhtml prop
func setinnerhtml(innerHTML string) {
	js.Rewrite("$<.innerHTML = innerHTML")
}

// mscontentzoomfactor prop
func mscontentzoomfactor() (msContentZoomFactor float32) {
	js.Rewrite("$<.msContentZoomFactor")
	return msContentZoomFactor
}

// setmscontentzoomfactor prop
func setmscontentzoomfactor(msContentZoomFactor float32) {
	js.Rewrite("$<.msContentZoomFactor = msContentZoomFactor")
}

// msregionoverflow prop
func msregionoverflow() (msRegionOverflow string) {
	js.Rewrite("$<.msRegionOverflow")
	return msRegionOverflow
}

// onariarequest prop
func onariarequest() (onariarequest func(Event)) {
	js.Rewrite("$<.onariarequest")
	return onariarequest
}

// setonariarequest prop
func setonariarequest(onariarequest func(Event)) {
	js.Rewrite("$<.onariarequest = onariarequest")
}

// oncommand prop
func oncommand() (oncommand func(Event)) {
	js.Rewrite("$<.oncommand")
	return oncommand
}

// setoncommand prop
func setoncommand(oncommand func(Event)) {
	js.Rewrite("$<.oncommand = oncommand")
}

// ongotpointercapture prop
func ongotpointercapture() (ongotpointercapture func(*PointerEvent)) {
	js.Rewrite("$<.ongotpointercapture")
	return ongotpointercapture
}

// setongotpointercapture prop
func setongotpointercapture(ongotpointercapture func(*PointerEvent)) {
	js.Rewrite("$<.ongotpointercapture = ongotpointercapture")
}

// onlostpointercapture prop
func onlostpointercapture() (onlostpointercapture func(*PointerEvent)) {
	js.Rewrite("$<.onlostpointercapture")
	return onlostpointercapture
}

// setonlostpointercapture prop
func setonlostpointercapture(onlostpointercapture func(*PointerEvent)) {
	js.Rewrite("$<.onlostpointercapture = onlostpointercapture")
}

// // onmsgesturechange prop
// func onmsgesturechange() (onmsgesturechange func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsgesturechange")
// 	return onmsgesturechange
// }

// // setonmsgesturechange prop
// func setonmsgesturechange(onmsgesturechange func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsgesturechange = onmsgesturechange")
// }

// // onmsgesturedoubletap prop
// func onmsgesturedoubletap() (onmsgesturedoubletap func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsgesturedoubletap")
// 	return onmsgesturedoubletap
// }

// // setonmsgesturedoubletap prop
// func setonmsgesturedoubletap(onmsgesturedoubletap func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsgesturedoubletap = onmsgesturedoubletap")
// }

// // onmsgestureend prop
// func onmsgestureend() (onmsgestureend func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsgestureend")
// 	return onmsgestureend
// }

// // setonmsgestureend prop
// func setonmsgestureend(onmsgestureend func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsgestureend = onmsgestureend")
// }

// // onmsgesturehold prop
// func onmsgesturehold() (onmsgesturehold func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsgesturehold")
// 	return onmsgesturehold
// }

// // setonmsgesturehold prop
// func setonmsgesturehold(onmsgesturehold func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsgesturehold = onmsgesturehold")
// }

// // onmsgesturestart prop
// func onmsgesturestart() (onmsgesturestart func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsgesturestart")
// 	return onmsgesturestart
// }

// // setonmsgesturestart prop
// func setonmsgesturestart(onmsgesturestart func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsgesturestart = onmsgesturestart")
// }

// // onmsgesturetap prop
// func onmsgesturetap() (onmsgesturetap func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsgesturetap")
// 	return onmsgesturetap
// }

// // setonmsgesturetap prop
// func setonmsgesturetap(onmsgesturetap func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsgesturetap = onmsgesturetap")
// }

// onmsgotpointercapture prop
func onmsgotpointercapture() (onmsgotpointercapture func(*MSPointerEvent)) {
	js.Rewrite("$<.onmsgotpointercapture")
	return onmsgotpointercapture
}

// setonmsgotpointercapture prop
func setonmsgotpointercapture(onmsgotpointercapture func(*MSPointerEvent)) {
	js.Rewrite("$<.onmsgotpointercapture = onmsgotpointercapture")
}

// // onmsinertiastart prop
// func onmsinertiastart() (onmsinertiastart func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsinertiastart")
// 	return onmsinertiastart
// }

// // setonmsinertiastart prop
// func setonmsinertiastart(onmsinertiastart func(*MSGestureEvent)) {
// 	js.Rewrite("$<.onmsinertiastart = onmsinertiastart")
// }

// onmslostpointercapture prop
func onmslostpointercapture() (onmslostpointercapture func(*MSPointerEvent)) {
	js.Rewrite("$<.onmslostpointercapture")
	return onmslostpointercapture
}

// setonmslostpointercapture prop
func setonmslostpointercapture(onmslostpointercapture func(*MSPointerEvent)) {
	js.Rewrite("$<.onmslostpointercapture = onmslostpointercapture")
}

// // onmspointercancel prop
// func onmspointercancel() (onmspointercancel func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointercancel")
// 	return onmspointercancel
// }

// // setonmspointercancel prop
// func setonmspointercancel(onmspointercancel func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointercancel = onmspointercancel")
// }

// // onmspointerdown prop
// func onmspointerdown() (onmspointerdown func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointerdown")
// 	return onmspointerdown
// }

// // setonmspointerdown prop
// func setonmspointerdown(onmspointerdown func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointerdown = onmspointerdown")
// }

// // onmspointerenter prop
// func onmspointerenter() (onmspointerenter func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointerenter")
// 	return onmspointerenter
// }

// // setonmspointerenter prop
// func setonmspointerenter(onmspointerenter func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointerenter = onmspointerenter")
// }

// // onmspointerleave prop
// func onmspointerleave() (onmspointerleave func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointerleave")
// 	return onmspointerleave
// }

// // setonmspointerleave prop
// func setonmspointerleave(onmspointerleave func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointerleave = onmspointerleave")
// }

// // onmspointermove prop
// func onmspointermove() (onmspointermove func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointermove")
// 	return onmspointermove
// }

// // setonmspointermove prop
// func setonmspointermove(onmspointermove func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointermove = onmspointermove")
// }

// // onmspointerout prop
// func onmspointerout() (onmspointerout func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointerout")
// 	return onmspointerout
// }

// // setonmspointerout prop
// func setonmspointerout(onmspointerout func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointerout = onmspointerout")
// }

// // onmspointerover prop
// func onmspointerover() (onmspointerover func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointerover")
// 	return onmspointerover
// }

// // setonmspointerover prop
// func setonmspointerover(onmspointerover func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointerover = onmspointerover")
// }

// // onmspointerup prop
// func onmspointerup() (onmspointerup func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointerup")
// 	return onmspointerup
// }

// // setonmspointerup prop
// func setonmspointerup(onmspointerup func(*MSPointerEvent)) {
// 	js.Rewrite("$<.onmspointerup = onmspointerup")
// }

// // ontouchcancel prop
// func ontouchcancel() (ontouchcancel func(*TouchEvent)) {
// 	js.Rewrite("$<.ontouchcancel")
// 	return ontouchcancel
// }

// // setontouchcancel prop
// func setontouchcancel(ontouchcancel func(*TouchEvent)) {
// 	js.Rewrite("$<.ontouchcancel = ontouchcancel")
// }

// // ontouchend prop
// func ontouchend() (ontouchend func(*TouchEvent)) {
// 	js.Rewrite("$<.ontouchend")
// 	return ontouchend
// }

// // setontouchend prop
// func setontouchend(ontouchend func(*TouchEvent)) {
// 	js.Rewrite("$<.ontouchend = ontouchend")
// }

// // ontouchmove prop
// func ontouchmove() (ontouchmove func(*TouchEvent)) {
// 	js.Rewrite("$<.ontouchmove")
// 	return ontouchmove
// }

// // setontouchmove prop
// func setontouchmove(ontouchmove func(*TouchEvent)) {
// 	js.Rewrite("$<.ontouchmove = ontouchmove")
// }

// // ontouchstart prop
// func ontouchstart() (ontouchstart func(*TouchEvent)) {
// 	js.Rewrite("$<.ontouchstart")
// 	return ontouchstart
// }

// // setontouchstart prop
// func setontouchstart(ontouchstart func(*TouchEvent)) {
// 	js.Rewrite("$<.ontouchstart = ontouchstart")
// }

// // onwebkitfullscreenchange prop
// func onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event)) {
// 	js.Rewrite("$<.onwebkitfullscreenchange")
// 	return onwebkitfullscreenchange
// }

// // setonwebkitfullscreenchange prop
// func setonwebkitfullscreenchange(onwebkitfullscreenchange func(Event)) {
// 	js.Rewrite("$<.onwebkitfullscreenchange = onwebkitfullscreenchange")
// }

// // onwebkitfullscreenerror prop
// func onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event)) {
// 	js.Rewrite("$<.onwebkitfullscreenerror")
// 	return onwebkitfullscreenerror
// }

// // setonwebkitfullscreenerror prop
// func setonwebkitfullscreenerror(onwebkitfullscreenerror func(Event)) {
// 	js.Rewrite("$<.onwebkitfullscreenerror = onwebkitfullscreenerror")
// }

// outerhtml prop
func outerhtml() (outerHTML string) {
	js.Rewrite("$<.outerHTML")
	return outerHTML
}

// setouterhtml prop
func setouterhtml(outerHTML string) {
	js.Rewrite("$<.outerHTML = outerHTML")
}

// prefix prop
func prefix() (prefix string) {
	js.Rewrite("$<.prefix")
	return prefix
}

// scrollheight prop
func scrollheight() (scrollHeight int) {
	js.Rewrite("$<.scrollHeight")
	return scrollHeight
}

// scrollleft prop
func scrollleft() (scrollLeft int) {
	js.Rewrite("$<.scrollLeft")
	return scrollLeft
}

// setscrollleft prop
func setscrollleft(scrollLeft int) {
	js.Rewrite("$<.scrollLeft = scrollLeft")
}

// scrolltop prop
func scrolltop() (scrollTop int) {
	js.Rewrite("$<.scrollTop")
	return scrollTop
}

// setscrolltop prop
func setscrolltop(scrollTop int) {
	js.Rewrite("$<.scrollTop = scrollTop")
}

// scrollwidth prop
func scrollwidth() (scrollWidth int) {
	js.Rewrite("$<.scrollWidth")
	return scrollWidth
}

// tagname prop
func tagname() (tagName string) {
	js.Rewrite("$<.tagName")
	return tagName
}
