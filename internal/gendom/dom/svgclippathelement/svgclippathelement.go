package svgclippathelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgunittypes"
	"github.com/matthewmueller/golly/js"
)

type SVGClipPathElement struct {
	*svggraphicselement.SVGGraphicsElement
	*svgunittypes.SVGUnitTypes
}

func (*SVGClipPathElement) GetClipPathUnits() (clipPathUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.clipPathUnits")
	return clipPathUnits
}
