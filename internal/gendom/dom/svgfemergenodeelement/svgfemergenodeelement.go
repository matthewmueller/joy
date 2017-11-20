package svgfemergenodeelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedstring"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/js"
)

type SVGFEMergeNodeElement struct {
	*svgelement.SVGElement
}

func (*SVGFEMergeNodeElement) GetIn1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}
