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
// js:"animVal"
func (*SVGAnimatedLengthList) AnimVal() (animVal *svglengthlist.SVGLengthList) {
	js.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedLengthList) BaseVal() (baseVal *svglengthlist.SVGLengthList) {
	js.Rewrite("$_.baseVal")
	return baseVal
}
