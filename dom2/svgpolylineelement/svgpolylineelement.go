package svgpolylineelement

import (
	"github.com/matthewmueller/golly/dom2/svgpointlist"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGPolylineElement struct
// js:"SVGPolylineElement,omit"
type SVGPolylineElement struct {
	window.SVGGraphicsElement
}

// AnimatedPoints
func (*SVGPolylineElement) AnimatedPoints() (animatedPoints *svgpointlist.SVGPointList) {
	js.Rewrite("$<.AnimatedPoints")
	return animatedPoints
}

// Points
func (*SVGPolylineElement) Points() (points *svgpointlist.SVGPointList) {
	js.Rewrite("$<.Points")
	return points
}
