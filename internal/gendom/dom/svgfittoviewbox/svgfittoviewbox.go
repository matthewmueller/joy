package svgfittoviewbox

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedrect"
	"github.com/matthewmueller/golly/js"
)

type SVGFitToViewBox struct {
}

func (*SVGFitToViewBox) GetPreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.preserveAspectRatio")
	return preserveAspectRatio
}

func (*SVGFitToViewBox) GetViewBox() (viewBox *svganimatedrect.SVGAnimatedRect) {
	js.Rewrite("$<.viewBox")
	return viewBox
}
