package svgfilterelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedinteger"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgunittypes"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgurireference"
	"github.com/matthewmueller/golly/js"
)

type SVGFilterElement struct {
	*svgelement.SVGElement
	*svgunittypes.SVGUnitTypes
	*svgurireference.SVGURIReference
}

func (*SVGFilterElement) SetFilterRes(filterResX uint, filterResY uint) {
	js.Rewrite("$<.setFilterRes($1, $2)", filterResX, filterResY)
}

func (*SVGFilterElement) GetFilterResX() (filterResX *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.filterResX")
	return filterResX
}

func (*SVGFilterElement) GetFilterResY() (filterResY *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.filterResY")
	return filterResY
}

func (*SVGFilterElement) GetFilterUnits() (filterUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.filterUnits")
	return filterUnits
}

func (*SVGFilterElement) GetHeight() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGFilterElement) GetPrimitiveUnits() (primitiveUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.primitiveUnits")
	return primitiveUnits
}

func (*SVGFilterElement) GetWidth() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGFilterElement) GetX() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGFilterElement) GetY() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
