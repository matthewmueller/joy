package svggradientelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedtransformlist"
	"github.com/matthewmueller/golly/dom2/svgunittypes"
	"github.com/matthewmueller/golly/dom2/svgurireference"
	"github.com/matthewmueller/golly/dom2/window"
)

// js:"SVGGradientElement,omit"
type SVGGradientElement interface {
	window.SVGElement
	svgunittypes.SVGUnitTypes
	svgurireference.SVGURIReference

	// GradientTransform
	GradientTransform() (gradientTransform *svganimatedtransformlist.SVGAnimatedTransformList)

	// GradientUnits
	GradientUnits() (gradientUnits *svganimatedenumeration.SVGAnimatedEnumeration)

	// SpreadMethod
	SpreadMethod() (spreadMethod *svganimatedenumeration.SVGAnimatedEnumeration)
}
