package svgtextcontentelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svgpoint"
	"github.com/matthewmueller/golly/dom/svgrect"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGTextContentElement,omit"
type SVGTextContentElement interface {
	window.SVGGraphicsElement

	// GetCharNumAtPosition
	// js:"getCharNumAtPosition,rewrite=getcharnumatposition"
	GetCharNumAtPosition(point *svgpoint.SVGPoint) (i int)

	// GetComputedTextLength
	// js:"getComputedTextLength,rewrite=getcomputedtextlength"
	GetComputedTextLength() (f float32)

	// GetEndPositionOfChar
	// js:"getEndPositionOfChar,rewrite=getendpositionofchar"
	GetEndPositionOfChar(charnum uint) (s *svgpoint.SVGPoint)

	// GetExtentOfChar
	// js:"getExtentOfChar,rewrite=getextentofchar"
	GetExtentOfChar(charnum uint) (s *svgrect.SVGRect)

	// GetNumberOfChars
	// js:"getNumberOfChars,rewrite=getnumberofchars"
	GetNumberOfChars() (i int)

	// GetRotationOfChar
	// js:"getRotationOfChar,rewrite=getrotationofchar"
	GetRotationOfChar(charnum uint) (f float32)

	// GetStartPositionOfChar
	// js:"getStartPositionOfChar,rewrite=getstartpositionofchar"
	GetStartPositionOfChar(charnum uint) (s *svgpoint.SVGPoint)

	// GetSubStringLength
	// js:"getSubStringLength,rewrite=getsubstringlength"
	GetSubStringLength(charnum uint, nchars uint) (f float32)

	// SelectSubString
	// js:"selectSubString,rewrite=selectsubstring"
	SelectSubString(charnum uint, nchars uint)

	// LengthAdjust prop
	// js:"lengthAdjust,rewrite=lengthadjust"
	LengthAdjust() (lengthAdjust *svganimatedenumeration.SVGAnimatedEnumeration)

	// TextLength prop
	// js:"textLength,rewrite=textlength"
	TextLength() (textLength *svganimatedlength.SVGAnimatedLength)
}

// getcharnumatposition fn
func getcharnumatposition(point *svgpoint.SVGPoint) (i int) {
	js.Rewrite("$<.getCharNumAtPosition($1)", point)
	return i
}

// getcomputedtextlength fn
func getcomputedtextlength() (f float32) {
	js.Rewrite("$<.getComputedTextLength()")
	return f
}

// getendpositionofchar fn
func getendpositionofchar(charnum uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.getEndPositionOfChar($1)", charnum)
	return s
}

// getextentofchar fn
func getextentofchar(charnum uint) (s *svgrect.SVGRect) {
	js.Rewrite("$<.getExtentOfChar($1)", charnum)
	return s
}

// getnumberofchars fn
func getnumberofchars() (i int) {
	js.Rewrite("$<.getNumberOfChars()")
	return i
}

// getrotationofchar fn
func getrotationofchar(charnum uint) (f float32) {
	js.Rewrite("$<.getRotationOfChar($1)", charnum)
	return f
}

// getstartpositionofchar fn
func getstartpositionofchar(charnum uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.getStartPositionOfChar($1)", charnum)
	return s
}

// getsubstringlength fn
func getsubstringlength(charnum uint, nchars uint) (f float32) {
	js.Rewrite("$<.getSubStringLength($1, $2)", charnum, nchars)
	return f
}

// selectsubstring fn
func selectsubstring(charnum uint, nchars uint) {
	js.Rewrite("$<.selectSubString($1, $2)", charnum, nchars)
}

// lengthadjust prop
func lengthadjust() (lengthAdjust *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.lengthAdjust")
	return lengthAdjust
}

// textlength prop
func textlength() (textLength *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.textLength")
	return textLength
}
