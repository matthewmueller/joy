package svganimatedangle

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgangle"
	"github.com/matthewmueller/golly/js"
)

type SVGAnimatedAngle struct {
}

func (*SVGAnimatedAngle) GetAnimVal() (animVal *svgangle.SVGAngle) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedAngle) GetBaseVal() (baseVal *svgangle.SVGAngle) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
