package svgtextcontentelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svgpoint"
	"github.com/matthewmueller/golly/dom/svgrect"
	"github.com/matthewmueller/golly/dom/window"
)

// js:"SVGTextContentElement,omit"
type SVGTextContentElement interface {
	window.SVGGraphicsElement

	// GetCharNumAtPosition
	// js:"getCharNumAtPosition"
	GetCharNumAtPosition(point *svgpoint.SVGPoint) (i int)

	// GetComputedTextLength
	// js:"getComputedTextLength"
	GetComputedTextLength() (f float32)

	// GetEndPositionOfChar
	// js:"getEndPositionOfChar"
	GetEndPositionOfChar(charnum uint) (s *svgpoint.SVGPoint)

	// GetExtentOfChar
	// js:"getExtentOfChar"
	GetExtentOfChar(charnum uint) (s *svgrect.SVGRect)

	// GetNumberOfChars
	// js:"getNumberOfChars"
	GetNumberOfChars() (i int)

	// GetRotationOfChar
	// js:"getRotationOfChar"
	GetRotationOfChar(charnum uint) (f float32)

	// GetStartPositionOfChar
	// js:"getStartPositionOfChar"
	GetStartPositionOfChar(charnum uint) (s *svgpoint.SVGPoint)

	// GetSubStringLength
	// js:"getSubStringLength"
	GetSubStringLength(charnum uint, nchars uint) (f float32)

	// SelectSubString
	// js:"selectSubString"
	SelectSubString(charnum uint, nchars uint)

	// LengthAdjust prop
	// js:"lengthAdjust"
	LengthAdjust() (lengthAdjust *svganimatedenumeration.SVGAnimatedEnumeration)

	// TextLength prop
	// js:"textLength"
	TextLength() (textLength *svganimatedlength.SVGAnimatedLength)
}
