package svgpathsegmovetorel

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGPathSegMovetoRel,omit"
type SVGPathSegMovetoRel struct {
	svgpathseg.SVGPathSeg
}

// X
func (*SVGPathSegMovetoRel) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegMovetoRel) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// Y
func (*SVGPathSegMovetoRel) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegMovetoRel) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}
