package window

// js:"SVGElement,omit"
type SVGElement interface {
	Element

	// ClassName
	ClassName() (className *svganimatedstring.SVGAnimatedString)

	// Onclick
	Onclick() (onclick func(MouseEvent))

	// Onclick
	SetOnclick(onclick func(MouseEvent))

	// Ondblclick
	Ondblclick() (ondblclick func(MouseEvent))

	// Ondblclick
	SetOndblclick(ondblclick func(MouseEvent))

	// Onfocusin
	Onfocusin() (onfocusin func(*FocusEvent))

	// Onfocusin
	SetOnfocusin(onfocusin func(*FocusEvent))

	// Onfocusout
	Onfocusout() (onfocusout func(*FocusEvent))

	// Onfocusout
	SetOnfocusout(onfocusout func(*FocusEvent))

	// Onload
	Onload() (onload func(Event))

	// Onload
	SetOnload(onload func(Event))

	// Onmousedown
	Onmousedown() (onmousedown func(MouseEvent))

	// Onmousedown
	SetOnmousedown(onmousedown func(MouseEvent))

	// Onmousemove
	Onmousemove() (onmousemove func(MouseEvent))

	// Onmousemove
	SetOnmousemove(onmousemove func(MouseEvent))

	// Onmouseout
	Onmouseout() (onmouseout func(MouseEvent))

	// Onmouseout
	SetOnmouseout(onmouseout func(MouseEvent))

	// Onmouseover
	Onmouseover() (onmouseover func(MouseEvent))

	// Onmouseover
	SetOnmouseover(onmouseover func(MouseEvent))

	// Onmouseup
	Onmouseup() (onmouseup func(MouseEvent))

	// Onmouseup
	SetOnmouseup(onmouseup func(MouseEvent))

	// OwnerSVGElement
	OwnerSVGElement() (ownerSVGElement *SVGSVGElement)

	// Style
	Style() (style *CSSStyleDeclaration)

	// ViewportElement
	ViewportElement() (viewportElement SVGElement)

	// Xmlbase
	Xmlbase() (xmlbase string)

	// Xmlbase
	SetXmlbase(xmlbase string)
}
