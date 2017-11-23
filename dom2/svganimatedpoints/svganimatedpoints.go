package svganimatedpoints

// js:"SVGAnimatedPoints,omit"
type SVGAnimatedPoints interface {

	// AnimatedPoints
	AnimatedPoints() (animatedPoints *svgpointlist.SVGPointList)

	// Points
	Points() (points *svgpointlist.SVGPointList)
}
