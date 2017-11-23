package svganimatedlength

import (
	"github.com/matthewmueller/golly/dom2/svglength"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedLength struct
// js:"SVGAnimatedLength,omit"
type SVGAnimatedLength struct {
}

// AnimVal prop
func (*SVGAnimatedLength) AnimVal() (animVal *svglength.SVGLength) {
	js.Rewrite("$<.animVal")
	return animVal
}

// BaseVal prop
func (*SVGAnimatedLength) BaseVal() (baseVal *svglength.SVGLength) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
