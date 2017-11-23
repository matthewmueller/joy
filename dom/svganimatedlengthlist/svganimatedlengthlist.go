package svganimatedlengthlist

import (
	"github.com/matthewmueller/golly/dom/svglengthlist"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedLengthList struct
// js:"SVGAnimatedLengthList,omit"
type SVGAnimatedLengthList struct {
}

// AnimVal prop
func (*SVGAnimatedLengthList) AnimVal() (animVal *svglengthlist.SVGLengthList) {
	js.Rewrite("$<.animVal")
	return animVal
}

// BaseVal prop
func (*SVGAnimatedLengthList) BaseVal() (baseVal *svglengthlist.SVGLengthList) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
