package svgpolylineelement

import "github.com/matthewmueller/golly/dom2/window"

// js:"SVGPolylineElement,omit"
type SVGPolylineElement struct {
	window.SVGGraphicsElement
	svganimatedpoints.SVGAnimatedPoints
}
