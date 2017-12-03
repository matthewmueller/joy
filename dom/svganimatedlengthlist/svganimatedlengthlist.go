package svganimatedlengthlist

import (
	"github.com/matthewmueller/joy/dom/svglengthlist"
	"github.com/matthewmueller/joy/macro"
)

// SVGAnimatedLengthList struct
// js:"SVGAnimatedLengthList,omit"
type SVGAnimatedLengthList struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedLengthList) AnimVal() (animVal *svglengthlist.SVGLengthList) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedLengthList) BaseVal() (baseVal *svglengthlist.SVGLengthList) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}
