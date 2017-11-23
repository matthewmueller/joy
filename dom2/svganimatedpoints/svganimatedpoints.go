package svganimatedpoints

import "github.com/matthewmueller/golly/dom2/svgpointlist"

// js:"SVGAnimatedPoints,omit"
type SVGAnimatedPoints interface {

	// AnimatedPoints
	AnimatedPoints() (animatedPoints *svgpointlist.SVGPointList)

	// Points
	Points() (points *svgpointlist.SVGPointList)
}
