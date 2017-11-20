package svgfecomponenttransferelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedstring"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/js"
)

type SVGFEComponentTransferElement struct {
	*svgelement.SVGElement
	*svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEComponentTransferElement) GetIn1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}
