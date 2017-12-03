package svganimatedboolean

import "github.com/matthewmueller/joy/js"

// SVGAnimatedBoolean struct
// js:"SVGAnimatedBoolean,omit"
type SVGAnimatedBoolean struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedBoolean) AnimVal() (animVal bool) {
	js.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedBoolean) BaseVal() (baseVal bool) {
	js.Rewrite("$_.baseVal")
	return baseVal
}

// SetBaseVal prop
// js:"baseVal"
func (*SVGAnimatedBoolean) SetBaseVal(baseVal bool) {
	js.Rewrite("$_.baseVal = $1", baseVal)
}
