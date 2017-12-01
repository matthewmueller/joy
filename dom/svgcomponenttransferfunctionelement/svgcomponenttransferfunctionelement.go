package svgcomponenttransferfunctionelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatednumber"
	"github.com/matthewmueller/golly/dom/svganimatednumberlist"
	"github.com/matthewmueller/golly/dom/window"
)

// SVGComponentTransferFunctionElement interface
// js:"SVGComponentTransferFunctionElement"
type SVGComponentTransferFunctionElement interface {
	window.SVGElement

	// Amplitude prop
	// js:"amplitude"
	// jsrewrite:"$_.amplitude"
	Amplitude() (amplitude *svganimatednumber.SVGAnimatedNumber)

	// Exponent prop
	// js:"exponent"
	// jsrewrite:"$_.exponent"
	Exponent() (exponent *svganimatednumber.SVGAnimatedNumber)

	// Intercept prop
	// js:"intercept"
	// jsrewrite:"$_.intercept"
	Intercept() (intercept *svganimatednumber.SVGAnimatedNumber)

	// Offset prop
	// js:"offset"
	// jsrewrite:"$_.offset"
	Offset() (offset *svganimatednumber.SVGAnimatedNumber)

	// Slope prop
	// js:"slope"
	// jsrewrite:"$_.slope"
	Slope() (slope *svganimatednumber.SVGAnimatedNumber)

	// TableValues prop
	// js:"tableValues"
	// jsrewrite:"$_.tableValues"
	TableValues() (tableValues *svganimatednumberlist.SVGAnimatedNumberList)

	// Type prop
	// js:"type"
	// jsrewrite:"$_.type"
	Type() (kind *svganimatedenumeration.SVGAnimatedEnumeration)
}
