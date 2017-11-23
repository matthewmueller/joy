package svgtextcontentelement

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
