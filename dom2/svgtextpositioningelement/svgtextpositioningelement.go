package svgtextpositioningelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlengthlist"
	"github.com/matthewmueller/golly/dom2/svganimatednumberlist"
	"github.com/matthewmueller/golly/dom2/svgtextcontentelement"
)

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
