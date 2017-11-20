package svganimatedlength

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svglength"
	"github.com/matthewmueller/golly/js"
)

type SVGAnimatedLength struct {
}

func (*SVGAnimatedLength) GetAnimVal() (animVal *svglength.SVGLength) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedLength) GetBaseVal() (baseVal *svglength.SVGLength) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
