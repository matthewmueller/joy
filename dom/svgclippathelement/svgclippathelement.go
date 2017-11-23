package svgclippathelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGClipPathElement struct
// js:"SVGClipPathElement,omit"
type SVGClipPathElement struct {
	window.SVGGraphicsElement
}

// ClipPathUnits prop
func (*SVGClipPathElement) ClipPathUnits() (clipPathUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.clipPathUnits")
	return clipPathUnits
}
