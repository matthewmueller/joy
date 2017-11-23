package svgpathsegcurvetoquadraticsmoothrel

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGPathSegCurvetoQuadraticSmoothRel,omit"
type SVGPathSegCurvetoQuadraticSmoothRel struct {
	svgpathseg.SVGPathSeg
}

// X
func (*SVGPathSegCurvetoQuadraticSmoothRel) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegCurvetoQuadraticSmoothRel) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// Y
func (*SVGPathSegCurvetoQuadraticSmoothRel) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegCurvetoQuadraticSmoothRel) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}
