package svganimatedstring

import "github.com/matthewmueller/golly/js"

type SVGAnimatedString struct {
}

func (*SVGAnimatedString) GetAnimVal() (animVal string) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedString) GetBaseVal() (baseVal string) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

func (*SVGAnimatedString) SetBaseVal(baseVal string) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}
