package svganimatedenumeration

import "github.com/matthewmueller/golly/js"

// SVGAnimatedEnumeration struct
// js:"SVGAnimatedEnumeration,omit"
type SVGAnimatedEnumeration struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedEnumeration) AnimVal() (animVal uint8) {
	js.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedEnumeration) BaseVal() (baseVal uint8) {
	js.Rewrite("$_.baseVal")
	return baseVal
}

// SetBaseVal prop
// js:"baseVal"
func (*SVGAnimatedEnumeration) SetBaseVal(baseVal uint8) {
	js.Rewrite("$_.baseVal = $1", baseVal)
}
