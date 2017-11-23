package svganimatedtransformlist

import (
	"github.com/matthewmueller/golly/dom2/svgtransformlist"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGAnimatedTransformList,omit"
type SVGAnimatedTransformList struct {
}

// AnimVal
func (*SVGAnimatedTransformList) AnimVal() (animVal *svgtransformlist.SVGTransformList) {
	js.Rewrite("$<.AnimVal")
	return animVal
}

// BaseVal
func (*SVGAnimatedTransformList) BaseVal() (baseVal *svgtransformlist.SVGTransformList) {
	js.Rewrite("$<.BaseVal")
	return baseVal
}
