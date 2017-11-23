package svgcomponenttransferfunctionelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svganimatednumberlist"
	"github.com/matthewmueller/golly/dom2/window"
)

// js:"SVGComponentTransferFunctionElement,omit"
type SVGComponentTransferFunctionElement interface {
	window.SVGElement

	// Amplitude
	Amplitude() (amplitude *svganimatednumber.SVGAnimatedNumber)

	// Exponent
	Exponent() (exponent *svganimatednumber.SVGAnimatedNumber)

	// Intercept
	Intercept() (intercept *svganimatednumber.SVGAnimatedNumber)

	// Offset
	Offset() (offset *svganimatednumber.SVGAnimatedNumber)

	// Slope
	Slope() (slope *svganimatednumber.SVGAnimatedNumber)

	// TableValues
	TableValues() (tableValues *svganimatednumberlist.SVGAnimatedNumberList)

	// Type
	Type() (kind *svganimatedenumeration.SVGAnimatedEnumeration)
}
