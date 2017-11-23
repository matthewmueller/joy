package svgviewelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGViewElement,omit"
type SVGViewElement struct {
	window.SVGElement
	svgzoomandpan.SVGZoomAndPan
	svgfittoviewbox.SVGFitToViewBox
}

// ViewTarget
func (*SVGViewElement) ViewTarget() (viewTarget *svgstringlist.SVGStringList) {
	js.Rewrite("$<.ViewTarget")
	return viewTarget
}
