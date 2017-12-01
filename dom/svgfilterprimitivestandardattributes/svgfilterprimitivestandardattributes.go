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
	// jsrewrite:"$_.height"
	Height() (height *svganimatedlength.SVGAnimatedLength)

	// Result prop
	// js:"result"
	// jsrewrite:"$_.result"
	Result() (result *svganimatedstring.SVGAnimatedString)

	// Width prop
	// js:"width"
	// jsrewrite:"$_.width"
	Width() (width *svganimatedlength.SVGAnimatedLength)

	// X prop
	// js:"x"
	// jsrewrite:"$_.x"
	X() (x *svganimatedlength.SVGAnimatedLength)

	// Y prop
	// js:"y"
	// jsrewrite:"$_.y"
	Y() (y *svganimatedlength.SVGAnimatedLength)
}
