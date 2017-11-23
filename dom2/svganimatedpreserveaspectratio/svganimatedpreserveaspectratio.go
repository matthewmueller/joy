package svganimatedpreserveaspectratio

import (
	"github.com/matthewmueller/golly/dom2/svgpreserveaspectratio"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedPreserveAspectRatio struct
// js:"SVGAnimatedPreserveAspectRatio,omit"
type SVGAnimatedPreserveAspectRatio struct {
}

// AnimVal
func (*SVGAnimatedPreserveAspectRatio) AnimVal() (animVal *svgpreserveaspectratio.SVGPreserveAspectRatio) {
	js.Rewrite("$<.AnimVal")
	return animVal
}

// BaseVal
func (*SVGAnimatedPreserveAspectRatio) BaseVal() (baseVal *svgpreserveaspectratio.SVGPreserveAspectRatio) {
	js.Rewrite("$<.BaseVal")
	return baseVal
}
