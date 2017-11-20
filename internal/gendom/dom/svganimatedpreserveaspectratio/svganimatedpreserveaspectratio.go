package svganimatedpreserveaspectratio

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpreserveaspectratio"
	"github.com/matthewmueller/golly/js"
)

type SVGAnimatedPreserveAspectRatio struct {
}

func (*SVGAnimatedPreserveAspectRatio) GetAnimVal() (animVal *svgpreserveaspectratio.SVGPreserveAspectRatio) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedPreserveAspectRatio) GetBaseVal() (baseVal *svgpreserveaspectratio.SVGPreserveAspectRatio) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
