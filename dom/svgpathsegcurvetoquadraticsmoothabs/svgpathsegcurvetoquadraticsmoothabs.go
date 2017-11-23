package svgpathsegcurvetoquadraticsmoothabs

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegCurvetoQuadraticSmoothAbs struct
// js:"SVGPathSegCurvetoQuadraticSmoothAbs,omit"
type SVGPathSegCurvetoQuadraticSmoothAbs struct {
	svgpathseg.SVGPathSeg
}

// X prop
func (*SVGPathSegCurvetoQuadraticSmoothAbs) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegCurvetoQuadraticSmoothAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// Y prop
func (*SVGPathSegCurvetoQuadraticSmoothAbs) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegCurvetoQuadraticSmoothAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
