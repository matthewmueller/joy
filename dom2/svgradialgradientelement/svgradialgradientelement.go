package svgradialgradientelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svggradientelement"
	"github.com/matthewmueller/golly/js"
)

// SVGRadialGradientElement struct
// js:"SVGRadialGradientElement,omit"
type SVGRadialGradientElement struct {
	svggradientelement.SVGGradientElement
}

// Cx
func (*SVGRadialGradientElement) Cx() (cx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Cx")
	return cx
}

// Cy
func (*SVGRadialGradientElement) Cy() (cy *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Cy")
	return cy
}

// Fx
func (*SVGRadialGradientElement) Fx() (fx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Fx")
	return fx
}

// Fy
func (*SVGRadialGradientElement) Fy() (fy *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Fy")
	return fy
}

// R
func (*SVGRadialGradientElement) R() (r *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.R")
	return r
}
