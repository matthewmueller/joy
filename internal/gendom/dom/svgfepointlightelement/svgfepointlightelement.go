package svgfepointlightelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumber"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/js"
)

type SVGFEPointLightElement struct {
	*svgelement.SVGElement
}

func (*SVGFEPointLightElement) GetX() (x *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGFEPointLightElement) GetY() (y *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGFEPointLightElement) GetZ() (z *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.z")
	return z
}
