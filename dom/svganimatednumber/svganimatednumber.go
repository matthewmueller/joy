package svganimatednumber

import "github.com/matthewmueller/joy/macro"

// SVGAnimatedNumber struct
// js:"SVGAnimatedNumber,omit"
type SVGAnimatedNumber struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedNumber) AnimVal() (animVal float32) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedNumber) BaseVal() (baseVal float32) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

// SetBaseVal prop
// js:"baseVal"
func (*SVGAnimatedNumber) SetBaseVal(baseVal float32) {
	macro.Rewrite("$_.baseVal = $1", baseVal)
}
