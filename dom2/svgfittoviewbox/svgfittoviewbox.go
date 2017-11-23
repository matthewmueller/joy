package svgfittoviewbox

import (
	"github.com/matthewmueller/golly/dom2/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/dom2/svganimatedrect"
	"github.com/matthewmueller/golly/js"
)

// SVGFitToViewBox struct
// js:"SVGFitToViewBox,omit"
type SVGFitToViewBox struct {
}

// PreserveAspectRatio
func (*SVGFitToViewBox) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.PreserveAspectRatio")
	return preserveAspectRatio
}

// ViewBox
func (*SVGFitToViewBox) ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect) {
	js.Rewrite("$<.ViewBox")
	return viewBox
}
