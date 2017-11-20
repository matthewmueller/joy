package svgrectelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
	"github.com/matthewmueller/golly/js"
)

type SVGRectElement struct {
	*svggraphicselement.SVGGraphicsElement
}

func (*SVGRectElement) GetHeight() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGRectElement) GetRx() (rx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.rx")
	return rx
}

func (*SVGRectElement) GetRy() (ry *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.ry")
	return ry
}

func (*SVGRectElement) GetWidth() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGRectElement) GetX() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGRectElement) GetY() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
