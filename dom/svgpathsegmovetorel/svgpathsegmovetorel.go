package svgpathsegmovetorel

import (
	"github.com/matthewmueller/golly/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegMovetoRel struct
// js:"SVGPathSegMovetoRel,omit"
type SVGPathSegMovetoRel struct {
	svgpathseg.SVGPathSeg
}

// X prop
func (*SVGPathSegMovetoRel) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegMovetoRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// Y prop
func (*SVGPathSegMovetoRel) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegMovetoRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
