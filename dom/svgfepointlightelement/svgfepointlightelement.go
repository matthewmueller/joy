package svgfepointlightelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEPointLightElement struct
// js:"SVGFEPointLightElement,omit"
type SVGFEPointLightElement struct {
	window.SVGElement
}

// X prop
func (*SVGFEPointLightElement) X() (x *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFEPointLightElement) Y() (y *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.y")
	return y
}

// Z prop
func (*SVGFEPointLightElement) Z() (z *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.z")
	return z
}
