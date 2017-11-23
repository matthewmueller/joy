package svganimatedinteger

import "github.com/matthewmueller/golly/js"

// SVGAnimatedInteger struct
// js:"SVGAnimatedInteger,omit"
type SVGAnimatedInteger struct {
}

// AnimVal prop
func (*SVGAnimatedInteger) AnimVal() (animVal int) {
	js.Rewrite("$<.animVal")
	return animVal
}

// BaseVal prop
func (*SVGAnimatedInteger) BaseVal() (baseVal int) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

// BaseVal prop
func (*SVGAnimatedInteger) SetBaseVal(baseVal int) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}
