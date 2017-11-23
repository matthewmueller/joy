package svganimatedrect

import (
	"github.com/matthewmueller/golly/dom/svgrect"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedRect struct
// js:"SVGAnimatedRect,omit"
type SVGAnimatedRect struct {
}

// AnimVal prop
func (*SVGAnimatedRect) AnimVal() (animVal *svgrect.SVGRect) {
	js.Rewrite("$<.animVal")
	return animVal
}

// BaseVal prop
func (*SVGAnimatedRect) BaseVal() (baseVal *svgrect.SVGRect) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
