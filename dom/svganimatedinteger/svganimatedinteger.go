package svganimatedinteger

import "github.com/matthewmueller/joy/macro"

// SVGAnimatedInteger struct
// js:"SVGAnimatedInteger,omit"
type SVGAnimatedInteger struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedInteger) AnimVal() (animVal int) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedInteger) BaseVal() (baseVal int) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

// SetBaseVal prop
// js:"baseVal"
func (*SVGAnimatedInteger) SetBaseVal(baseVal int) {
	macro.Rewrite("$_.baseVal = $1", baseVal)
}
