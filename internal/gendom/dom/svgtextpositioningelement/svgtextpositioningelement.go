package svgtextpositioningelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlengthlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumberlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgtextcontentelement"
	"github.com/matthewmueller/golly/js"
)

type SVGTextPositioningElement struct {
	*svgtextcontentelement.SVGTextContentElement
}

func (*SVGTextPositioningElement) GetDx() (dx *svganimatedlengthlist.SVGAnimatedLengthList) {
	js.Rewrite("$<.dx")
	return dx
}

func (*SVGTextPositioningElement) GetDy() (dy *svganimatedlengthlist.SVGAnimatedLengthList) {
	js.Rewrite("$<.dy")
	return dy
}

func (*SVGTextPositioningElement) GetRotate() (rotate *svganimatednumberlist.SVGAnimatedNumberList) {
	js.Rewrite("$<.rotate")
	return rotate
}

func (*SVGTextPositioningElement) GetX() (x *svganimatedlengthlist.SVGAnimatedLengthList) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGTextPositioningElement) GetY() (y *svganimatedlengthlist.SVGAnimatedLengthList) {
	js.Rewrite("$<.y")
	return y
}
