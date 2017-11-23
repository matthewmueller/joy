package svglineargradientelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svggradientelement"
	"github.com/matthewmueller/golly/js"
)

// SVGLinearGradientElement struct
// js:"SVGLinearGradientElement,omit"
type SVGLinearGradientElement struct {
	svggradientelement.SVGGradientElement
}

// X1
func (*SVGLinearGradientElement) X1() (x1 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X1")
	return x1
}

// X2
func (*SVGLinearGradientElement) X2() (x2 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X2")
	return x2
}

// Y1
func (*SVGLinearGradientElement) Y1() (y1 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y1")
	return y1
}

// Y2
func (*SVGLinearGradientElement) Y2() (y2 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y2")
	return y2
}
