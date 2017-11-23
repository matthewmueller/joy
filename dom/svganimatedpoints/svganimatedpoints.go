package svganimatedpoints

import (
	"github.com/matthewmueller/golly/dom/svgpointlist"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedPoints struct
// js:"SVGAnimatedPoints,omit"
type SVGAnimatedPoints struct {
}

// AnimatedPoints prop
func (*SVGAnimatedPoints) AnimatedPoints() (animatedPoints *svgpointlist.SVGPointList) {
	js.Rewrite("$<.animatedPoints")
	return animatedPoints
}

// Points prop
func (*SVGAnimatedPoints) Points() (points *svgpointlist.SVGPointList) {
	js.Rewrite("$<.points")
	return points
}
