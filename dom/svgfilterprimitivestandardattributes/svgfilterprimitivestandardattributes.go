package svgfilterprimitivestandardattributes

import (
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
)

// SVGFilterPrimitiveStandardAttributes interface
// js:"SVGFilterPrimitiveStandardAttributes"
type SVGFilterPrimitiveStandardAttributes interface {

	// Height prop
	// js:"height"
	Height() (height *svganimatedlength.SVGAnimatedLength)

	// Result prop
	// js:"result"
	Result() (result *svganimatedstring.SVGAnimatedString)

	// Width prop
	// js:"width"
	Width() (width *svganimatedlength.SVGAnimatedLength)

	// X prop
	// js:"x"
	X() (x *svganimatedlength.SVGAnimatedLength)

	// Y prop
	// js:"y"
	Y() (y *svganimatedlength.SVGAnimatedLength)
}
