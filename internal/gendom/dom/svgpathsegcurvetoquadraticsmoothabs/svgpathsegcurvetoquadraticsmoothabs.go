package svgpathsegcurvetoquadraticsmoothabs

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegCurvetoQuadraticSmoothAbs struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
