package svgfittoviewbox

import (
	"github.com/matthewmueller/golly/dom/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/dom/svganimatedrect"
	"github.com/matthewmueller/golly/js"
)

// SVGFitToViewBox struct
// js:"SVGFitToViewBox,omit"
type SVGFitToViewBox struct {
}

// PreserveAspectRatio prop
func (*SVGFitToViewBox) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.preserveAspectRatio")
	return preserveAspectRatio
}

// ViewBox prop
func (*SVGFitToViewBox) ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect) {
	js.Rewrite("$<.viewBox")
	return viewBox
}
