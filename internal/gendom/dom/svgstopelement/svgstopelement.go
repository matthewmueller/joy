package svgstopelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumber"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/js"
)

type SVGStopElement struct {
	*svgelement.SVGElement
}

func (*SVGStopElement) GetOffset() (offset *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.offset")
	return offset
}
