package window

// js:"Element,omit"
type Element interface {
	Node
	GlobalEventHandlers
	ElementTraversal
	NodeSelector
	childnode.ChildNode

	// GetAttribute
	GetAttribute(qualifiedName string) (s string)

	// GetAttributeNode
	GetAttributeNode(name string) (a *Attr)

	// GetAttributeNodeNS
	GetAttributeNodeNS(namespaceURI string, localName string) (a *Attr)

	// GetAttributeNS
	GetAttributeNS(namespaceURI string, localName string) (s string)

	// GetBoundingClientRect
	GetBoundingClientRect() (c *clientrect.ClientRect)

	// GetClientRects
	GetClientRects() (c *clientrectlist.ClientRectList)

	// GetElementsByTagName
	GetElementsByTagName(name string) (n *NodeList)

	// GetElementsByTagNameNS
	GetElementsByTagNameNS(namespaceURI string, localName string) (n *NodeList)

	// HasAttribute
	HasAttribute(name string) (b bool)

	// HasAttributeNS
	HasAttributeNS(namespaceURI string, localName string) (b bool)

	// MsGetRegionContent
	MsGetRegionContent() (m *MSRangeCollection)

	// MsGetUntransformedBounds
	MsGetUntransformedBounds() (c *clientrect.ClientRect)

	// MsMatchesSelector
	MsMatchesSelector(selectors string) (b bool)

	// MsReleasePointerCapture
	MsReleasePointerCapture(pointerId int)

	// MsSetPointerCapture
	MsSetPointerCapture(pointerId int)

	// MsZoomTo
	MsZoomTo(args *mszoomtooptions.MsZoomToOptions)

	// ReleasePointerCapture
	ReleasePointerCapture(pointerId int)

	// RemoveAttribute
	RemoveAttribute(qualifiedName string)

	// RemoveAttributeNode
	RemoveAttributeNode(oldAttr *Attr) (a *Attr)

	// RemoveAttributeNS
	RemoveAttributeNS(namespaceURI string, localName string)

	// RequestFullscreen
	RequestFullscreen()

	// RequestPointerLock
	RequestPointerLock()

	// SetAttribute
	SetAttribute(qualifiedName string, value string)

	// SetAttributeNode
	SetAttributeNode(newAttr *Attr) (a *Attr)

	// SetAttributeNodeNS
	SetAttributeNodeNS(newAttr *Attr) (a *Attr)

	// SetAttributeNS
	SetAttributeNS(namespaceURI string, qualifiedName string, value string)

	// SetPointerCapture
	SetPointerCapture(pointerId int)

	// WebkitMatchesSelector
	WebkitMatchesSelector(selectors string) (b bool)

	// WebkitRequestFullscreen
	WebkitRequestFullscreen()

	// WebkitRequestFullScreen
	WebkitRequestFullScreen()

	// ClassList
	ClassList() (classList domtokenlist.DOMTokenList)

	// ClassName
	ClassName() (className string)

	// ClassName
	SetClassName(className string)

	// ClientHeight
	ClientHeight() (clientHeight int)

	// ClientLeft
	ClientLeft() (clientLeft int)

	// ClientTop
	ClientTop() (clientTop int)

	// ClientWidth
	ClientWidth() (clientWidth int)

	// ID
	ID() (id string)

	// ID
	SetID(id string)

	// InnerHTML
	InnerHTML() (innerHTML string)

	// InnerHTML
	SetInnerHTML(innerHTML string)

	// MsContentZoomFactor
	MsContentZoomFactor() (msContentZoomFactor float32)

	// MsContentZoomFactor
	SetMsContentZoomFactor(msContentZoomFactor float32)

	// MsRegionOverflow
	MsRegionOverflow() (msRegionOverflow string)

	// Onariarequest
	Onariarequest() (onariarequest func(Event))

	// Onariarequest
	SetOnariarequest(onariarequest func(Event))

	// Oncommand
	Oncommand() (oncommand func(Event))

	// Oncommand
	SetOncommand(oncommand func(Event))

	// Ongotpointercapture
	Ongotpointercapture() (ongotpointercapture func(*PointerEvent))

	// Ongotpointercapture
	SetOngotpointercapture(ongotpointercapture func(*PointerEvent))

	// Onlostpointercapture
	Onlostpointercapture() (onlostpointercapture func(*PointerEvent))

	// Onlostpointercapture
	SetOnlostpointercapture(onlostpointercapture func(*PointerEvent))

	// Onmsgesturechange
	Onmsgesturechange() (onmsgesturechange func(*MSGestureEvent))

	// Onmsgesturechange
	SetOnmsgesturechange(onmsgesturechange func(*MSGestureEvent))

	// Onmsgesturedoubletap
	Onmsgesturedoubletap() (onmsgesturedoubletap func(*MSGestureEvent))

	// Onmsgesturedoubletap
	SetOnmsgesturedoubletap(onmsgesturedoubletap func(*MSGestureEvent))

	// Onmsgestureend
	Onmsgestureend() (onmsgestureend func(*MSGestureEvent))

	// Onmsgestureend
	SetOnmsgestureend(onmsgestureend func(*MSGestureEvent))

	// Onmsgesturehold
	Onmsgesturehold() (onmsgesturehold func(*MSGestureEvent))

	// Onmsgesturehold
	SetOnmsgesturehold(onmsgesturehold func(*MSGestureEvent))

	// Onmsgesturestart
	Onmsgesturestart() (onmsgesturestart func(*MSGestureEvent))

	// Onmsgesturestart
	SetOnmsgesturestart(onmsgesturestart func(*MSGestureEvent))

	// Onmsgesturetap
	Onmsgesturetap() (onmsgesturetap func(*MSGestureEvent))

	// Onmsgesturetap
	SetOnmsgesturetap(onmsgesturetap func(*MSGestureEvent))

	// Onmsgotpointercapture
	Onmsgotpointercapture() (onmsgotpointercapture func(*MSPointerEvent))

	// Onmsgotpointercapture
	SetOnmsgotpointercapture(onmsgotpointercapture func(*MSPointerEvent))

	// Onmsinertiastart
	Onmsinertiastart() (onmsinertiastart func(*MSGestureEvent))

	// Onmsinertiastart
	SetOnmsinertiastart(onmsinertiastart func(*MSGestureEvent))

	// Onmslostpointercapture
	Onmslostpointercapture() (onmslostpointercapture func(*MSPointerEvent))

	// Onmslostpointercapture
	SetOnmslostpointercapture(onmslostpointercapture func(*MSPointerEvent))

	// Onmspointercancel
	Onmspointercancel() (onmspointercancel func(*MSPointerEvent))

	// Onmspointercancel
	SetOnmspointercancel(onmspointercancel func(*MSPointerEvent))

	// Onmspointerdown
	Onmspointerdown() (onmspointerdown func(*MSPointerEvent))

	// Onmspointerdown
	SetOnmspointerdown(onmspointerdown func(*MSPointerEvent))

	// Onmspointerenter
	Onmspointerenter() (onmspointerenter func(*MSPointerEvent))

	// Onmspointerenter
	SetOnmspointerenter(onmspointerenter func(*MSPointerEvent))

	// Onmspointerleave
	Onmspointerleave() (onmspointerleave func(*MSPointerEvent))

	// Onmspointerleave
	SetOnmspointerleave(onmspointerleave func(*MSPointerEvent))

	// Onmspointermove
	Onmspointermove() (onmspointermove func(*MSPointerEvent))

	// Onmspointermove
	SetOnmspointermove(onmspointermove func(*MSPointerEvent))

	// Onmspointerout
	Onmspointerout() (onmspointerout func(*MSPointerEvent))

	// Onmspointerout
	SetOnmspointerout(onmspointerout func(*MSPointerEvent))

	// Onmspointerover
	Onmspointerover() (onmspointerover func(*MSPointerEvent))

	// Onmspointerover
	SetOnmspointerover(onmspointerover func(*MSPointerEvent))

	// Onmspointerup
	Onmspointerup() (onmspointerup func(*MSPointerEvent))

	// Onmspointerup
	SetOnmspointerup(onmspointerup func(*MSPointerEvent))

	// Ontouchcancel
	Ontouchcancel() (ontouchcancel func(*TouchEvent))

	// Ontouchcancel
	SetOntouchcancel(ontouchcancel func(*TouchEvent))

	// Ontouchend
	Ontouchend() (ontouchend func(*TouchEvent))

	// Ontouchend
	SetOntouchend(ontouchend func(*TouchEvent))

	// Ontouchmove
	Ontouchmove() (ontouchmove func(*TouchEvent))

	// Ontouchmove
	SetOntouchmove(ontouchmove func(*TouchEvent))

	// Ontouchstart
	Ontouchstart() (ontouchstart func(*TouchEvent))

	// Ontouchstart
	SetOntouchstart(ontouchstart func(*TouchEvent))

	// Onwebkitfullscreenchange
	Onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event))

	// Onwebkitfullscreenchange
	SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(Event))

	// Onwebkitfullscreenerror
	Onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event))

	// Onwebkitfullscreenerror
	SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(Event))

	// OuterHTML
	OuterHTML() (outerHTML string)

	// OuterHTML
	SetOuterHTML(outerHTML string)

	// Prefix
	Prefix() (prefix string)

	// ScrollHeight
	ScrollHeight() (scrollHeight int)

	// ScrollLeft
	ScrollLeft() (scrollLeft int)

	// ScrollLeft
	SetScrollLeft(scrollLeft int)

	// ScrollTop
	ScrollTop() (scrollTop int)

	// ScrollTop
	SetScrollTop(scrollTop int)

	// ScrollWidth
	ScrollWidth() (scrollWidth int)

	// TagName
	TagName() (tagName string)
}
