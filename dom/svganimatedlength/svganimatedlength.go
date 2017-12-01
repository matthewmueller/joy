package svganimatedlength

import (
	"github.com/matthewmueller/golly/dom/svglength"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedLength struct
// js:"SVGAnimatedLength,omit"
type SVGAnimatedLength struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedLength) AnimVal() (animVal *svglength.SVGLength) {
	js.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedLength) BaseVal() (baseVal *svglength.SVGLength) {
	js.Rewrite("$_.baseVal")
	return baseVal
}
