package svgpathseglinetorel

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegLinetoRel struct
// js:"SVGPathSegLinetoRel,omit"
type SVGPathSegLinetoRel struct {
	svgpathseg.SVGPathSeg
}

// X
func (*SVGPathSegLinetoRel) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegLinetoRel) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// Y
func (*SVGPathSegLinetoRel) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegLinetoRel) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}
