package svgfeoffsetelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumber"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedstring"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/js"
)

type SVGFEOffsetElement struct {
	*svgelement.SVGElement
	*svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEOffsetElement) GetDx() (dx *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.dx")
	return dx
}

func (*SVGFEOffsetElement) GetDy() (dy *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.dy")
	return dy
}

func (*SVGFEOffsetElement) GetIn1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}
