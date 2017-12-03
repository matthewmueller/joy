package svganimatedpreserveaspectratio

import (
	"github.com/matthewmueller/joy/dom/svgpreserveaspectratio"
	"github.com/matthewmueller/joy/macro"
)

// SVGAnimatedPreserveAspectRatio struct
// js:"SVGAnimatedPreserveAspectRatio,omit"
type SVGAnimatedPreserveAspectRatio struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedPreserveAspectRatio) AnimVal() (animVal *svgpreserveaspectratio.SVGPreserveAspectRatio) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedPreserveAspectRatio) BaseVal() (baseVal *svgpreserveaspectratio.SVGPreserveAspectRatio) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}
