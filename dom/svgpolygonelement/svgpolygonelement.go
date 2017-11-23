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

// AnimatedPoints prop
func (*SVGPolygonElement) AnimatedPoints() (animatedPoints *svgpointlist.SVGPointList) {
	js.Rewrite("$<.animatedPoints")
	return animatedPoints
}

// Points prop
func (*SVGPolygonElement) Points() (points *svgpointlist.SVGPointList) {
	js.Rewrite("$<.points")
	return points
}
