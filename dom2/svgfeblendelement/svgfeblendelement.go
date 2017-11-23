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

// Height
func (*SVGFEBlendElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFEBlendElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFEBlendElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFEBlendElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFEBlendElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}

// In1
func (*SVGFEBlendElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}

// In2
func (*SVGFEBlendElement) In2() (in2 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In2")
	return in2
}

// Mode
func (*SVGFEBlendElement) Mode() (mode *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.Mode")
	return mode
}
