package svganimatednumber

import "github.com/matthewmueller/golly/js"

type SVGAnimatedNumber struct {
}

func (*SVGAnimatedNumber) GetAnimVal() (animVal float32) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedNumber) GetBaseVal() (baseVal float32) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

func (*SVGAnimatedNumber) SetBaseVal(baseVal float32) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}
