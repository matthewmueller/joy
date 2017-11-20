package svggradientelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedtransformlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgunittypes"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgurireference"
	"github.com/matthewmueller/golly/js"
)

type SVGGradientElement struct {
	*svgelement.SVGElement
	*svgunittypes.SVGUnitTypes
	*svgurireference.SVGURIReference
}

func (*SVGGradientElement) GetGradientTransform() (gradientTransform *svganimatedtransformlist.SVGAnimatedTransformList) {
	js.Rewrite("$<.gradientTransform")
	return gradientTransform
}

func (*SVGGradientElement) GetGradientUnits() (gradientUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.gradientUnits")
	return gradientUnits
}

func (*SVGGradientElement) GetSpreadMethod() (spreadMethod *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.spreadMethod")
	return spreadMethod
}
