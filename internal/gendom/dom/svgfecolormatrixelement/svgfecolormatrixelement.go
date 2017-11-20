package svgfecolormatrixelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumberlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedstring"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/js"
)

type SVGFEColorMatrixElement struct {
	*svgelement.SVGElement
	*svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEColorMatrixElement) GetIn1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFEColorMatrixElement) GetType() (kind *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.type")
	return kind
}

func (*SVGFEColorMatrixElement) GetValues() (values *svganimatednumberlist.SVGAnimatedNumberList) {
	js.Rewrite("$<.values")
	return values
}
