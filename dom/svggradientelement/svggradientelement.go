package svggradientelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/svganimatedtransformlist"
	"github.com/matthewmueller/golly/dom/window"
)

// js:"SVGGradientElement,omit"
type SVGGradientElement interface {
	window.SVGElement

	// Href prop
	// js:"href"
	Href() (href *svganimatedstring.SVGAnimatedString)

	// GradientTransform prop
	// js:"gradientTransform"
	GradientTransform() (gradientTransform *svganimatedtransformlist.SVGAnimatedTransformList)

	// GradientUnits prop
	// js:"gradientUnits"
	GradientUnits() (gradientUnits *svganimatedenumeration.SVGAnimatedEnumeration)

	// SpreadMethod prop
	// js:"spreadMethod"
	SpreadMethod() (spreadMethod *svganimatedenumeration.SVGAnimatedEnumeration)
}
