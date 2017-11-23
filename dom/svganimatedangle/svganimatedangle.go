package svganimatedangle

import (
	"github.com/matthewmueller/golly/dom2/svgangle"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedAngle struct
// js:"SVGAnimatedAngle,omit"
type SVGAnimatedAngle struct {
}

// AnimVal prop
func (*SVGAnimatedAngle) AnimVal() (animVal *svgangle.SVGAngle) {
	js.Rewrite("$<.animVal")
	return animVal
}

// BaseVal prop
func (*SVGAnimatedAngle) BaseVal() (baseVal *svgangle.SVGAngle) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
