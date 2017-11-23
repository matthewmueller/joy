package svganimatedstring

import "github.com/matthewmueller/golly/js"

// js:"SVGAnimatedString,omit"
type SVGAnimatedString struct {
}

// AnimVal
func (*SVGAnimatedString) AnimVal() (animVal string) {
	js.Rewrite("$<.AnimVal")
	return animVal
}

// BaseVal
func (*SVGAnimatedString) BaseVal() (baseVal string) {
	js.Rewrite("$<.BaseVal")
	return baseVal
}

// BaseVal
func (*SVGAnimatedString) SetBaseVal(baseVal string) {
	js.Rewrite("$<.BaseVal = $1", baseVal)
}
