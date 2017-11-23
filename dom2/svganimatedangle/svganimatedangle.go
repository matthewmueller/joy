package svganimatedangle

import "github.com/matthewmueller/golly/js"

// js:"SVGAnimatedAngle,omit"
type SVGAnimatedAngle struct {
}

// AnimVal
func (*SVGAnimatedAngle) AnimVal() (animVal *svgangle.SVGAngle) {
	js.Rewrite("$<.AnimVal")
	return animVal
}

// BaseVal
func (*SVGAnimatedAngle) BaseVal() (baseVal *svgangle.SVGAngle) {
	js.Rewrite("$<.BaseVal")
	return baseVal
}
