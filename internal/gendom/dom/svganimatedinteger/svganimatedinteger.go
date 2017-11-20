package svganimatedinteger

import "github.com/matthewmueller/golly/js"

type SVGAnimatedInteger struct {
}

func (*SVGAnimatedInteger) GetAnimVal() (animVal int) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedInteger) GetBaseVal() (baseVal int) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

func (*SVGAnimatedInteger) SetBaseVal(baseVal int) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}
