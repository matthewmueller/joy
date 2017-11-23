package svgellipseelement

import "github.com/matthewmueller/golly/js"

// js:"SVGEllipseElement,omit"
type SVGEllipseElement struct {
	window.SVGGraphicsElement
}

// Cx
func (*SVGEllipseElement) Cx() (cx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Cx")
	return cx
}

// Cy
func (*SVGEllipseElement) Cy() (cy *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Cy")
	return cy
}

// Rx
func (*SVGEllipseElement) Rx() (rx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Rx")
	return rx
}

// Ry
func (*SVGEllipseElement) Ry() (ry *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Ry")
	return ry
}
