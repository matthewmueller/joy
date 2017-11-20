package element

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/attr"
	"github.com/matthewmueller/golly/internal/gendom/dom/childnode"
	"github.com/matthewmueller/golly/internal/gendom/dom/clientrect"
	"github.com/matthewmueller/golly/internal/gendom/dom/clientrectlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/domtokenlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/elementtraversal"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/globaleventhandlers"
	"github.com/matthewmueller/golly/internal/gendom/dom/msgestureevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/mspointerevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/msrangecollection"
	"github.com/matthewmueller/golly/internal/gendom/dom/mszoomtooptions"
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodelist"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodeselector"
	"github.com/matthewmueller/golly/internal/gendom/dom/pointerevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/touchevent"
	"github.com/matthewmueller/golly/js"
)

type Element struct {
	*node.Node
	*globaleventhandlers.GlobalEventHandlers
	*elementtraversal.ElementTraversal
	*nodeselector.NodeSelector
	*childnode.ChildNode
}

func (*Element) GetAttribute(qualifiedName string) (s string) {
	js.Rewrite("$<.getAttribute($1)", qualifiedName)
	return s
}

func (*Element) GetAttributeNode(name string) (a *attr.Attr) {
	js.Rewrite("$<.getAttributeNode($1)", name)
	return a
}

