package svgclippathelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svgunittypes"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGClipPathElement,omit"
type SVGClipPathElement struct {
	window.SVGGraphicsElement
	svgunittypes.SVGUnitTypes
}

// ClipPathUnits
func (*SVGClipPathElement) ClipPathUnits() (clipPathUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.ClipPathUnits")
	return clipPathUnits
}
