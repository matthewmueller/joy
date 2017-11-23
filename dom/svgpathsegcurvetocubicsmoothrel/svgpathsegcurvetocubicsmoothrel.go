package svgpathsegcurvetocubicsmoothrel

import (
	"github.com/matthewmueller/golly/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegCurvetoCubicSmoothRel struct
// js:"SVGPathSegCurvetoCubicSmoothRel,omit"
type SVGPathSegCurvetoCubicSmoothRel struct {
	svgpathseg.SVGPathSeg
}

// X prop
func (*SVGPathSegCurvetoCubicSmoothRel) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegCurvetoCubicSmoothRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// X2 prop
func (*SVGPathSegCurvetoCubicSmoothRel) X2() (x2 float32) {
	js.Rewrite("$<.x2")
	return x2
}

// X2 prop
func (*SVGPathSegCurvetoCubicSmoothRel) SetX2(x2 float32) {
	js.Rewrite("$<.x2 = $1", x2)
}

// Y prop
func (*SVGPathSegCurvetoCubicSmoothRel) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegCurvetoCubicSmoothRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

// Y2 prop
func (*SVGPathSegCurvetoCubicSmoothRel) Y2() (y2 float32) {
	js.Rewrite("$<.y2")
	return y2
}

// Y2 prop
func (*SVGPathSegCurvetoCubicSmoothRel) SetY2(y2 float32) {
	js.Rewrite("$<.y2 = $1", y2)
}
