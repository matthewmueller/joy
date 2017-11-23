package svganimatedboolean

import "github.com/matthewmueller/golly/js"

// SVGAnimatedBoolean struct
// js:"SVGAnimatedBoolean,omit"
type SVGAnimatedBoolean struct {
}

// AnimVal prop
func (*SVGAnimatedBoolean) AnimVal() (animVal bool) {
	js.Rewrite("$<.animVal")
	return animVal
}

// BaseVal prop
func (*SVGAnimatedBoolean) BaseVal() (baseVal bool) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

// BaseVal prop
func (*SVGAnimatedBoolean) SetBaseVal(baseVal bool) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}
