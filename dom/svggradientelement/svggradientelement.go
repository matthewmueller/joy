package svggradientelement

import (
	"github.com/matthewmueller/joy/dom/svganimatedenumeration"
	"github.com/matthewmueller/joy/dom/svganimatedstring"
	"github.com/matthewmueller/joy/dom/svganimatedtransformlist"
	"github.com/matthewmueller/joy/dom/window"
)

// SVGGradientElement interface
// js:"SVGGradientElement"
type SVGGradientElement interface {
	window.SVGElement

	// Href prop
	// js:"href"
	// jsrewrite:"$_.href"
	Href() (href *svganimatedstring.SVGAnimatedString)

	// GradientTransform prop
	// js:"gradientTransform"
	// jsrewrite:"$_.gradientTransform"
	GradientTransform() (gradientTransform *svganimatedtransformlist.SVGAnimatedTransformList)

	// GradientUnits prop
	// js:"gradientUnits"
	// jsrewrite:"$_.gradientUnits"
	GradientUnits() (gradientUnits *svganimatedenumeration.SVGAnimatedEnumeration)

	// SpreadMethod prop
	// js:"spreadMethod"
	// jsrewrite:"$_.spreadMethod"
	SpreadMethod() (spreadMethod *svganimatedenumeration.SVGAnimatedEnumeration)
}
