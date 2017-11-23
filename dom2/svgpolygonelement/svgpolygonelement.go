package svgpolygonelement

import (
	"github.com/matthewmueller/golly/dom2/svgpointlist"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGPolygonElement struct
// js:"SVGPolygonElement,omit"
type SVGPolygonElement struct {
	window.SVGGraphicsElement
}

// AnimatedPoints
func (*SVGPolygonElement) AnimatedPoints() (animatedPoints *svgpointlist.SVGPointList) {
	js.Rewrite("$<.AnimatedPoints")
	return animatedPoints
}

// Points
func (*SVGPolygonElement) Points() (points *svgpointlist.SVGPointList) {
	js.Rewrite("$<.Points")
	return points
}
