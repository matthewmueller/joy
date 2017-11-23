package svgpathseglinetoverticalrel

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegLinetoVerticalRel struct
// js:"SVGPathSegLinetoVerticalRel,omit"
type SVGPathSegLinetoVerticalRel struct {
	svgpathseg.SVGPathSeg
}

// Y
func (*SVGPathSegLinetoVerticalRel) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegLinetoVerticalRel) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}
