package svganimatedangle

import (
	"github.com/matthewmueller/joy/dom/svgangle"
	"github.com/matthewmueller/joy/js"
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
