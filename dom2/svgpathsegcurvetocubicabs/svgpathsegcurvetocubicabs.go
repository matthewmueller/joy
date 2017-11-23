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

// X
func (*SVGPathSegCurvetoCubicAbs) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegCurvetoCubicAbs) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// X1
func (*SVGPathSegCurvetoCubicAbs) X1() (x1 float32) {
	js.Rewrite("$<.X1")
	return x1
}

// X1
func (*SVGPathSegCurvetoCubicAbs) SetX1(x1 float32) {
	js.Rewrite("$<.X1 = $1", x1)
}

// X2
func (*SVGPathSegCurvetoCubicAbs) X2() (x2 float32) {
	js.Rewrite("$<.X2")
	return x2
}

// X2
func (*SVGPathSegCurvetoCubicAbs) SetX2(x2 float32) {
	js.Rewrite("$<.X2 = $1", x2)
}

// Y
func (*SVGPathSegCurvetoCubicAbs) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegCurvetoCubicAbs) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}

// Y1
func (*SVGPathSegCurvetoCubicAbs) Y1() (y1 float32) {
	js.Rewrite("$<.Y1")
	return y1
}

// Y1
func (*SVGPathSegCurvetoCubicAbs) SetY1(y1 float32) {
	js.Rewrite("$<.Y1 = $1", y1)
}

// Y2
func (*SVGPathSegCurvetoCubicAbs) Y2() (y2 float32) {
	js.Rewrite("$<.Y2")
	return y2
}

// Y2
func (*SVGPathSegCurvetoCubicAbs) SetY2(y2 float32) {
	js.Rewrite("$<.Y2 = $1", y2)
}
