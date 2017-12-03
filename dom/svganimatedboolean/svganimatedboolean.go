package svganimatedboolean

import "github.com/matthewmueller/joy/macro"

// SVGAnimatedBoolean struct
// js:"SVGAnimatedBoolean,omit"
type SVGAnimatedBoolean struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedBoolean) AnimVal() (animVal bool) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedBoolean) BaseVal() (baseVal bool) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

// SetBaseVal prop
// js:"baseVal"
func (*SVGAnimatedBoolean) SetBaseVal(baseVal bool) {
	macro.Rewrite("$_.baseVal = $1", baseVal)
}
