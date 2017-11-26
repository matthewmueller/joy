package svgtextpositioningelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedlengthlist"
	"github.com/matthewmueller/golly/dom/svganimatednumberlist"
	"github.com/matthewmueller/golly/dom/svgtextcontentelement"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGTextPositioningElement,omit"
type SVGTextPositioningElement interface {
	svgtextcontentelement.SVGTextContentElement

	// Dx prop
	// js:"dx,rewrite=dx"
	Dx() (dx *svganimatedlengthlist.SVGAnimatedLengthList)

	// Dy prop
	// js:"dy,rewrite=dy"
	Dy() (dy *svganimatedlengthlist.SVGAnimatedLengthList)

	// Rotate prop
	// js:"rotate,rewrite=rotate"
	Rotate() (rotate *svganimatednumberlist.SVGAnimatedNumberList)

	// X prop
	// js:"x,rewrite=x"
	X() (x *svganimatedlengthlist.SVGAnimatedLengthList)

	// Y prop
	// js:"y,rewrite=y"
	Y() (y *svganimatedlengthlist.SVGAnimatedLengthList)
}

// dx prop
func dx() (dx *svganimatedlengthlist.SVGAnimatedLengthList) {
	js.Rewrite("$<.dx")
	return dx
}

// dy prop
func dy() (dy *svganimatedlengthlist.SVGAnimatedLengthList) {
	js.Rewrite("$<.dy")
	return dy
}

// rotate prop
func rotate() (rotate *svganimatednumberlist.SVGAnimatedNumberList) {
	js.Rewrite("$<.rotate")
	return rotate
}

// x prop
func x() (x *svganimatedlengthlist.SVGAnimatedLengthList) {
	js.Rewrite("$<.x")
	return x
}

// y prop
func y() (y *svganimatedlengthlist.SVGAnimatedLengthList) {
	js.Rewrite("$<.y")
	return y
}
