package svglineargradientelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svggradientelement"
	"github.com/matthewmueller/golly/js"
)

// SVGLinearGradientElement struct
// js:"SVGLinearGradientElement,omit"
type SVGLinearGradientElement struct {
	svggradientelement.SVGGradientElement
}

// X1 prop
func (*SVGLinearGradientElement) X1() (x1 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x1")
	return x1
}

// X2 prop
func (*SVGLinearGradientElement) X2() (x2 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x2")
	return x2
}

// Y1 prop
func (*SVGLinearGradientElement) Y1() (y1 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y1")
	return y1
}

// Y2 prop
func (*SVGLinearGradientElement) Y2() (y2 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y2")
	return y2
}
