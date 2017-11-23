package svgellipseelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGEllipseElement struct
// js:"SVGEllipseElement,omit"
type SVGEllipseElement struct {
	window.SVGGraphicsElement
}

// Cx prop
func (*SVGEllipseElement) Cx() (cx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.cx")
	return cx
}

// Cy prop
func (*SVGEllipseElement) Cy() (cy *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.cy")
	return cy
}

// Rx prop
func (*SVGEllipseElement) Rx() (rx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.rx")
	return rx
}

// Ry prop
func (*SVGEllipseElement) Ry() (ry *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.ry")
	return ry
}
