package svglineelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
	"github.com/matthewmueller/golly/js"
)

type SVGLineElement struct {
	*svggraphicselement.SVGGraphicsElement
}

func (*SVGLineElement) GetX1() (x1 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x1")
	return x1
}

func (*SVGLineElement) GetX2() (x2 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x2")
	return x2
}

func (*SVGLineElement) GetY1() (y1 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y1")
	return y1
}

func (*SVGLineElement) GetY2() (y2 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y2")
	return y2
}
