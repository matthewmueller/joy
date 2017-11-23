package svgfecompositeelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFECompositeElement struct
// js:"SVGFECompositeElement,omit"
type SVGFECompositeElement struct {
	window.SVGElement
}

// Height
func (*SVGFECompositeElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFECompositeElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFECompositeElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFECompositeElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFECompositeElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
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
