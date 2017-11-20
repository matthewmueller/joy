package svgpolygonelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedpoints"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
)

type SVGPolygonElement struct {
	*svggraphicselement.SVGGraphicsElement
	*svganimatedpoints.SVGAnimatedPoints
}
