package svganimatedboolean

import "github.com/matthewmueller/golly/js"

type SVGAnimatedBoolean struct {
}

func (*SVGAnimatedBoolean) GetAnimVal() (animVal bool) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedBoolean) GetBaseVal() (baseVal bool) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

func (*SVGAnimatedBoolean) SetBaseVal(baseVal bool) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}
