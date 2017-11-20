package svgpatternelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedtransformlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfittoviewbox"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgtests"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgunittypes"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgurireference"
	"github.com/matthewmueller/golly/js"
)

type SVGPatternElement struct {
	*svgelement.SVGElement
	*svgtests.SVGTests
	*svgunittypes.SVGUnitTypes
	*svgfittoviewbox.SVGFitToViewBox
	*svgurireference.SVGURIReference
}

func (*SVGPatternElement) GetHeight() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGPatternElement) GetPatternContentUnits() (patternContentUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.patternContentUnits")
	return patternContentUnits
}

func (*SVGPatternElement) GetPatternTransform() (patternTransform *svganimatedtransformlist.SVGAnimatedTransformList) {
	js.Rewrite("$<.patternTransform")
	return patternTransform
}

func (*SVGPatternElement) GetPatternUnits() (patternUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.patternUnits")
	return patternUnits
}

func (*SVGPatternElement) GetWidth() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGPatternElement) GetX() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPatternElement) GetY() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
