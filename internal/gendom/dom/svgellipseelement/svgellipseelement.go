package svgellipseelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
	"github.com/matthewmueller/golly/js"
)

type SVGEllipseElement struct {
	*svggraphicselement.SVGGraphicsElement
}

func (*SVGEllipseElement) GetCx() (cx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.cx")
	return cx
}

func (*SVGEllipseElement) GetCy() (cy *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.cy")
	return cy
}

func (*SVGEllipseElement) GetRx() (rx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.rx")
	return rx
}

func (*SVGEllipseElement) GetRy() (ry *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.ry")
	return ry
}
