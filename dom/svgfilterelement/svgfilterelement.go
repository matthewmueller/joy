package svgfilterelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatedinteger"
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFilterElement struct
// js:"SVGFilterElement,omit"
type SVGFilterElement struct {
	window.SVGElement
}

// SetFilterRes fn
func (*SVGFilterElement) SetFilterRes(filterResX uint, filterResY uint) {
	js.Rewrite("$<.setFilterRes($1, $2)", filterResX, filterResY)
}

// Href prop
func (*SVGFilterElement) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.href")
	return href
}

// FilterResX prop
func (*SVGFilterElement) FilterResX() (filterResX *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.filterResX")
	return filterResX
}

// FilterResY prop
func (*SVGFilterElement) FilterResY() (filterResY *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.filterResY")
	return filterResY
}

// FilterUnits prop
func (*SVGFilterElement) FilterUnits() (filterUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.filterUnits")
	return filterUnits
}

// Height prop
func (*SVGFilterElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// PrimitiveUnits prop
func (*SVGFilterElement) PrimitiveUnits() (primitiveUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.primitiveUnits")
	return primitiveUnits
}

// Width prop
func (*SVGFilterElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGFilterElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFilterElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
