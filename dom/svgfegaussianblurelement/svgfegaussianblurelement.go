package svgfegaussianblurelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svganimatednumber"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEGaussianBlurElement struct
// js:"SVGFEGaussianBlurElement,omit"
type SVGFEGaussianBlurElement struct {
	window.SVGElement
}

// SetStdDeviation fn
func (*SVGFEGaussianBlurElement) SetStdDeviation(stdDeviationX float32, stdDeviationY float32) {
	js.Rewrite("$<.setStdDeviation($1, $2)", stdDeviationX, stdDeviationY)
}

// Height prop
func (*SVGFEGaussianBlurElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Result prop
func (*SVGFEGaussianBlurElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

// Width prop
func (*SVGFEGaussianBlurElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGFEGaussianBlurElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFEGaussianBlurElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

// In1 prop
func (*SVGFEGaussianBlurElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

// StdDeviationX prop
func (*SVGFEGaussianBlurElement) StdDeviationX() (stdDeviationX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.stdDeviationX")
	return stdDeviationX
}

// StdDeviationY prop
func (*SVGFEGaussianBlurElement) StdDeviationY() (stdDeviationY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.stdDeviationY")
	return stdDeviationY
}