func (*Element) GetAttributeNodeNS(namespaceURI string, localName string) (a *attr.Attr) {
	js.Rewrite("$<.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return a
}

func (*Element) GetAttributeNS(namespaceURI string, localName string) (s string) {
	js.Rewrite("$<.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

func (*Element) GetBoundingClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$<.getBoundingClientRect()")
	return c
}

func (*Element) GetClientRects() (c *clientrectlist.ClientRectList) {
	js.Rewrite("$<.getClientRects()")
	return c
}

func (*Element) GetElementsByTagName(name string) (n *nodelist.NodeList) {
	js.Rewrite("$<.getElementsByTagName($1)", name)
	return n
}

func (*Element) GetElementsByTagNameNS(namespaceURI string, localName string) (n *nodelist.NodeList) {
	js.Rewrite("$<.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return n
}

func (*Element) HasAttribute(name string) (b bool) {
	js.Rewrite("$<.hasAttribute($1)", name)
	return b
}

func (*Element) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	js.Rewrite("$<.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

func (*Element) MsGetRegionContent() (m *msrangecollection.MSRangeCollection) {
	js.Rewrite("$<.msGetRegionContent()")
	return m
}

func (*Element) MsGetUntransformedBounds() (c *clientrect.ClientRect) {
	js.Rewrite("$<.msGetUntransformedBounds()")
	return c
}

func (*Element) MsMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$<.msMatchesSelector($1)", selectors)
	return b
}

func (*Element) MsReleasePointerCapture(pointerId int) {
	js.Rewrite("$<.msReleasePointerCapture($1)", pointerId)
}

func (*Element) MsSetPointerCapture(pointerId int) {
	js.Rewrite("$<.msSetPointerCapture($1)", pointerId)
}

func (*Element) MsZoomTo(args *mszoomtooptions.MsZoomToOptions) {
	js.Rewrite("$<.msZoomTo($1)", args)
}

func (*Element) ReleasePointerCapture(pointerId int) {
	js.Rewrite("$<.releasePointerCapture($1)", pointerId)
}

func (*Element) RemoveAttribute(qualifiedName string) {
	js.Rewrite("$<.removeAttribute($1)", qualifiedName)
}

func (*Element) RemoveAttributeNode(oldAttr *attr.Attr) (a *attr.Attr) {
	js.Rewrite("$<.removeAttributeNode($1)", oldAttr)
	return a
}

func (*Element) RemoveAttributeNS(namespaceURI string, localName string) {
	js.Rewrite("$<.removeAttributeNS($1, $2)", namespaceURI, localName)
}

func (*Element) RequestFullscreen() {
	js.Rewrite("$<.requestFullscreen()")
}

func (*Element) RequestPointerLock() {
	js.Rewrite("$<.requestPointerLock()")
}

func (*Element) SetAttribute(qualifiedName string, value string) {
	js.Rewrite("$<.setAttribute($1, $2)", qualifiedName, value)
}

func (*Element) SetAttributeNode(newAttr *attr.Attr) (a *attr.Attr) {
	js.Rewrite("$<.setAttributeNode($1)", newAttr)
	return a
}

func (*Element) SetAttributeNodeNS(newAttr *attr.Attr) (a *attr.Attr) {
	js.Rewrite("$<.setAttributeNodeNS($1)", newAttr)
	return a
}

func (*Element) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	js.Rewrite("$<.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

func (*Element) SetPointerCapture(pointerId int) {
	js.Rewrite("$<.setPointerCapture($1)", pointerId)
}

func (*Element) WebkitMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$<.webkitMatchesSelector($1)", selectors)
	return b
}

func (*Element) WebkitRequestFullscreen() {
	js.Rewrite("$<.webkitRequestFullscreen()")
}

func (*Element) WebkitRequestFullScreen() {
	js.Rewrite("$<.webkitRequestFullScreen()")
}

func (*Element) GetClassList() (classList *domtokenlist.DOMTokenList) {
	js.Rewrite("$<.classList")
	return classList
}

func (*Element) GetClassName() (className string) {
	js.Rewrite("$<.className")
	return className
}

func (*Element) SetClassName(className string) {
	js.Rewrite("$<.className = $1", className)
}

func (*Element) GetClientHeight() (clientHeight int) {
	js.Rewrite("$<.clientHeight")
	return clientHeight
}

func (*Element) GetClientLeft() (clientLeft int) {
	js.Rewrite("$<.clientLeft")
	return clientLeft
}

func (*Element) GetClientTop() (clientTop int) {
	js.Rewrite("$<.clientTop")
	return clientTop
}

func (*Element) GetClientWidth() (clientWidth int) {
	js.Rewrite("$<.clientWidth")
	return clientWidth
}

func (*Element) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*Element) SetID(id string) {
	js.Rewrite("$<.id = $1", id)
}

func (*Element) GetInnerHTML() (innerHTML string) {
	js.Rewrite("$<.innerHTML")
	return innerHTML
}

func (*Element) SetInnerHTML(innerHTML string) {
	js.Rewrite("$<.innerHTML = $1", innerHTML)
}

func (*Element) GetMsContentZoomFactor() (msContentZoomFactor float32) {
	js.Rewrite("$<.msContentZoomFactor")
	return msContentZoomFactor
}

func (*Element) SetMsContentZoomFactor(msContentZoomFactor float32) {
	js.Rewrite("$<.msContentZoomFactor = $1", msContentZoomFactor)
}

func (*Element) GetMsRegionOverflow() (msRegionOverflow string) {
	js.Rewrite("$<.msRegionOverflow")
	return msRegionOverflow
}

func (*Element) GetOnariarequest() (e *event.Event) {
	js.Rewrite("$<.onariarequest")
	return e
}

func (*Element) SetOnariarequest(e *event.Event) {
	js.Rewrite("$<.onariarequest = $1", e)
}

func (*Element) GetOncommand() (e *event.Event) {
	js.Rewrite("$<.oncommand")
	return e
}

func (*Element) SetOncommand(e *event.Event) {
	js.Rewrite("$<.oncommand = $1", e)
}

func (*Element) GetOngotpointercapture() (gotpointercapture *pointerevent.PointerEvent) {
	js.Rewrite("$<.ongotpointercapture")
	return gotpointercapture
}

func (*Element) SetOngotpointercapture(gotpointercapture *pointerevent.PointerEvent) {
	js.Rewrite("$<.ongotpointercapture = $1", gotpointercapture)
}

func (*Element) GetOnlostpointercapture() (lostpointercapture *pointerevent.PointerEvent) {
	js.Rewrite("$<.onlostpointercapture")
	return lostpointercapture
}

func (*Element) SetOnlostpointercapture(lostpointercapture *pointerevent.PointerEvent) {
	js.Rewrite("$<.onlostpointercapture = $1", lostpointercapture)
}

func (*Element) GetOnmsgesturechange() (MSGestureChange *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsgesturechange")
	return MSGestureChange
}

func (*Element) SetOnmsgesturechange(MSGestureChange *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsgesturechange = $1", MSGestureChange)
}

func (*Element) GetOnmsgesturedoubletap() (MSGestureDoubleTap *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsgesturedoubletap")
	return MSGestureDoubleTap
}

func (*Element) SetOnmsgesturedoubletap(MSGestureDoubleTap *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsgesturedoubletap = $1", MSGestureDoubleTap)
}

func (*Element) GetOnmsgestureend() (MSGestureEnd *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsgestureend")
	return MSGestureEnd
}

func (*Element) SetOnmsgestureend(MSGestureEnd *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsgestureend = $1", MSGestureEnd)
}

func (*Element) GetOnmsgesturehold() (MSGestureHold *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsgesturehold")
	return MSGestureHold
}

func (*Element) SetOnmsgesturehold(MSGestureHold *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsgesturehold = $1", MSGestureHold)
}

func (*Element) GetOnmsgesturestart() (MSGestureStart *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsgesturestart")
	return MSGestureStart
}

func (*Element) SetOnmsgesturestart(MSGestureStart *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsgesturestart = $1", MSGestureStart)
}

func (*Element) GetOnmsgesturetap() (MSGestureTap *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsgesturetap")
	return MSGestureTap
}

func (*Element) SetOnmsgesturetap(MSGestureTap *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsgesturetap = $1", MSGestureTap)
}

func (*Element) GetOnmsgotpointercapture() (MSGotPointerCapture *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmsgotpointercapture")
	return MSGotPointerCapture
}

func (*Element) SetOnmsgotpointercapture(MSGotPointerCapture *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmsgotpointercapture = $1", MSGotPointerCapture)
}

func (*Element) GetOnmsinertiastart() (MSInertiaStart *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsinertiastart")
	return MSInertiaStart
}

func (*Element) SetOnmsinertiastart(MSInertiaStart *msgestureevent.MSGestureEvent) {
	js.Rewrite("$<.onmsinertiastart = $1", MSInertiaStart)
}

func (*Element) GetOnmslostpointercapture() (MSLostPointerCapture *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmslostpointercapture")
	return MSLostPointerCapture
}

func (*Element) SetOnmslostpointercapture(MSLostPointerCapture *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmslostpointercapture = $1", MSLostPointerCapture)
}

func (*Element) GetOnmspointercancel() (MSPointerCancel *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointercancel")
	return MSPointerCancel
}

func (*Element) SetOnmspointercancel(MSPointerCancel *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointercancel = $1", MSPointerCancel)
}

func (*Element) GetOnmspointerdown() (MSPointerDown *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointerdown")
	return MSPointerDown
}

func (*Element) SetOnmspointerdown(MSPointerDown *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointerdown = $1", MSPointerDown)
}

func (*Element) GetOnmspointerenter() (MSPointerEnter *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointerenter")
	return MSPointerEnter
}

func (*Element) SetOnmspointerenter(MSPointerEnter *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointerenter = $1", MSPointerEnter)
}

func (*Element) GetOnmspointerleave() (MSPointerLeave *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointerleave")
	return MSPointerLeave
}

func (*Element) SetOnmspointerleave(MSPointerLeave *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointerleave = $1", MSPointerLeave)
}

func (*Element) GetOnmspointermove() (MSPointerMove *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointermove")
	return MSPointerMove
}

func (*Element) SetOnmspointermove(MSPointerMove *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointermove = $1", MSPointerMove)
}

func (*Element) GetOnmspointerout() (MSPointerOut *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointerout")
	return MSPointerOut
}

func (*Element) SetOnmspointerout(MSPointerOut *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointerout = $1", MSPointerOut)
}

func (*Element) GetOnmspointerover() (MSPointerOver *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointerover")
	return MSPointerOver
}

func (*Element) SetOnmspointerover(MSPointerOver *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointerover = $1", MSPointerOver)
}

func (*Element) GetOnmspointerup() (MSPointerUp *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointerup")
	return MSPointerUp
}

func (*Element) SetOnmspointerup(MSPointerUp *mspointerevent.MSPointerEvent) {
	js.Rewrite("$<.onmspointerup = $1", MSPointerUp)
}

func (*Element) GetOntouchcancel() (touchcancel *touchevent.TouchEvent) {
	js.Rewrite("$<.ontouchcancel")
	return touchcancel
}

func (*Element) SetOntouchcancel(touchcancel *touchevent.TouchEvent) {
	js.Rewrite("$<.ontouchcancel = $1", touchcancel)
}

func (*Element) GetOntouchend() (touchend *touchevent.TouchEvent) {
	js.Rewrite("$<.ontouchend")
	return touchend
}

func (*Element) SetOntouchend(touchend *touchevent.TouchEvent) {
	js.Rewrite("$<.ontouchend = $1", touchend)
}

func (*Element) GetOntouchmove() (touchmove *touchevent.TouchEvent) {
	js.Rewrite("$<.ontouchmove")
	return touchmove
}

func (*Element) SetOntouchmove(touchmove *touchevent.TouchEvent) {
	js.Rewrite("$<.ontouchmove = $1", touchmove)
}

func (*Element) GetOntouchstart() (touchstart *touchevent.TouchEvent) {
	js.Rewrite("$<.ontouchstart")
	return touchstart
}

func (*Element) SetOntouchstart(touchstart *touchevent.TouchEvent) {
	js.Rewrite("$<.ontouchstart = $1", touchstart)
}

func (*Element) GetOnwebkitfullscreenchange() (e *event.Event) {
	js.Rewrite("$<.onwebkitfullscreenchange")
	return e
}

func (*Element) SetOnwebkitfullscreenchange(e *event.Event) {
	js.Rewrite("$<.onwebkitfullscreenchange = $1", e)
}

func (*Element) GetOnwebkitfullscreenerror() (e *event.Event) {
	js.Rewrite("$<.onwebkitfullscreenerror")
	return e
}

func (*Element) SetOnwebkitfullscreenerror(e *event.Event) {
	js.Rewrite("$<.onwebkitfullscreenerror = $1", e)
}

func (*Element) GetOuterHTML() (outerHTML string) {
	js.Rewrite("$<.outerHTML")
	return outerHTML
}

func (*Element) SetOuterHTML(outerHTML string) {
	js.Rewrite("$<.outerHTML = $1", outerHTML)
}

func (*Element) GetPrefix() (prefix string) {
	js.Rewrite("$<.prefix")
	return prefix
}

func (*Element) GetScrollHeight() (scrollHeight int) {
	js.Rewrite("$<.scrollHeight")
	return scrollHeight
}

func (*Element) GetScrollLeft() (scrollLeft int) {
	js.Rewrite("$<.scrollLeft")
	return scrollLeft
}

func (*Element) SetScrollLeft(scrollLeft int) {
	js.Rewrite("$<.scrollLeft = $1", scrollLeft)
}

func (*Element) GetScrollTop() (scrollTop int) {
	js.Rewrite("$<.scrollTop")
	return scrollTop
}

func (*Element) SetScrollTop(scrollTop int) {
	js.Rewrite("$<.scrollTop = $1", scrollTop)
}

func (*Element) GetScrollWidth() (scrollWidth int) {
	js.Rewrite("$<.scrollWidth")
	return scrollWidth
}

func (*Element) GetTagName() (tagName string) {
	js.Rewrite("$<.tagName")
	return tagName
}
