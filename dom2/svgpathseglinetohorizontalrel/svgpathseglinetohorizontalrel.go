package svgpathseglinetohorizontalrel

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegLinetoHorizontalRel struct
// js:"SVGPathSegLinetoHorizontalRel,omit"
type SVGPathSegLinetoHorizontalRel struct {
	svgpathseg.SVGPathSeg
}

// X
func (*SVGPathSegLinetoHorizontalRel) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegLinetoHorizontalRel) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}
