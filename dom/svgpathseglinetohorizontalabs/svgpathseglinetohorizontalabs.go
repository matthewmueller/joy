package svgpathseglinetohorizontalabs

import (
	"github.com/matthewmueller/golly/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegLinetoHorizontalAbs struct
// js:"SVGPathSegLinetoHorizontalAbs,omit"
type SVGPathSegLinetoHorizontalAbs struct {
	svgpathseg.SVGPathSeg
}

// X prop
func (*SVGPathSegLinetoHorizontalAbs) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegLinetoHorizontalAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}
