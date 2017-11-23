package svgfegaussianblurelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEGaussianBlurElement struct
// js:"SVGFEGaussianBlurElement,omit"
type SVGFEGaussianBlurElement struct {
	window.SVGElement
}

// SetStdDeviation
func (*SVGFEGaussianBlurElement) SetStdDeviation(stdDeviationX float32, stdDeviationY float32) {
	js.Rewrite("$<.SetStdDeviation($1, $2)", stdDeviationX, stdDeviationY)
}

// Height
func (*SVGFEGaussianBlurElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFEGaussianBlurElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFEGaussianBlurElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFEGaussianBlurElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFEGaussianBlurElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}

// In1
func (*SVGFEGaussianBlurElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}

// StdDeviationX
func (*SVGFEGaussianBlurElement) StdDeviationX() (stdDeviationX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.StdDeviationX")
	return stdDeviationX
}

// StdDeviationY
func (*SVGFEGaussianBlurElement) StdDeviationY() (stdDeviationY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.StdDeviationY")
	return stdDeviationY
}
