package svgpathsegcurvetoquadraticabs

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGPathSegCurvetoQuadraticAbs,omit"
type SVGPathSegCurvetoQuadraticAbs struct {
	svgpathseg.SVGPathSeg
}

// X
func (*SVGPathSegCurvetoQuadraticAbs) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegCurvetoQuadraticAbs) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// X1
func (*SVGPathSegCurvetoQuadraticAbs) X1() (x1 float32) {
	js.Rewrite("$<.X1")
	return x1
}

// X1
func (*SVGPathSegCurvetoQuadraticAbs) SetX1(x1 float32) {
	js.Rewrite("$<.X1 = $1", x1)
}

// Y
func (*SVGPathSegCurvetoQuadraticAbs) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegCurvetoQuadraticAbs) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}

// Y1
func (*SVGPathSegCurvetoQuadraticAbs) Y1() (y1 float32) {
	js.Rewrite("$<.Y1")
	return y1
}

// Y1
func (*SVGPathSegCurvetoQuadraticAbs) SetY1(y1 float32) {
	js.Rewrite("$<.Y1 = $1", y1)
}
