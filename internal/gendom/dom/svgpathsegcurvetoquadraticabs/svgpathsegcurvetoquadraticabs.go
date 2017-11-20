package svgpathsegcurvetoquadraticabs

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegCurvetoQuadraticAbs struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegCurvetoQuadraticAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoQuadraticAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoQuadraticAbs) GetX1() (x1 float32) {
	js.Rewrite("$<.x1")
	return x1
}

func (*SVGPathSegCurvetoQuadraticAbs) SetX1(x1 float32) {
	js.Rewrite("$<.x1 = $1", x1)
}

func (*SVGPathSegCurvetoQuadraticAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoQuadraticAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

func (*SVGPathSegCurvetoQuadraticAbs) GetY1() (y1 float32) {
	js.Rewrite("$<.y1")
	return y1
}

func (*SVGPathSegCurvetoQuadraticAbs) SetY1(y1 float32) {
	js.Rewrite("$<.y1 = $1", y1)
}
