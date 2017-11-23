package svgtextpositioningelement

// js:"SVGTextPositioningElement,omit"
type SVGTextPositioningElement interface {
	svgtextcontentelement.SVGTextContentElement

	// Dx
	Dx() (dx *svganimatedlengthlist.SVGAnimatedLengthList)

	// Dy
	Dy() (dy *svganimatedlengthlist.SVGAnimatedLengthList)

	// Rotate
	Rotate() (rotate *svganimatednumberlist.SVGAnimatedNumberList)

	// X
	X() (x *svganimatedlengthlist.SVGAnimatedLengthList)

	// Y
	Y() (y *svganimatedlengthlist.SVGAnimatedLengthList)
}
