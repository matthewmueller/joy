package svgpathsegmovetoabs

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGPathSegMovetoAbs,omit"
type SVGPathSegMovetoAbs struct {
	svgpathseg.SVGPathSeg
}

// X
func (*SVGPathSegMovetoAbs) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegMovetoAbs) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// Y
func (*SVGPathSegMovetoAbs) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegMovetoAbs) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}
