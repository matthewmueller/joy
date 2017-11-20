package svganimatedrect

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgrect"
	"github.com/matthewmueller/golly/js"
)

type SVGAnimatedRect struct {
}

func (*SVGAnimatedRect) GetAnimVal() (animVal *svgrect.SVGRect) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedRect) GetBaseVal() (baseVal *svgrect.SVGRect) {
	js.Rewrite("$<.baseVal")
	return baseVal
}
