package svgpathseglinetoabs

import (
	"github.com/matthewmueller/golly/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegLinetoAbs struct
// js:"SVGPathSegLinetoAbs,omit"
type SVGPathSegLinetoAbs struct {
	svgpathseg.SVGPathSeg
}

// X prop
func (*SVGPathSegLinetoAbs) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegLinetoAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// Y prop
func (*SVGPathSegLinetoAbs) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegLinetoAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
