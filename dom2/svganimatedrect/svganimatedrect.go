package svganimatedrect

import (
	"github.com/matthewmueller/golly/dom2/svgrect"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGAnimatedRect,omit"
type SVGAnimatedRect struct {
}

// AnimVal
func (*SVGAnimatedRect) AnimVal() (animVal *svgrect.SVGRect) {
	js.Rewrite("$<.AnimVal")
	return animVal
}

// BaseVal
func (*SVGAnimatedRect) BaseVal() (baseVal *svgrect.SVGRect) {
	js.Rewrite("$<.BaseVal")
	return baseVal
}
