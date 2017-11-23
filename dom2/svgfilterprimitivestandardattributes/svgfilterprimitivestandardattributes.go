package svgfilterprimitivestandardattributes

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/js"
)

// SVGFilterPrimitiveStandardAttributes struct
// js:"SVGFilterPrimitiveStandardAttributes,omit"
type SVGFilterPrimitiveStandardAttributes struct {
}

// Height
func (*SVGFilterPrimitiveStandardAttributes) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFilterPrimitiveStandardAttributes) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFilterPrimitiveStandardAttributes) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFilterPrimitiveStandardAttributes) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFilterPrimitiveStandardAttributes) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}
