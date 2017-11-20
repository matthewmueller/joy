package svgfegaussianblurelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumber"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedstring"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/js"
)

type SVGFEGaussianBlurElement struct {
	*svgelement.SVGElement
	*svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEGaussianBlurElement) SetStdDeviation(stdDeviationX float32, stdDeviationY float32) {
	js.Rewrite("$<.setStdDeviation($1, $2)", stdDeviationX, stdDeviationY)
}

func (*SVGFEGaussianBlurElement) GetIn1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFEGaussianBlurElement) GetStdDeviationX() (stdDeviationX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.stdDeviationX")
	return stdDeviationX
}

func (*SVGFEGaussianBlurElement) GetStdDeviationY() (stdDeviationY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.stdDeviationY")
	return stdDeviationY
}
