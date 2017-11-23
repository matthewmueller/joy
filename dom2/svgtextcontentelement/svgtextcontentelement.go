package svgtextcontentelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svgpoint"
	"github.com/matthewmueller/golly/dom2/svgrect"
	"github.com/matthewmueller/golly/dom2/window"
)

// js:"SVGTextContentElement,omit"
type SVGTextContentElement interface {
	window.SVGGraphicsElement

	// GetCharNumAtPosition
	GetCharNumAtPosition(point *svgpoint.SVGPoint) (i int)

	// GetComputedTextLength
	GetComputedTextLength() (f float32)

	// GetEndPositionOfChar
	GetEndPositionOfChar(charnum uint) (s *svgpoint.SVGPoint)

	// GetExtentOfChar
	GetExtentOfChar(charnum uint) (s *svgrect.SVGRect)

	// GetNumberOfChars
	GetNumberOfChars() (i int)

	// GetRotationOfChar
	GetRotationOfChar(charnum uint) (f float32)

	// GetStartPositionOfChar
	GetStartPositionOfChar(charnum uint) (s *svgpoint.SVGPoint)

	// GetSubStringLength
	GetSubStringLength(charnum uint, nchars uint) (f float32)

	// SelectSubString
	SelectSubString(charnum uint, nchars uint)

	// LengthAdjust
	LengthAdjust() (lengthAdjust *svganimatedenumeration.SVGAnimatedEnumeration)

	// TextLength
	TextLength() (textLength *svganimatedlength.SVGAnimatedLength)
}
