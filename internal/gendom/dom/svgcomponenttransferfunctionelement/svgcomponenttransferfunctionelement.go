package svgcomponenttransferfunctionelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumber"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumberlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/js"
)

type SVGComponentTransferFunctionElement struct {
	*svgelement.SVGElement
}

func (*SVGComponentTransferFunctionElement) GetAmplitude() (amplitude *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.amplitude")
	return amplitude
}

func (*SVGComponentTransferFunctionElement) GetExponent() (exponent *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.exponent")
	return exponent
}

func (*SVGComponentTransferFunctionElement) GetIntercept() (intercept *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.intercept")
	return intercept
}

func (*SVGComponentTransferFunctionElement) GetOffset() (offset *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.offset")
	return offset
}

func (*SVGComponentTransferFunctionElement) GetSlope() (slope *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.slope")
	return slope
}

func (*SVGComponentTransferFunctionElement) GetTableValues() (tableValues *svganimatednumberlist.SVGAnimatedNumberList) {
	js.Rewrite("$<.tableValues")
	return tableValues
}

func (*SVGComponentTransferFunctionElement) GetType() (kind *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.type")
	return kind
}
