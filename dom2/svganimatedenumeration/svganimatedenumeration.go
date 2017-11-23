package svganimatedenumeration

import "github.com/matthewmueller/golly/js"

// SVGAnimatedEnumeration struct
// js:"SVGAnimatedEnumeration,omit"
type SVGAnimatedEnumeration struct {
}

// AnimVal
func (*SVGAnimatedEnumeration) AnimVal() (animVal uint8) {
	js.Rewrite("$<.AnimVal")
	return animVal
}

// BaseVal
func (*SVGAnimatedEnumeration) BaseVal() (baseVal uint8) {
	js.Rewrite("$<.BaseVal")
	return baseVal
}

// BaseVal
func (*SVGAnimatedEnumeration) SetBaseVal(baseVal uint8) {
	js.Rewrite("$<.BaseVal = $1", baseVal)
}
