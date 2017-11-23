package svgcircleelement

import "github.com/matthewmueller/golly/js"

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
