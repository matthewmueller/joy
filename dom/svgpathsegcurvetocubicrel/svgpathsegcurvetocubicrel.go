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

// X prop
func (*SVGPathSegCurvetoCubicRel) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegCurvetoCubicRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// X1 prop
func (*SVGPathSegCurvetoCubicRel) X1() (x1 float32) {
	js.Rewrite("$<.x1")
	return x1
}

// X1 prop
func (*SVGPathSegCurvetoCubicRel) SetX1(x1 float32) {
	js.Rewrite("$<.x1 = $1", x1)
}

// X2 prop
func (*SVGPathSegCurvetoCubicRel) X2() (x2 float32) {
	js.Rewrite("$<.x2")
	return x2
}

// X2 prop
func (*SVGPathSegCurvetoCubicRel) SetX2(x2 float32) {
	js.Rewrite("$<.x2 = $1", x2)
}

// Y prop
func (*SVGPathSegCurvetoCubicRel) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegCurvetoCubicRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

// Y1 prop
func (*SVGPathSegCurvetoCubicRel) Y1() (y1 float32) {
	js.Rewrite("$<.y1")
	return y1
}

// Y1 prop
func (*SVGPathSegCurvetoCubicRel) SetY1(y1 float32) {
	js.Rewrite("$<.y1 = $1", y1)
}

// Y2 prop
func (*SVGPathSegCurvetoCubicRel) Y2() (y2 float32) {
	js.Rewrite("$<.y2")
	return y2
}

// Y2 prop
func (*SVGPathSegCurvetoCubicRel) SetY2(y2 float32) {
	js.Rewrite("$<.y2 = $1", y2)
}
