package svganimatedenumeration

import "github.com/matthewmueller/golly/js"

// SVGAnimatedEnumeration struct
// js:"SVGAnimatedEnumeration,omit"
type SVGAnimatedEnumeration struct {
}

// AnimVal prop
func (*SVGAnimatedEnumeration) AnimVal() (animVal uint8) {
	js.Rewrite("$<.animVal")
	return animVal
}

// BaseVal prop
func (*SVGAnimatedEnumeration) BaseVal() (baseVal uint8) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

// BaseVal prop
func (*SVGAnimatedEnumeration) SetBaseVal(baseVal uint8) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}
