package svganimatedtransformlist

import (
	"github.com/matthewmueller/joy/dom/svgtransformlist"
	"github.com/matthewmueller/joy/js"
)

// SVGAnimatedTransformList struct
// js:"SVGAnimatedTransformList,omit"
type SVGAnimatedTransformList struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedTransformList) AnimVal() (animVal *svgtransformlist.SVGTransformList) {
	js.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedTransformList) BaseVal() (baseVal *svgtransformlist.SVGTransformList) {
	js.Rewrite("$_.baseVal")
	return baseVal
}
