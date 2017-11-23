package svgpathsegcurvetoquadraticsmoothrel

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegCurvetoQuadraticSmoothRel struct
// js:"SVGPathSegCurvetoQuadraticSmoothRel,omit"
type SVGPathSegCurvetoQuadraticSmoothRel struct {
	svgpathseg.SVGPathSeg
}

// X prop
func (*SVGPathSegCurvetoQuadraticSmoothRel) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegCurvetoQuadraticSmoothRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// Y prop
func (*SVGPathSegCurvetoQuadraticSmoothRel) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegCurvetoQuadraticSmoothRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
