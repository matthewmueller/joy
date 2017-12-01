package svganimatedinteger

import "github.com/matthewmueller/golly/js"

// SVGAnimatedInteger struct
// js:"SVGAnimatedInteger,omit"
type SVGAnimatedInteger struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedInteger) AnimVal() (animVal int) {
	js.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedInteger) BaseVal() (baseVal int) {
	js.Rewrite("$_.baseVal")
	return baseVal
}

// SetBaseVal prop
// js:"baseVal"
func (*SVGAnimatedInteger) SetBaseVal(baseVal int) {
	js.Rewrite("$_.baseVal = $1", baseVal)
}
