package svganimatedstring

import "github.com/matthewmueller/joy/js"

// SVGAnimatedString struct
// js:"SVGAnimatedString,omit"
type SVGAnimatedString struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedString) AnimVal() (animVal string) {
	js.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedString) BaseVal() (baseVal string) {
	js.Rewrite("$_.baseVal")
	return baseVal
}

// SetBaseVal prop
// js:"baseVal"
func (*SVGAnimatedString) SetBaseVal(baseVal string) {
	js.Rewrite("$_.baseVal = $1", baseVal)
}
