package svgfecompositeelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGFECompositeElement,omit"
type SVGFECompositeElement struct {
	window.SVGElement
	svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

// In1
func (*SVGFECompositeElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}

// In2
func (*SVGFECompositeElement) In2() (in2 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In2")
	return in2
}

// K1
func (*SVGFECompositeElement) K1() (k1 *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.K1")
	return k1
}

// K2
func (*SVGFECompositeElement) K2() (k2 *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.K2")
	return k2
}

// K3
func (*SVGFECompositeElement) K3() (k3 *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.K3")
	return k3
}

// K4
func (*SVGFECompositeElement) K4() (k4 *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.K4")
	return k4
}

// Operator
func (*SVGFECompositeElement) Operator() (operator *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.Operator")
	return operator
}
