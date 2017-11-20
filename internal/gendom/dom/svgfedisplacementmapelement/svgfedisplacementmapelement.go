package svgfedisplacementmapelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumber"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedstring"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/js"
)

type SVGFEDisplacementMapElement struct {
	*svgelement.SVGElement
	*svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEDisplacementMapElement) GetIn1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFEDisplacementMapElement) GetIn2() (in2 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in2")
	return in2
}

func (*SVGFEDisplacementMapElement) GetScale() (scale *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.scale")
	return scale
}

func (*SVGFEDisplacementMapElement) GetXChannelSelector() (xChannelSelector *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.xChannelSelector")
	return xChannelSelector
}

func (*SVGFEDisplacementMapElement) GetYChannelSelector() (yChannelSelector *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.yChannelSelector")
	return yChannelSelector
}
