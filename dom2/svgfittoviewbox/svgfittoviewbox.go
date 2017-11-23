package svgfittoviewbox

import (
	"github.com/matthewmueller/golly/dom2/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/dom2/svganimatedrect"
)

// js:"SVGFitToViewBox,omit"
type SVGFitToViewBox interface {

	// PreserveAspectRatio
	PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio)

	// ViewBox
	ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect)
}
