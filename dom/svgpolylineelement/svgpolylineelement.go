package svgpolylineelement

import (
	"github.com/matthewmueller/golly/dom/svgpointlist"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGPolylineElement struct
// js:"SVGPolylineElement,omit"
type SVGPolylineElement struct {
	window.SVGGraphicsElement
}

// AnimatedPoints prop
func (*SVGPolylineElement) AnimatedPoints() (animatedPoints *svgpointlist.SVGPointList) {
	js.Rewrite("$<.animatedPoints")
	return animatedPoints
}

// Points prop
func (*SVGPolylineElement) Points() (points *svgpointlist.SVGPointList) {
	js.Rewrite("$<.points")
	return points
}
