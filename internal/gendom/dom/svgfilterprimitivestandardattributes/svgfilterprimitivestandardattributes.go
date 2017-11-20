package svgfilterprimitivestandardattributes

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedstring"
	"github.com/matthewmueller/golly/js"
)

type SVGFilterPrimitiveStandardAttributes struct {
}

func (*SVGFilterPrimitiveStandardAttributes) GetHeight() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGFilterPrimitiveStandardAttributes) GetResult() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

func (*SVGFilterPrimitiveStandardAttributes) GetWidth() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGFilterPrimitiveStandardAttributes) GetX() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGFilterPrimitiveStandardAttributes) GetY() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
