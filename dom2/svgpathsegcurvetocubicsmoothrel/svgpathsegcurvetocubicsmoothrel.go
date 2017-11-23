package svgpathsegcurvetocubicsmoothrel

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGPathSegCurvetoCubicSmoothRel,omit"
type SVGPathSegCurvetoCubicSmoothRel struct {
	svgpathseg.SVGPathSeg
}

// X
func (*SVGPathSegCurvetoCubicSmoothRel) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegCurvetoCubicSmoothRel) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// X2
func (*SVGPathSegCurvetoCubicSmoothRel) X2() (x2 float32) {
	js.Rewrite("$<.X2")
	return x2
}

// X2
func (*SVGPathSegCurvetoCubicSmoothRel) SetX2(x2 float32) {
	js.Rewrite("$<.X2 = $1", x2)
}

// Y
func (*SVGPathSegCurvetoCubicSmoothRel) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegCurvetoCubicSmoothRel) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}

// Y2
func (*SVGPathSegCurvetoCubicSmoothRel) Y2() (y2 float32) {
	js.Rewrite("$<.Y2")
	return y2
}

// Y2
func (*SVGPathSegCurvetoCubicSmoothRel) SetY2(y2 float32) {
	js.Rewrite("$<.Y2 = $1", y2)
}
