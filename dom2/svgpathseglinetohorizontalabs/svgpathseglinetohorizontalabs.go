package svgpathseglinetohorizontalabs

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegLinetoHorizontalAbs struct
// js:"SVGPathSegLinetoHorizontalAbs,omit"
type SVGPathSegLinetoHorizontalAbs struct {
	svgpathseg.SVGPathSeg
}

// X
func (*SVGPathSegLinetoHorizontalAbs) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegLinetoHorizontalAbs) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}
