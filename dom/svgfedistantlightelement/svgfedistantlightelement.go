package svgfedistantlightelement

import (
	"github.com/matthewmueller/golly/dom/svganimatednumber"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEDistantLightElement struct
// js:"SVGFEDistantLightElement,omit"
type SVGFEDistantLightElement struct {
	window.SVGElement
}

// Azimuth prop
func (*SVGFEDistantLightElement) Azimuth() (azimuth *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.azimuth")
	return azimuth
}

// Elevation prop
func (*SVGFEDistantLightElement) Elevation() (elevation *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.elevation")
	return elevation
}
