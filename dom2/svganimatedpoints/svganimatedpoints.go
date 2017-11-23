package svganimatedpoints

import (
	"github.com/matthewmueller/golly/dom2/svgpointlist"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedPoints struct
// js:"SVGAnimatedPoints,omit"
type SVGAnimatedPoints struct {
}

// AnimatedPoints
func (*SVGAnimatedPoints) AnimatedPoints() (animatedPoints *svgpointlist.SVGPointList) {
	js.Rewrite("$<.AnimatedPoints")
	return animatedPoints
}

// Points
func (*SVGAnimatedPoints) Points() (points *svgpointlist.SVGPointList) {
	js.Rewrite("$<.Points")
	return points
}
