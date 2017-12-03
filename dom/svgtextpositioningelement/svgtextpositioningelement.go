package svgtextpositioningelement

import (
	"github.com/matthewmueller/joy/dom/svganimatedlengthlist"
	"github.com/matthewmueller/joy/dom/svganimatednumberlist"
	"github.com/matthewmueller/joy/dom/svgtextcontentelement"
)

// SVGTextPositioningElement interface
// js:"SVGTextPositioningElement"
type SVGTextPositioningElement interface {
	svgtextcontentelement.SVGTextContentElement

	// Dx prop
	// js:"dx"
	// jsrewrite:"$_.dx"
	Dx() (dx *svganimatedlengthlist.SVGAnimatedLengthList)

	// Dy prop
	// js:"dy"
	// jsrewrite:"$_.dy"
	Dy() (dy *svganimatedlengthlist.SVGAnimatedLengthList)

	// Rotate prop
	// js:"rotate"
	// jsrewrite:"$_.rotate"
	Rotate() (rotate *svganimatednumberlist.SVGAnimatedNumberList)

	// X prop
	// js:"x"
	// jsrewrite:"$_.x"
	X() (x *svganimatedlengthlist.SVGAnimatedLengthList)

	// Y prop
	// js:"y"
	// jsrewrite:"$_.y"
	Y() (y *svganimatedlengthlist.SVGAnimatedLengthList)
}
