package svgpathsegcurvetoquadraticabs

import (
	"github.com/matthewmueller/golly/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegCurvetoQuadraticAbs struct
// js:"SVGPathSegCurvetoQuadraticAbs,omit"
type SVGPathSegCurvetoQuadraticAbs struct {
	svgpathseg.SVGPathSeg
}

// X prop
func (*SVGPathSegCurvetoQuadraticAbs) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegCurvetoQuadraticAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// X1 prop
func (*SVGPathSegCurvetoQuadraticAbs) X1() (x1 float32) {
	js.Rewrite("$<.x1")
	return x1
}

// X1 prop
func (*SVGPathSegCurvetoQuadraticAbs) SetX1(x1 float32) {
	js.Rewrite("$<.x1 = $1", x1)
}

// Y prop
func (*SVGPathSegCurvetoQuadraticAbs) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegCurvetoQuadraticAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

// Y1 prop
func (*SVGPathSegCurvetoQuadraticAbs) Y1() (y1 float32) {
	js.Rewrite("$<.y1")
	return y1
}

// Y1 prop
func (*SVGPathSegCurvetoQuadraticAbs) SetY1(y1 float32) {
	js.Rewrite("$<.y1 = $1", y1)
}
