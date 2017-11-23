package svgpathsegcurvetoquadraticsmoothabs

import "github.com/matthewmueller/golly/js"

// js:"SVGPathSegCurvetoQuadraticSmoothAbs,omit"
type SVGPathSegCurvetoQuadraticSmoothAbs struct {
	svgpathseg.SVGPathSeg
}

// X
func (*SVGPathSegCurvetoQuadraticSmoothAbs) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegCurvetoQuadraticSmoothAbs) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// Y
func (*SVGPathSegCurvetoQuadraticSmoothAbs) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegCurvetoQuadraticSmoothAbs) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}
