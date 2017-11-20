package svgpolylineelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedpoints"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
)

type SVGPolylineElement struct {
	*svggraphicselement.SVGGraphicsElement
	*svganimatedpoints.SVGAnimatedPoints
}
