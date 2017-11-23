package svgfilterelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedinteger"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svgunittypes"
	"github.com/matthewmueller/golly/dom2/svgurireference"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGFilterElement,omit"
type SVGFilterElement struct {
	window.SVGElement
	svgunittypes.SVGUnitTypes
	svgurireference.SVGURIReference
}

// SetFilterRes
func (*SVGFilterElement) SetFilterRes(filterResX uint, filterResY uint) {
	js.Rewrite("$<.SetFilterRes($1, $2)", filterResX, filterResY)
}

// FilterResX
func (*SVGFilterElement) FilterResX() (filterResX *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.FilterResX")
	return filterResX
}

// FilterResY
func (*SVGFilterElement) FilterResY() (filterResY *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.FilterResY")
	return filterResY
}

// FilterUnits
func (*SVGFilterElement) FilterUnits() (filterUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.FilterUnits")
	return filterUnits
}

// Height
func (*SVGFilterElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// PrimitiveUnits
func (*SVGFilterElement) PrimitiveUnits() (primitiveUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.PrimitiveUnits")
	return primitiveUnits
}

// Width
func (*SVGFilterElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFilterElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFilterElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}
