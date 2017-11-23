package svgcomponenttransferfunctionelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatednumber"
	"github.com/matthewmueller/golly/dom/svganimatednumberlist"
	"github.com/matthewmueller/golly/dom/window"
)

// js:"SVGComponentTransferFunctionElement,omit"
type SVGComponentTransferFunctionElement interface {
	window.SVGElement

	// Amplitude prop
	// js:"amplitude"
	Amplitude() (amplitude *svganimatednumber.SVGAnimatedNumber)

	// Exponent prop
	// js:"exponent"
	Exponent() (exponent *svganimatednumber.SVGAnimatedNumber)

	// Intercept prop
	// js:"intercept"
	Intercept() (intercept *svganimatednumber.SVGAnimatedNumber)

	// Offset prop
	// js:"offset"
	Offset() (offset *svganimatednumber.SVGAnimatedNumber)

	// Slope prop
	// js:"slope"
	Slope() (slope *svganimatednumber.SVGAnimatedNumber)

	// TableValues prop
	// js:"tableValues"
	TableValues() (tableValues *svganimatednumberlist.SVGAnimatedNumberList)

	// Type prop
	// js:"type"
	Type() (kind *svganimatedenumeration.SVGAnimatedEnumeration)
}
