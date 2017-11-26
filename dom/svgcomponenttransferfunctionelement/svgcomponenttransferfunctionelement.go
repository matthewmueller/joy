package svgcomponenttransferfunctionelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatednumber"
	"github.com/matthewmueller/golly/dom/svganimatednumberlist"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGComponentTransferFunctionElement,omit"
type SVGComponentTransferFunctionElement interface {
	window.SVGElement

	// Amplitude prop
	// js:"amplitude,rewrite=amplitude"
	Amplitude() (amplitude *svganimatednumber.SVGAnimatedNumber)

	// Exponent prop
	// js:"exponent,rewrite=exponent"
	Exponent() (exponent *svganimatednumber.SVGAnimatedNumber)

	// Intercept prop
	// js:"intercept,rewrite=intercept"
	Intercept() (intercept *svganimatednumber.SVGAnimatedNumber)

	// Offset prop
	// js:"offset,rewrite=offset"
	Offset() (offset *svganimatednumber.SVGAnimatedNumber)

	// Slope prop
	// js:"slope,rewrite=slope"
	Slope() (slope *svganimatednumber.SVGAnimatedNumber)

	// TableValues prop
	// js:"tableValues,rewrite=tablevalues"
	TableValues() (tableValues *svganimatednumberlist.SVGAnimatedNumberList)

	// Type prop
	// js:"type,rewrite=kind"
	Type() (kind *svganimatedenumeration.SVGAnimatedEnumeration)
}

// amplitude prop
func amplitude() (amplitude *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.amplitude")
	return amplitude
}

// exponent prop
func exponent() (exponent *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.exponent")
	return exponent
}

// intercept prop
func intercept() (intercept *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.intercept")
	return intercept
}

// offset prop
func offset() (offset *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.offset")
	return offset
}

// slope prop
func slope() (slope *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.slope")
	return slope
}

// tablevalues prop
func tablevalues() (tableValues *svganimatednumberlist.SVGAnimatedNumberList) {
	js.Rewrite("$<.tableValues")
	return tableValues
}

// kind prop
func kind() (kind *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.type")
	return kind
}
