package element


		// Element interface
		// js:"Element"
		type Element interface {
			Node
		
			
		// QuerySelector 
		// js:"querySelector"
		QuerySelector(selectors string) (e Element)
	
		// QuerySelectorAll 
		// js:"querySelectorAll"
		QuerySelectorAll(selectors string) (n *NodeList)
	
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
		

		// SetOnpointercancel prop 
		// js:"onpointercancel"
		SetOnpointercancel (onpointercancel func(Event))
		
		// Onpointerdown prop 
		// js:"onpointerdown"
		Onpointerdown() (onpointerdown func(Event))
		

		// SetOnpointerdown prop 
		// js:"onpointerdown"
		SetOnpointerdown (onpointerdown func(Event))
		
		// Onpointerenter prop 
		// js:"onpointerenter"
		Onpointerenter() (onpointerenter func(Event))
		

		// SetOnpointerenter prop 
		// js:"onpointerenter"
		SetOnpointerenter (onpointerenter func(Event))
		
		// Onpointerleave prop 
		// js:"onpointerleave"
		Onpointerleave() (onpointerleave func(Event))
		

		// SetOnpointerleave prop 
		// js:"onpointerleave"
		SetOnpointerleave (onpointerleave func(Event))
		
		// Onpointermove prop 
		// js:"onpointermove"
		Onpointermove() (onpointermove func(Event))
		

		// SetOnpointermove prop 
		// js:"onpointermove"
		SetOnpointermove (onpointermove func(Event))
		
		// Onpointerout prop 
		// js:"onpointerout"
		Onpointerout() (onpointerout func(Event))
		

		// SetOnpointerout prop 
		// js:"onpointerout"
		SetOnpointerout (onpointerout func(Event))
		
		// Onpointerover prop 
		// js:"onpointerover"
		Onpointerover() (onpointerover func(Event))
		

		// SetOnpointerover prop 
		// js:"onpointerover"
		SetOnpointerover (onpointerover func(Event))
		
		// Onpointerup prop 
		// js:"onpointerup"
		Onpointerup() (onpointerup func(Event))
		

		// SetOnpointerup prop 
		// js:"onpointerup"
		SetOnpointerup (onpointerup func(Event))
		
		// Onwheel prop 
		// js:"onwheel"
		Onwheel() (onwheel func(Event))
		

		// SetOnwheel prop 
		// js:"onwheel"
		SetOnwheel (onwheel func(Event))
		
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
		

		// SetClassName prop 
		// js:"className"
		SetClassName (className string)
		
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
		

		// SetID prop 
		// js:"id"
		SetID (id string)
		
		// InnerHTML prop 
		// js:"innerHTML"
		InnerHTML() (innerHTML string)
		

		// SetInnerHTML prop 
		// js:"innerHTML"
		SetInnerHTML (innerHTML string)
		
		// MsContentZoomFactor prop 
		// js:"msContentZoomFactor"
		MsContentZoomFactor() (msContentZoomFactor float32)
		

		// SetMsContentZoomFactor prop 
		// js:"msContentZoomFactor"
		SetMsContentZoomFactor (msContentZoomFactor float32)
		
		// MsRegionOverflow prop 
		// js:"msRegionOverflow"
		MsRegionOverflow() (msRegionOverflow string)
		
		// Onariarequest prop 
		// js:"onariarequest"
		Onariarequest() (onariarequest func(Event))
		

		// SetOnariarequest prop 
		// js:"onariarequest"
		SetOnariarequest (onariarequest func(Event))
		
		// Oncommand prop 
		// js:"oncommand"
		Oncommand() (oncommand func(Event))
		

		// SetOncommand prop 
		// js:"oncommand"
		SetOncommand (oncommand func(Event))
		
		// Ongotpointercapture prop 
		// js:"ongotpointercapture"
		Ongotpointercapture() (ongotpointercapture func(*PointerEvent))
		

		// SetOngotpointercapture prop 
		// js:"ongotpointercapture"
		SetOngotpointercapture (ongotpointercapture func(*PointerEvent))
		
		// Onlostpointercapture prop 
		// js:"onlostpointercapture"
		Onlostpointercapture() (onlostpointercapture func(*PointerEvent))
		

		// SetOnlostpointercapture prop 
		// js:"onlostpointercapture"
		SetOnlostpointercapture (onlostpointercapture func(*PointerEvent))
		
		// Onmsgesturechange prop 
		// js:"onmsgesturechange"
		Onmsgesturechange() (onmsgesturechange func(*MSGestureEvent))
		

		// SetOnmsgesturechange prop 
		// js:"onmsgesturechange"
		SetOnmsgesturechange (onmsgesturechange func(*MSGestureEvent))
		
		// Onmsgesturedoubletap prop 
		// js:"onmsgesturedoubletap"
		Onmsgesturedoubletap() (onmsgesturedoubletap func(*MSGestureEvent))
		

		// SetOnmsgesturedoubletap prop 
		// js:"onmsgesturedoubletap"
		SetOnmsgesturedoubletap (onmsgesturedoubletap func(*MSGestureEvent))
		
		// Onmsgestureend prop 
		// js:"onmsgestureend"
		Onmsgestureend() (onmsgestureend func(*MSGestureEvent))
		

		// SetOnmsgestureend prop 
		// js:"onmsgestureend"
		SetOnmsgestureend (onmsgestureend func(*MSGestureEvent))
		
		// Onmsgesturehold prop 
		// js:"onmsgesturehold"
		Onmsgesturehold() (onmsgesturehold func(*MSGestureEvent))
		

		// SetOnmsgesturehold prop 
		// js:"onmsgesturehold"
		SetOnmsgesturehold (onmsgesturehold func(*MSGestureEvent))
		
		// Onmsgesturestart prop 
		// js:"onmsgesturestart"
		Onmsgesturestart() (onmsgesturestart func(*MSGestureEvent))
		

		// SetOnmsgesturestart prop 
		// js:"onmsgesturestart"
		SetOnmsgesturestart (onmsgesturestart func(*MSGestureEvent))
		
		// Onmsgesturetap prop 
		// js:"onmsgesturetap"
		Onmsgesturetap() (onmsgesturetap func(*MSGestureEvent))
		

		// SetOnmsgesturetap prop 
		// js:"onmsgesturetap"
		SetOnmsgesturetap (onmsgesturetap func(*MSGestureEvent))
		
		// Onmsgotpointercapture prop 
		// js:"onmsgotpointercapture"
		Onmsgotpointercapture() (onmsgotpointercapture func(*MSPointerEvent))
		

		// SetOnmsgotpointercapture prop 
		// js:"onmsgotpointercapture"
		SetOnmsgotpointercapture (onmsgotpointercapture func(*MSPointerEvent))
		
		// Onmsinertiastart prop 
		// js:"onmsinertiastart"
		Onmsinertiastart() (onmsinertiastart func(*MSGestureEvent))
		

		// SetOnmsinertiastart prop 
		// js:"onmsinertiastart"
		SetOnmsinertiastart (onmsinertiastart func(*MSGestureEvent))
		
		// Onmslostpointercapture prop 
		// js:"onmslostpointercapture"
		Onmslostpointercapture() (onmslostpointercapture func(*MSPointerEvent))
		

		// SetOnmslostpointercapture prop 
		// js:"onmslostpointercapture"
		SetOnmslostpointercapture (onmslostpointercapture func(*MSPointerEvent))
		
		// Onmspointercancel prop 
		// js:"onmspointercancel"
		Onmspointercancel() (onmspointercancel func(*MSPointerEvent))
		

		// SetOnmspointercancel prop 
		// js:"onmspointercancel"
		SetOnmspointercancel (onmspointercancel func(*MSPointerEvent))
		
		// Onmspointerdown prop 
		// js:"onmspointerdown"
		Onmspointerdown() (onmspointerdown func(*MSPointerEvent))
		

		// SetOnmspointerdown prop 
		// js:"onmspointerdown"
		SetOnmspointerdown (onmspointerdown func(*MSPointerEvent))
		
		// Onmspointerenter prop 
		// js:"onmspointerenter"
		Onmspointerenter() (onmspointerenter func(*MSPointerEvent))
		

		// SetOnmspointerenter prop 
		// js:"onmspointerenter"
		SetOnmspointerenter (onmspointerenter func(*MSPointerEvent))
		
		// Onmspointerleave prop 
		// js:"onmspointerleave"
		Onmspointerleave() (onmspointerleave func(*MSPointerEvent))
		

		// SetOnmspointerleave prop 
		// js:"onmspointerleave"
		SetOnmspointerleave (onmspointerleave func(*MSPointerEvent))
		
		// Onmspointermove prop 
		// js:"onmspointermove"
		Onmspointermove() (onmspointermove func(*MSPointerEvent))
		

		// SetOnmspointermove prop 
		// js:"onmspointermove"
		SetOnmspointermove (onmspointermove func(*MSPointerEvent))
		
		// Onmspointerout prop 
		// js:"onmspointerout"
		Onmspointerout() (onmspointerout func(*MSPointerEvent))
		

		// SetOnmspointerout prop 
		// js:"onmspointerout"
		SetOnmspointerout (onmspointerout func(*MSPointerEvent))
		
		// Onmspointerover prop 
		// js:"onmspointerover"
		Onmspointerover() (onmspointerover func(*MSPointerEvent))
		

		// SetOnmspointerover prop 
		// js:"onmspointerover"
		SetOnmspointerover (onmspointerover func(*MSPointerEvent))
		
		// Onmspointerup prop 
		// js:"onmspointerup"
		Onmspointerup() (onmspointerup func(*MSPointerEvent))
		

		// SetOnmspointerup prop 
		// js:"onmspointerup"
		SetOnmspointerup (onmspointerup func(*MSPointerEvent))
		
		// Ontouchcancel prop 
		// js:"ontouchcancel"
		Ontouchcancel() (ontouchcancel func(*TouchEvent))
		

		// SetOntouchcancel prop 
		// js:"ontouchcancel"
		SetOntouchcancel (ontouchcancel func(*TouchEvent))
		
		// Ontouchend prop 
		// js:"ontouchend"
		Ontouchend() (ontouchend func(*TouchEvent))
		

		// SetOntouchend prop 
		// js:"ontouchend"
		SetOntouchend (ontouchend func(*TouchEvent))
		
		// Ontouchmove prop 
		// js:"ontouchmove"
		Ontouchmove() (ontouchmove func(*TouchEvent))
		

		// SetOntouchmove prop 
		// js:"ontouchmove"
		SetOntouchmove (ontouchmove func(*TouchEvent))
		
		// Ontouchstart prop 
		// js:"ontouchstart"
		Ontouchstart() (ontouchstart func(*TouchEvent))
		

		// SetOntouchstart prop 
		// js:"ontouchstart"
		SetOntouchstart (ontouchstart func(*TouchEvent))
		
		// Onwebkitfullscreenchange prop 
		// js:"onwebkitfullscreenchange"
		Onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event))
		

		// SetOnwebkitfullscreenchange prop 
		// js:"onwebkitfullscreenchange"
		SetOnwebkitfullscreenchange (onwebkitfullscreenchange func(Event))
		
		// Onwebkitfullscreenerror prop 
		// js:"onwebkitfullscreenerror"
		Onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event))
		

		// SetOnwebkitfullscreenerror prop 
		// js:"onwebkitfullscreenerror"
		SetOnwebkitfullscreenerror (onwebkitfullscreenerror func(Event))
		
		// OuterHTML prop 
		// js:"outerHTML"
		OuterHTML() (outerHTML string)
		

		// SetOuterHTML prop 
		// js:"outerHTML"
		SetOuterHTML (outerHTML string)
		
		// Prefix prop 
		// js:"prefix"
		Prefix() (prefix string)
		
		// ScrollHeight prop 
		// js:"scrollHeight"
		ScrollHeight() (scrollHeight int)
		
		// ScrollLeft prop 
		// js:"scrollLeft"
		ScrollLeft() (scrollLeft int)
		

		// SetScrollLeft prop 
		// js:"scrollLeft"
		SetScrollLeft (scrollLeft int)
		
		// ScrollTop prop 
		// js:"scrollTop"
		ScrollTop() (scrollTop int)
		

		// SetScrollTop prop 
		// js:"scrollTop"
		SetScrollTop (scrollTop int)
		
		// ScrollWidth prop 
		// js:"scrollWidth"
		ScrollWidth() (scrollWidth int)
		
		// TagName prop 
		// js:"tagName"
		TagName() (tagName string)
		
		}
	