package svgfecompositeelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumber"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedstring"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/js"
)

type SVGFECompositeElement struct {
	*svgelement.SVGElement
	*svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

func (*SVGFECompositeElement) GetIn1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFECompositeElement) GetIn2() (in2 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in2")
	return in2
}

func (*SVGFECompositeElement) GetK1() (k1 *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.k1")
	return k1
}

func (*SVGFECompositeElement) GetK2() (k2 *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.k2")
	return k2
}

func (*SVGFECompositeElement) GetK3() (k3 *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.k3")
	return k3
}

func (*SVGFECompositeElement) GetK4() (k4 *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.k4")
	return k4
}

func (*SVGFECompositeElement) GetOperator() (operator *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.operator")
	return operator
}
