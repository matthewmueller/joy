package svganimatedpreserveaspectratio

import (
	"github.com/matthewmueller/golly/dom2/svgpreserveaspectratio"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedPreserveAspectRatio struct
// js:"SVGAnimatedPreserveAspectRatio,omit"
type SVGAnimatedPreserveAspectRatio struct {
}

// AnimVal prop
func (*SVGAnimatedPreserveAspectRatio) AnimVal() (animVal *svgpreserveaspectratio.SVGPreserveAspectRatio) {
	js.Rewrite("$<.animVal")
	return animVal
}

// BaseVal prop
func (*SVGAnimatedPreserveAspectRatio) BaseVal() (baseVal *svgpreserveaspectratio.SVGPreserveAspectRatio) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
