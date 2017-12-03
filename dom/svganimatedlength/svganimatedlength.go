package svganimatedlength

import (
	"github.com/matthewmueller/joy/dom/svglength"
	"github.com/matthewmueller/joy/js"
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
