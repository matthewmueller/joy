package svganimatedlengthlist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svglengthlist"
	"github.com/matthewmueller/golly/js"
)

type SVGAnimatedLengthList struct {
}

func (*SVGAnimatedLengthList) GetAnimVal() (animVal *svglengthlist.SVGLengthList) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedLengthList) GetBaseVal() (baseVal *svglengthlist.SVGLengthList) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
