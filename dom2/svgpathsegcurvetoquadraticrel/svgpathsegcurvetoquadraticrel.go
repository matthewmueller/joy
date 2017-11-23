package svgpathsegcurvetoquadraticrel

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGPathSegCurvetoQuadraticRel,omit"
type SVGPathSegCurvetoQuadraticRel struct {
	svgpathseg.SVGPathSeg
}

// X
func (*SVGPathSegCurvetoQuadraticRel) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegCurvetoQuadraticRel) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// X1
func (*SVGPathSegCurvetoQuadraticRel) X1() (x1 float32) {
	js.Rewrite("$<.X1")
	return x1
}

// X1
func (*SVGPathSegCurvetoQuadraticRel) SetX1(x1 float32) {
	js.Rewrite("$<.X1 = $1", x1)
}

// Y
func (*SVGPathSegCurvetoQuadraticRel) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegCurvetoQuadraticRel) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}

// Y1
func (*SVGPathSegCurvetoQuadraticRel) Y1() (y1 float32) {
	js.Rewrite("$<.Y1")
	return y1
}

// Y1
func (*SVGPathSegCurvetoQuadraticRel) SetY1(y1 float32) {
	js.Rewrite("$<.Y1 = $1", y1)
}
