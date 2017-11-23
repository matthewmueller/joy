package svgtextpositioningelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedlengthlist"
	"github.com/matthewmueller/golly/dom/svganimatednumberlist"
	"github.com/matthewmueller/golly/dom/svgtextcontentelement"
)

// js:"SVGTextPositioningElement,omit"
type SVGTextPositioningElement interface {
	svgtextcontentelement.SVGTextContentElement

	// Dx prop
	// js:"dx"
	Dx() (dx *svganimatedlengthlist.SVGAnimatedLengthList)

	// Dy prop
	// js:"dy"
	Dy() (dy *svganimatedlengthlist.SVGAnimatedLengthList)

	// Rotate prop
	// js:"rotate"
	Rotate() (rotate *svganimatednumberlist.SVGAnimatedNumberList)

	// X prop
	// js:"x"
	X() (x *svganimatedlengthlist.SVGAnimatedLengthList)

	// Y prop
	// js:"y"
	Y() (y *svganimatedlengthlist.SVGAnimatedLengthList)
}
