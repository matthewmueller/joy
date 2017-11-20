package svganimatednumberlist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgnumberlist"
	"github.com/matthewmueller/golly/js"
)

type SVGAnimatedNumberList struct {
}

func (*SVGAnimatedNumberList) GetAnimVal() (animVal *svgnumberlist.SVGNumberList) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedNumberList) GetBaseVal() (baseVal *svgnumberlist.SVGNumberList) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
