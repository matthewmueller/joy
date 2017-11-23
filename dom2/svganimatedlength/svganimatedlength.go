package svganimatedlength

import (
	"github.com/matthewmueller/golly/dom2/svglength"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedLength struct
// js:"SVGAnimatedLength,omit"
type SVGAnimatedLength struct {
}

// AnimVal
func (*SVGAnimatedLength) AnimVal() (animVal *svglength.SVGLength) {
	js.Rewrite("$<.AnimVal")
	return animVal
}

// BaseVal
func (*SVGAnimatedLength) BaseVal() (baseVal *svglength.SVGLength) {
	js.Rewrite("$<.BaseVal")
	return baseVal
}
