package svgfittoviewbox

import (
	"github.com/matthewmueller/golly/dom/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/dom/svganimatedrect"
)

// SVGFitToViewBox interface
// js:"SVGFitToViewBox"
type SVGFitToViewBox interface {

	// PreserveAspectRatio prop
	// js:"preserveAspectRatio"
	// jsrewrite:"$_.preserveAspectRatio"
	PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio)

	// ViewBox prop
	// js:"viewBox"
	// jsrewrite:"$_.viewBox"
	ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect)
}
