package svgpathsegcurvetocubicrel

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegCurvetoCubicRel struct
// js:"SVGPathSegCurvetoCubicRel,omit"
type SVGPathSegCurvetoCubicRel struct {
	svgpathseg.SVGPathSeg
}

// X
func (*SVGPathSegCurvetoCubicRel) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegCurvetoCubicRel) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// X1
func (*SVGPathSegCurvetoCubicRel) X1() (x1 float32) {
	js.Rewrite("$<.X1")
	return x1
}

// X1
func (*SVGPathSegCurvetoCubicRel) SetX1(x1 float32) {
	js.Rewrite("$<.X1 = $1", x1)
}

// X2
func (*SVGPathSegCurvetoCubicRel) X2() (x2 float32) {
	js.Rewrite("$<.X2")
	return x2
}

// X2
func (*SVGPathSegCurvetoCubicRel) SetX2(x2 float32) {
	js.Rewrite("$<.X2 = $1", x2)
}

// Y
func (*SVGPathSegCurvetoCubicRel) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegCurvetoCubicRel) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}

// Y1
func (*SVGPathSegCurvetoCubicRel) Y1() (y1 float32) {
	js.Rewrite("$<.Y1")
	return y1
}

// Y1
func (*SVGPathSegCurvetoCubicRel) SetY1(y1 float32) {
	js.Rewrite("$<.Y1 = $1", y1)
}

// Y2
func (*SVGPathSegCurvetoCubicRel) Y2() (y2 float32) {
	js.Rewrite("$<.Y2")
	return y2
}

// Y2
func (*SVGPathSegCurvetoCubicRel) SetY2(y2 float32) {
	js.Rewrite("$<.Y2 = $1", y2)
}
