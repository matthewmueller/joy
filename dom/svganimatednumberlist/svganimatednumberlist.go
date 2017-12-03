package svganimatednumberlist

import (
	"github.com/matthewmueller/joy/dom/svgnumberlist"
	"github.com/matthewmueller/joy/macro"
)

// SVGAnimatedNumberList struct
// js:"SVGAnimatedNumberList,omit"
type SVGAnimatedNumberList struct {
}

// AnimVal prop
// js:"animVal"
func (*SVGAnimatedNumberList) AnimVal() (animVal *svgnumberlist.SVGNumberList) {
	macro.Rewrite("$_.animVal")
	return animVal
}

// BaseVal prop
// js:"baseVal"
func (*SVGAnimatedNumberList) BaseVal() (baseVal *svgnumberlist.SVGNumberList) {
	macro.Rewrite("$_.baseVal")
	return baseVal
}
