package svgcircleelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGCircleElement struct
// js:"SVGCircleElement,omit"
type SVGCircleElement struct {
	window.SVGGraphicsElement
}

// Cx
func (*SVGCircleElement) Cx() (cx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Cx")
	return cx
}

// Cy
func (*SVGCircleElement) Cy() (cy *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Cy")
	return cy
}

// R
func (*SVGCircleElement) R() (r *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.R")
	return r
}
