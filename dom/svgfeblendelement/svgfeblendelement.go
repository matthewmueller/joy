package svgfeblendelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEBlendElement struct
// js:"SVGFEBlendElement,omit"
type SVGFEBlendElement struct {
	window.SVGElement
}

// Height prop
func (*SVGFEBlendElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Result prop
func (*SVGFEBlendElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

// Width prop
func (*SVGFEBlendElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGFEBlendElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFEBlendElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

// In1 prop
func (*SVGFEBlendElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

// In2 prop
func (*SVGFEBlendElement) In2() (in2 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in2")
	return in2
}

// Mode prop
func (*SVGFEBlendElement) Mode() (mode *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.mode")
	return mode
}
