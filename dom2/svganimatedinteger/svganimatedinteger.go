package svganimatedinteger

import "github.com/matthewmueller/golly/js"

// SVGAnimatedInteger struct
// js:"SVGAnimatedInteger,omit"
type SVGAnimatedInteger struct {
}

// AnimVal
func (*SVGAnimatedInteger) AnimVal() (animVal int) {
	js.Rewrite("$<.AnimVal")
	return animVal
}

// BaseVal
func (*SVGAnimatedInteger) BaseVal() (baseVal int) {
	js.Rewrite("$<.BaseVal")
	return baseVal
}

// BaseVal
func (*SVGAnimatedInteger) SetBaseVal(baseVal int) {
	js.Rewrite("$<.BaseVal = $1", baseVal)
}
