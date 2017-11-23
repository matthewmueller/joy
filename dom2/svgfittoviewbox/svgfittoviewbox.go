package svgfittoviewbox

// js:"SVGFitToViewBox,omit"
type SVGFitToViewBox interface {

	// PreserveAspectRatio
	PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio)

	// ViewBox
	ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect)
}
