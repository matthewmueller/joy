package svganimatedtransformlist

import (
	"github.com/matthewmueller/golly/dom2/svgtransformlist"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedTransformList struct
// js:"SVGAnimatedTransformList,omit"
type SVGAnimatedTransformList struct {
}

// AnimVal prop
func (*SVGAnimatedTransformList) AnimVal() (animVal *svgtransformlist.SVGTransformList) {
	js.Rewrite("$<.animVal")
	return animVal
}

// BaseVal prop
func (*SVGAnimatedTransformList) BaseVal() (baseVal *svgtransformlist.SVGTransformList) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
