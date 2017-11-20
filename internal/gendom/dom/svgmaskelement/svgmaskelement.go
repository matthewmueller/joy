package svgmaskelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgtests"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgunittypes"
	"github.com/matthewmueller/golly/js"
)

type SVGMaskElement struct {
	*svgelement.SVGElement
	*svgtests.SVGTests
	*svgunittypes.SVGUnitTypes
}

func (*SVGMaskElement) GetHeight() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGMaskElement) GetMaskContentUnits() (maskContentUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.maskContentUnits")
	return maskContentUnits
}

func (*SVGMaskElement) GetMaskUnits() (maskUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.maskUnits")
	return maskUnits
}

func (*SVGMaskElement) GetWidth() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGMaskElement) GetX() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGMaskElement) GetY() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
