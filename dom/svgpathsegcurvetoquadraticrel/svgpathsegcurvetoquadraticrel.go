package svgpathsegcurvetoquadraticrel

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegCurvetoQuadraticRel struct
// js:"SVGPathSegCurvetoQuadraticRel,omit"
type SVGPathSegCurvetoQuadraticRel struct {
	svgpathseg.SVGPathSeg
}

// X prop
func (*SVGPathSegCurvetoQuadraticRel) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegCurvetoQuadraticRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// X1 prop
func (*SVGPathSegCurvetoQuadraticRel) X1() (x1 float32) {
	js.Rewrite("$<.x1")
	return x1
}

// X1 prop
func (*SVGPathSegCurvetoQuadraticRel) SetX1(x1 float32) {
	js.Rewrite("$<.x1 = $1", x1)
}

// Y prop
func (*SVGPathSegCurvetoQuadraticRel) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegCurvetoQuadraticRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

// Y1 prop
func (*SVGPathSegCurvetoQuadraticRel) Y1() (y1 float32) {
	js.Rewrite("$<.y1")
	return y1
}

// Y1 prop
func (*SVGPathSegCurvetoQuadraticRel) SetY1(y1 float32) {
	js.Rewrite("$<.y1 = $1", y1)
}
