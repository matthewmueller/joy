package svgfilterprimitivestandardattributes

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
)

// js:"SVGFilterPrimitiveStandardAttributes,omit"
type SVGFilterPrimitiveStandardAttributes interface {

	// Height
	Height() (height *svganimatedlength.SVGAnimatedLength)

	// Result
	Result() (result *svganimatedstring.SVGAnimatedString)

	// Width
	Width() (width *svganimatedlength.SVGAnimatedLength)

	// X
	X() (x *svganimatedlength.SVGAnimatedLength)

	// Y
	Y() (y *svganimatedlength.SVGAnimatedLength)
}
