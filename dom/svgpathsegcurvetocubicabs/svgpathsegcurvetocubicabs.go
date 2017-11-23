package svgpathsegcurvetocubicabs

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegCurvetoCubicAbs struct
// js:"SVGPathSegCurvetoCubicAbs,omit"
type SVGPathSegCurvetoCubicAbs struct {
	svgpathseg.SVGPathSeg
}

// X prop
func (*SVGPathSegCurvetoCubicAbs) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegCurvetoCubicAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// X1 prop
func (*SVGPathSegCurvetoCubicAbs) X1() (x1 float32) {
	js.Rewrite("$<.x1")
	return x1
}

// X1 prop
func (*SVGPathSegCurvetoCubicAbs) SetX1(x1 float32) {
	js.Rewrite("$<.x1 = $1", x1)
}

// X2 prop
func (*SVGPathSegCurvetoCubicAbs) X2() (x2 float32) {
	js.Rewrite("$<.x2")
	return x2
}

// X2 prop
func (*SVGPathSegCurvetoCubicAbs) SetX2(x2 float32) {
	js.Rewrite("$<.x2 = $1", x2)
}

// Y prop
func (*SVGPathSegCurvetoCubicAbs) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegCurvetoCubicAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

// Y1 prop
func (*SVGPathSegCurvetoCubicAbs) Y1() (y1 float32) {
	js.Rewrite("$<.y1")
	return y1
}

// Y1 prop
func (*SVGPathSegCurvetoCubicAbs) SetY1(y1 float32) {
	js.Rewrite("$<.y1 = $1", y1)
}

// Y2 prop
func (*SVGPathSegCurvetoCubicAbs) Y2() (y2 float32) {
	js.Rewrite("$<.y2")
	return y2
}

// Y2 prop
func (*SVGPathSegCurvetoCubicAbs) SetY2(y2 float32) {
	js.Rewrite("$<.y2 = $1", y2)
}
