package svgclippathelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGClipPathElement struct
// js:"SVGClipPathElement,omit"
type SVGClipPathElement struct {
	window.SVGGraphicsElement
}

// ClipPathUnits
func (*SVGClipPathElement) ClipPathUnits() (clipPathUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.ClipPathUnits")
	return clipPathUnits
}
