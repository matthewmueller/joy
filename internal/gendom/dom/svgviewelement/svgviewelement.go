package svgviewelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfittoviewbox"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgstringlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgzoomandpan"
	"github.com/matthewmueller/golly/js"
)

type SVGViewElement struct {
	*svgelement.SVGElement
	*svgzoomandpan.SVGZoomAndPan
	*svgfittoviewbox.SVGFitToViewBox
}

func (*SVGViewElement) GetViewTarget() (viewTarget *svgstringlist.SVGStringList) {
	js.Rewrite("$<.viewTarget")
	return viewTarget
}
