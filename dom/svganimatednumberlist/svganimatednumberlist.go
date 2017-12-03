package svganimatednumberlist

import (
	"github.com/matthewmueller/joy/dom/svgnumberlist"
	"github.com/matthewmueller/joy/js"
)

// SVGAnimatedNumberList struct
// js:"SVGAnimatedNumberList,omit"
type SVGAnimatedNumberList struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedNumberList) AnimVal() (animVal *svgnumberlist.SVGNumberList) {
	js.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedNumberList) BaseVal() (baseVal *svgnumberlist.SVGNumberList) {
	js.Rewrite("$_.baseVal")
	return baseVal
}
