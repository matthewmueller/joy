package svganimatedrect

import (
	"github.com/matthewmueller/joy/dom/svgrect"
	"github.com/matthewmueller/joy/js"
)

// SVGAnimatedRect struct
// js:"SVGAnimatedRect,omit"
type SVGAnimatedRect struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedRect) AnimVal() (animVal *svgrect.SVGRect) {
	js.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedRect) BaseVal() (baseVal *svgrect.SVGRect) {
	js.Rewrite("$_.baseVal")
	return baseVal
}
