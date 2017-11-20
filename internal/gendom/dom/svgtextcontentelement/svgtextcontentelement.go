package svgtextcontentelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpoint"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgrect"
	"github.com/matthewmueller/golly/js"
)

type SVGTextContentElement struct {
	*svggraphicselement.SVGGraphicsElement
}

func (*SVGTextContentElement) GetCharNumAtPosition(point *svgpoint.SVGPoint) (i int) {
	js.Rewrite("$<.getCharNumAtPosition($1)", point)
	return i
}

func (*SVGTextContentElement) GetComputedTextLength() (f float32) {
	js.Rewrite("$<.getComputedTextLength()")
	return f
}

func (*SVGTextContentElement) GetEndPositionOfChar(charnum uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.getEndPositionOfChar($1)", charnum)
	return s
}

func (*SVGTextContentElement) GetExtentOfChar(charnum uint) (s *svgrect.SVGRect) {
	js.Rewrite("$<.getExtentOfChar($1)", charnum)
	return s
}

func (*SVGTextContentElement) GetNumberOfChars() (i int) {
	js.Rewrite("$<.getNumberOfChars()")
	return i
}

func (*SVGTextContentElement) GetRotationOfChar(charnum uint) (f float32) {
	js.Rewrite("$<.getRotationOfChar($1)", charnum)
	return f
}

func (*SVGTextContentElement) GetStartPositionOfChar(charnum uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.getStartPositionOfChar($1)", charnum)
	return s
}

func (*SVGTextContentElement) GetSubStringLength(charnum uint, nchars uint) (f float32) {
	js.Rewrite("$<.getSubStringLength($1, $2)", charnum, nchars)
	return f
}

func (*SVGTextContentElement) SelectSubString(charnum uint, nchars uint) {
	js.Rewrite("$<.selectSubString($1, $2)", charnum, nchars)
}

func (*SVGTextContentElement) GetLengthAdjust() (lengthAdjust *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.lengthAdjust")
	return lengthAdjust
}

func (*SVGTextContentElement) GetTextLength() (textLength *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.textLength")
	return textLength
}
