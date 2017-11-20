package svganimatedpoints

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpointlist"
	"github.com/matthewmueller/golly/js"
)

type SVGAnimatedPoints struct {
}

func (*SVGAnimatedPoints) GetAnimatedPoints() (animatedPoints *svgpointlist.SVGPointList) {
	js.Rewrite("$<.animatedPoints")
	return animatedPoints
}

func (*SVGAnimatedPoints) GetPoints() (points *svgpointlist.SVGPointList) {
	js.Rewrite("$<.points")
	return points
}
