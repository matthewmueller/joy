package svgpathsegcurvetocubicsmoothabs

import (
	"github.com/matthewmueller/golly/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegCurvetoCubicSmoothAbs struct
// js:"SVGPathSegCurvetoCubicSmoothAbs,omit"
type SVGPathSegCurvetoCubicSmoothAbs struct {
	svgpathseg.SVGPathSeg
}

// X prop
func (*SVGPathSegCurvetoCubicSmoothAbs) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegCurvetoCubicSmoothAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// X2 prop
func (*SVGPathSegCurvetoCubicSmoothAbs) X2() (x2 float32) {
	js.Rewrite("$<.x2")
	return x2
}

// X2 prop
func (*SVGPathSegCurvetoCubicSmoothAbs) SetX2(x2 float32) {
	js.Rewrite("$<.x2 = $1", x2)
}

// Y prop
func (*SVGPathSegCurvetoCubicSmoothAbs) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegCurvetoCubicSmoothAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

// Y2 prop
func (*SVGPathSegCurvetoCubicSmoothAbs) Y2() (y2 float32) {
	js.Rewrite("$<.y2")
	return y2
}

// Y2 prop
func (*SVGPathSegCurvetoCubicSmoothAbs) SetY2(y2 float32) {
	js.Rewrite("$<.y2 = $1", y2)
}
