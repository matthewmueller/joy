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

// X prop
func (*SVGPathSegLinetoRel) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegLinetoRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// Y prop
func (*SVGPathSegLinetoRel) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegLinetoRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
