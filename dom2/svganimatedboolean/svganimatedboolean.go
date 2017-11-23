package svganimatedboolean

import "github.com/matthewmueller/golly/js"

// SVGAnimatedBoolean struct
// js:"SVGAnimatedBoolean,omit"
type SVGAnimatedBoolean struct {
}

// AnimVal
func (*SVGAnimatedBoolean) AnimVal() (animVal bool) {
	js.Rewrite("$<.AnimVal")
	return animVal
}

// BaseVal
func (*SVGAnimatedBoolean) BaseVal() (baseVal bool) {
	js.Rewrite("$<.BaseVal")
	return baseVal
}

// BaseVal
func (*SVGAnimatedBoolean) SetBaseVal(baseVal bool) {
	js.Rewrite("$<.BaseVal = $1", baseVal)
}
