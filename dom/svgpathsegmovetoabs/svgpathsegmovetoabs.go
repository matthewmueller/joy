package svgpathsegmovetoabs

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegMovetoAbs struct
// js:"SVGPathSegMovetoAbs,omit"
type SVGPathSegMovetoAbs struct {
	svgpathseg.SVGPathSeg
}

// X prop
func (*SVGPathSegMovetoAbs) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegMovetoAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// Y prop
func (*SVGPathSegMovetoAbs) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegMovetoAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
