package svgpathsegcurvetoquadraticrel

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegCurvetoQuadraticRel struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegCurvetoQuadraticRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoQuadraticRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoQuadraticRel) GetX1() (x1 float32) {
	js.Rewrite("$<.x1")
	return x1
}

func (*SVGPathSegCurvetoQuadraticRel) SetX1(x1 float32) {
	js.Rewrite("$<.x1 = $1", x1)
}

func (*SVGPathSegCurvetoQuadraticRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoQuadraticRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

func (*SVGPathSegCurvetoQuadraticRel) GetY1() (y1 float32) {
	js.Rewrite("$<.y1")
	return y1
}

func (*SVGPathSegCurvetoQuadraticRel) SetY1(y1 float32) {
	js.Rewrite("$<.y1 = $1", y1)
}
