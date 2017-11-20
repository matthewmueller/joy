package svganimatedtransformlist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgtransformlist"
	"github.com/matthewmueller/golly/js"
)

type SVGAnimatedTransformList struct {
}

func (*SVGAnimatedTransformList) GetAnimVal() (animVal *svgtransformlist.SVGTransformList) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedTransformList) GetBaseVal() (baseVal *svgtransformlist.SVGTransformList) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
