package svganimatedlengthlist

import (
	"github.com/matthewmueller/golly/dom2/svglengthlist"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGAnimatedLengthList,omit"
type SVGAnimatedLengthList struct {
}

// AnimVal
func (*SVGAnimatedLengthList) AnimVal() (animVal *svglengthlist.SVGLengthList) {
	js.Rewrite("$<.AnimVal")
	return animVal
}

// BaseVal
func (*SVGAnimatedLengthList) BaseVal() (baseVal *svglengthlist.SVGLengthList) {
	js.Rewrite("$<.BaseVal")
	return baseVal
}
