package svganimatedstring

import "github.com/matthewmueller/golly/js"

// SVGAnimatedString struct
// js:"SVGAnimatedString,omit"
type SVGAnimatedString struct {
}

// AnimVal prop
func (*SVGAnimatedString) AnimVal() (animVal string) {
	js.Rewrite("$<.animVal")
	return animVal
}

// BaseVal prop
func (*SVGAnimatedString) BaseVal() (baseVal string) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

// BaseVal prop
func (*SVGAnimatedString) SetBaseVal(baseVal string) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}
