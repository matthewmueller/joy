package svgtextcontentelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svgpoint"
	"github.com/matthewmueller/golly/dom/svgrect"
	"github.com/matthewmueller/golly/dom/window"
)

// SVGTextContentElement interface
// js:"SVGTextContentElement"
type SVGTextContentElement interface {
	window.SVGGraphicsElement

	// GetCharNumAtPosition
	// js:"getCharNumAtPosition"
	// jsrewrite:"$_.getCharNumAtPosition($1)"
	GetCharNumAtPosition(point *svgpoint.SVGPoint) (i int)

	// GetComputedTextLength
	// js:"getComputedTextLength"
	// jsrewrite:"$_.getComputedTextLength()"
	GetComputedTextLength() (f float32)

	// GetEndPositionOfChar
	// js:"getEndPositionOfChar"
	// jsrewrite:"$_.getEndPositionOfChar($1)"
	GetEndPositionOfChar(charnum uint) (s *svgpoint.SVGPoint)

	// GetExtentOfChar
	// js:"getExtentOfChar"
	// jsrewrite:"$_.getExtentOfChar($1)"
	GetExtentOfChar(charnum uint) (s *svgrect.SVGRect)

	// GetNumberOfChars
	// js:"getNumberOfChars"
	// jsrewrite:"$_.getNumberOfChars()"
	GetNumberOfChars() (i int)

	// GetRotationOfChar
	// js:"getRotationOfChar"
	// jsrewrite:"$_.getRotationOfChar($1)"
	GetRotationOfChar(charnum uint) (f float32)

	// GetStartPositionOfChar
	// js:"getStartPositionOfChar"
	// jsrewrite:"$_.getStartPositionOfChar($1)"
	GetStartPositionOfChar(charnum uint) (s *svgpoint.SVGPoint)

	// GetSubStringLength
	// js:"getSubStringLength"
	// jsrewrite:"$_.getSubStringLength($1, $2)"
	GetSubStringLength(charnum uint, nchars uint) (f float32)

	// SelectSubString
	// js:"selectSubString"
	// jsrewrite:"$_.selectSubString($1, $2)"
	SelectSubString(charnum uint, nchars uint)

	// LengthAdjust prop
	// js:"lengthAdjust"
	// jsrewrite:"$_.lengthAdjust"
	LengthAdjust() (lengthAdjust *svganimatedenumeration.SVGAnimatedEnumeration)

	// TextLength prop
	// js:"textLength"
	// jsrewrite:"$_.textLength"
	TextLength() (textLength *svganimatedlength.SVGAnimatedLength)
}
