package svgcircleelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGCircleElement struct
// js:"SVGCircleElement,omit"
type SVGCircleElement struct {
	window.SVGGraphicsElement
}

// Cx prop
func (*SVGCircleElement) Cx() (cx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.cx")
	return cx
}

// Cy prop
func (*SVGCircleElement) Cy() (cy *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.cy")
	return cy
}

// R prop
func (*SVGCircleElement) R() (r *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.r")
	return r
}
