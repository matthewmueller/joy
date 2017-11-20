package svguseelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelementinstance"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgurireference"
	"github.com/matthewmueller/golly/js"
)

type SVGUseElement struct {
	*svggraphicselement.SVGGraphicsElement
	*svgurireference.SVGURIReference
}

func (*SVGUseElement) GetAnimatedInstanceRoot() (animatedInstanceRoot *svgelementinstance.SVGElementInstance) {
	js.Rewrite("$<.animatedInstanceRoot")
	return animatedInstanceRoot
}

func (*SVGUseElement) GetHeight() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGUseElement) GetInstanceRoot() (instanceRoot *svgelementinstance.SVGElementInstance) {
	js.Rewrite("$<.instanceRoot")
	return instanceRoot
}

func (*SVGUseElement) GetWidth() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGUseElement) GetX() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGUseElement) GetY() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
