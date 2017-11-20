package svgpathsegcurvetoquadraticsmoothrel

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegCurvetoQuadraticSmoothRel struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegCurvetoQuadraticSmoothRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoQuadraticSmoothRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoQuadraticSmoothRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoQuadraticSmoothRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
