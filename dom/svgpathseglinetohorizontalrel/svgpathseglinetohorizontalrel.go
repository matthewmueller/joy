package svgpathseglinetohorizontalrel

import (
	"github.com/matthewmueller/golly/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegLinetoHorizontalRel struct
// js:"SVGPathSegLinetoHorizontalRel,omit"
type SVGPathSegLinetoHorizontalRel struct {
	svgpathseg.SVGPathSeg
}

// X prop
func (*SVGPathSegLinetoHorizontalRel) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegLinetoHorizontalRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}
