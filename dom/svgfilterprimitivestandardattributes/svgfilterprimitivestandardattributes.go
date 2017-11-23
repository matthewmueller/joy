package svgfilterprimitivestandardattributes

import (
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/js"
)

// SVGFilterPrimitiveStandardAttributes struct
// js:"SVGFilterPrimitiveStandardAttributes,omit"
type SVGFilterPrimitiveStandardAttributes struct {
}

// Height prop
func (*SVGFilterPrimitiveStandardAttributes) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Result prop
func (*SVGFilterPrimitiveStandardAttributes) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

// Width prop
func (*SVGFilterPrimitiveStandardAttributes) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGFilterPrimitiveStandardAttributes) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFilterPrimitiveStandardAttributes) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
