package svganimatedpoints

import "github.com/matthewmueller/golly/dom/svgpointlist"

// SVGAnimatedPoints interface
// js:"SVGAnimatedPoints"
type SVGAnimatedPoints interface {

	// AnimatedPoints prop
	// js:"animatedPoints"
	AnimatedPoints() (animatedPoints *svgpointlist.SVGPointList)

	// Points prop
	// js:"points"
	Points() (points *svgpointlist.SVGPointList)
}
