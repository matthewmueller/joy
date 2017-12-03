package svganimatedstring

import "github.com/matthewmueller/joy/macro"

// SVGAnimatedString struct
// js:"SVGAnimatedString,omit"
type SVGAnimatedString struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedString) AnimVal() (animVal string) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedString) BaseVal() (baseVal string) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}

// SetBaseVal prop
// js:"baseVal"
func (*SVGAnimatedString) SetBaseVal(baseVal string) {
	macro.Rewrite("$_.baseVal = $1", baseVal)
}
