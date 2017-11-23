package svganimatednumberlist

import (
	"github.com/matthewmueller/golly/dom2/svgnumberlist"
	"github.com/matthewmueller/golly/js"
)

// SVGAnimatedNumberList struct
// js:"SVGAnimatedNumberList,omit"
type SVGAnimatedNumberList struct {
}

// AnimVal prop
func (*SVGAnimatedNumberList) AnimVal() (animVal *svgnumberlist.SVGNumberList) {
	js.Rewrite("$<.animVal")
	return animVal
}

// BaseVal prop
func (*SVGAnimatedNumberList) BaseVal() (baseVal *svgnumberlist.SVGNumberList) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
