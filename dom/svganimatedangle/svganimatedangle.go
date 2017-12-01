package svganimatedangle

import (
	"github.com/matthewmueller/golly/dom/svgangle"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedAngle struct
// js:"SVGAnimatedAngle,omit"
type SVGAnimatedAngle struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedAngle) AnimVal() (animVal *svgangle.SVGAngle) {
	js.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedAngle) BaseVal() (baseVal *svgangle.SVGAngle) {
	js.Rewrite("$_.baseVal")
	return baseVal
}
