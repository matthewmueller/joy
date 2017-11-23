package svgfegaussianblurelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGFEGaussianBlurElement,omit"
type SVGFEGaussianBlurElement struct {
	window.SVGElement
	svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

// SetStdDeviation
func (*SVGFEGaussianBlurElement) SetStdDeviation(stdDeviationX float32, stdDeviationY float32) {
	js.Rewrite("$<.SetStdDeviation($1, $2)", stdDeviationX, stdDeviationY)
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
