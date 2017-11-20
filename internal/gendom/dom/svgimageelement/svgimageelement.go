package svgimageelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgurireference"
	"github.com/matthewmueller/golly/js"
)

type SVGImageElement struct {
	*svggraphicselement.SVGGraphicsElement
	*svgurireference.SVGURIReference
}

func (*SVGImageElement) GetHeight() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGImageElement) GetPreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.preserveAspectRatio")
	return preserveAspectRatio
}

func (*SVGImageElement) GetWidth() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGImageElement) GetX() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGImageElement) GetY() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
