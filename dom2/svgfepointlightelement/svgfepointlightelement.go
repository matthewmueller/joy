package svgfepointlightelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGFEPointLightElement,omit"
type SVGFEPointLightElement struct {
	window.SVGElement
}

// X
func (*SVGFEPointLightElement) X() (x *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFEPointLightElement) Y() (y *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Y")
	return y
}

// Z
func (*SVGFEPointLightElement) Z() (z *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Z")
	return z
}
