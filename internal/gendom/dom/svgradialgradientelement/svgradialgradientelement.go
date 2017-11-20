package svgradialgradientelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggradientelement"
	"github.com/matthewmueller/golly/js"
)

type SVGRadialGradientElement struct {
	*svggradientelement.SVGGradientElement
}

func (*SVGRadialGradientElement) GetCx() (cx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.cx")
	return cx
}

func (*SVGRadialGradientElement) GetCy() (cy *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.cy")
	return cy
}

func (*SVGRadialGradientElement) GetFx() (fx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.fx")
	return fx
}

func (*SVGRadialGradientElement) GetFy() (fy *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.fy")
	return fy
}

func (*SVGRadialGradientElement) GetR() (r *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.r")
	return r
}
