package svgforeignobjectelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
	"github.com/matthewmueller/golly/js"
)

type SVGForeignObjectElement struct {
	*svggraphicselement.SVGGraphicsElement
}

func (*SVGForeignObjectElement) GetHeight() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGForeignObjectElement) GetWidth() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGForeignObjectElement) GetX() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGForeignObjectElement) GetY() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
