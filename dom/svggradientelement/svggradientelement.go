package svggradientelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/svganimatedtransformlist"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGGradientElement,omit"
type SVGGradientElement interface {
	window.SVGElement

	// Href prop
	// js:"href,rewrite=href"
	Href() (href *svganimatedstring.SVGAnimatedString)

	// GradientTransform prop
	// js:"gradientTransform,rewrite=gradienttransform"
	GradientTransform() (gradientTransform *svganimatedtransformlist.SVGAnimatedTransformList)

	// GradientUnits prop
	// js:"gradientUnits,rewrite=gradientunits"
	GradientUnits() (gradientUnits *svganimatedenumeration.SVGAnimatedEnumeration)

	// SpreadMethod prop
	// js:"spreadMethod,rewrite=spreadmethod"
	SpreadMethod() (spreadMethod *svganimatedenumeration.SVGAnimatedEnumeration)
}

// gradienttransform prop
func gradienttransform() (gradientTransform *svganimatedtransformlist.SVGAnimatedTransformList) {
	js.Rewrite("$<.gradientTransform")
	return gradientTransform
}

// gradientunits prop
func gradientunits() (gradientUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.gradientUnits")
	return gradientUnits
}

// spreadmethod prop
func spreadmethod() (spreadMethod *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.spreadMethod")
	return spreadMethod
}
