package svganimatedpoints

import "github.com/matthewmueller/joy/dom/svgpointlist"

// SVGAnimatedPoints interface
// js:"SVGAnimatedPoints"
type SVGAnimatedPoints interface {

	// AnimatedPoints prop
	// js:"animatedPoints"
	// jsrewrite:"$_.animatedPoints"
	AnimatedPoints() (animatedPoints *svgpointlist.SVGPointList)

	// Points prop
	// js:"points"
	// jsrewrite:"$_.points"
	Points() (points *svgpointlist.SVGPointList)
}
