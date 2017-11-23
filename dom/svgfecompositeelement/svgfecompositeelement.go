package svgfecompositeelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svganimatednumber"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFECompositeElement struct
// js:"SVGFECompositeElement,omit"
type SVGFECompositeElement struct {
	window.SVGElement
}

// Height prop
func (*SVGFECompositeElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Result prop
func (*SVGFECompositeElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

// Width prop
func (*SVGFECompositeElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGFECompositeElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFECompositeElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

// In1 prop
func (*SVGFECompositeElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

// In2 prop
func (*SVGFECompositeElement) In2() (in2 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in2")
	return in2
}

// K1 prop
func (*SVGFECompositeElement) K1() (k1 *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.k1")
	return k1
}

// K2 prop
func (*SVGFECompositeElement) K2() (k2 *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.k2")
	return k2
}

// K3 prop
func (*SVGFECompositeElement) K3() (k3 *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.k3")
	return k3
}

// K4 prop
func (*SVGFECompositeElement) K4() (k4 *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.k4")
	return k4
}

// Operator prop
func (*SVGFECompositeElement) Operator() (operator *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.operator")
	return operator
}
