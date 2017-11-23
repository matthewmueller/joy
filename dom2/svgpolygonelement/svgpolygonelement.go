package svgpolygonelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedpoints"
	"github.com/matthewmueller/golly/dom2/window"
)

// js:"SVGPolygonElement,omit"
type SVGPolygonElement struct {
	window.SVGGraphicsElement
	svganimatedpoints.SVGAnimatedPoints
}
