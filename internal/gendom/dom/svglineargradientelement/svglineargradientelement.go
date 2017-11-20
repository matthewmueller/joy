package svglineargradientelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggradientelement"
	"github.com/matthewmueller/golly/js"
)

type SVGLinearGradientElement struct {
	*svggradientelement.SVGGradientElement
}

func (*SVGLinearGradientElement) GetX1() (x1 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x1")
	return x1
}

func (*SVGLinearGradientElement) GetX2() (x2 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x2")
	return x2
}

func (*SVGLinearGradientElement) GetY1() (y1 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y1")
	return y1
}

func (*SVGLinearGradientElement) GetY2() (y2 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y2")
	return y2
}
