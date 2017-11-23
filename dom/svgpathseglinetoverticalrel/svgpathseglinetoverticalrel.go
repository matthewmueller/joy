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

// Y prop
func (*SVGPathSegLinetoVerticalRel) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegLinetoVerticalRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
