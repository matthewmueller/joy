package svgpathsegcurvetocubicsmoothabs

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGPathSegCurvetoCubicSmoothAbs,omit"
type SVGPathSegCurvetoCubicSmoothAbs struct {
	svgpathseg.SVGPathSeg
}

// X
func (*SVGPathSegCurvetoCubicSmoothAbs) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegCurvetoCubicSmoothAbs) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// X2
func (*SVGPathSegCurvetoCubicSmoothAbs) X2() (x2 float32) {
	js.Rewrite("$<.X2")
	return x2
}

// X2
func (*SVGPathSegCurvetoCubicSmoothAbs) SetX2(x2 float32) {
	js.Rewrite("$<.X2 = $1", x2)
}

// Y
func (*SVGPathSegCurvetoCubicSmoothAbs) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegCurvetoCubicSmoothAbs) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}

// Y2
func (*SVGPathSegCurvetoCubicSmoothAbs) Y2() (y2 float32) {
	js.Rewrite("$<.Y2")
	return y2
}

// Y2
func (*SVGPathSegCurvetoCubicSmoothAbs) SetY2(y2 float32) {
	js.Rewrite("$<.Y2 = $1", y2)
}
