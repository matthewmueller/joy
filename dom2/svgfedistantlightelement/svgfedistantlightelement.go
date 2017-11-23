package svgfedistantlightelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGFEDistantLightElement,omit"
type SVGFEDistantLightElement struct {
	window.SVGElement
}

// Azimuth
func (*SVGFEDistantLightElement) Azimuth() (azimuth *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Azimuth")
	return azimuth
}

// Elevation
func (*SVGFEDistantLightElement) Elevation() (elevation *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Elevation")
	return elevation
}
