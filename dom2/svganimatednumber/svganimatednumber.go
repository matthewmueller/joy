package svganimatednumber

import "github.com/matthewmueller/golly/js"

// SVGAnimatedNumber struct
// js:"SVGAnimatedNumber,omit"
type SVGAnimatedNumber struct {
}

// AnimVal
func (*SVGAnimatedNumber) AnimVal() (animVal float32) {
	js.Rewrite("$<.AnimVal")
	return animVal
}

// BaseVal
func (*SVGAnimatedNumber) BaseVal() (baseVal float32) {
	js.Rewrite("$<.BaseVal")
	return baseVal
}

// BaseVal
func (*SVGAnimatedNumber) SetBaseVal(baseVal float32) {
	js.Rewrite("$<.BaseVal = $1", baseVal)
}
