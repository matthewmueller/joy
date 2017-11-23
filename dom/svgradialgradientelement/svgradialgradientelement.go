package svgradialgradientelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svggradientelement"
	"github.com/matthewmueller/golly/js"
)

// SVGRadialGradientElement struct
// js:"SVGRadialGradientElement,omit"
type SVGRadialGradientElement struct {
	svggradientelement.SVGGradientElement
}

// Cx prop
func (*SVGRadialGradientElement) Cx() (cx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.cx")
	return cx
}

// Cy prop
func (*SVGRadialGradientElement) Cy() (cy *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.cy")
	return cy
}

// Fx prop
func (*SVGRadialGradientElement) Fx() (fx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.fx")
	return fx
}

// Fy prop
func (*SVGRadialGradientElement) Fy() (fy *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.fy")
	return fy
}

// R prop
func (*SVGRadialGradientElement) R() (r *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.r")
	return r
}
