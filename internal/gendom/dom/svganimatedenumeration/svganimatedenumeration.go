package svganimatedenumeration

import "github.com/matthewmueller/golly/js"

type SVGAnimatedEnumeration struct {
}

func (*SVGAnimatedEnumeration) GetAnimVal() (animVal uint8) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedEnumeration) GetBaseVal() (baseVal uint8) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

func (*SVGAnimatedEnumeration) SetBaseVal(baseVal uint8) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}
