package svganimatednumber

import "github.com/matthewmueller/golly/js"

// SVGAnimatedNumber struct
// js:"SVGAnimatedNumber,omit"
type SVGAnimatedNumber struct {
}

// AnimVal prop
func (*SVGAnimatedNumber) AnimVal() (animVal float32) {
	js.Rewrite("$<.animVal")
	return animVal
}

// BaseVal prop
func (*SVGAnimatedNumber) BaseVal() (baseVal float32) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

// BaseVal prop
func (*SVGAnimatedNumber) SetBaseVal(baseVal float32) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}
