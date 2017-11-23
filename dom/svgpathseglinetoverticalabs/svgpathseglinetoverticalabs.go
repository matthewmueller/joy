package svgpathseglinetoverticalabs

import (
	"github.com/matthewmueller/golly/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegLinetoVerticalAbs struct
// js:"SVGPathSegLinetoVerticalAbs,omit"
type SVGPathSegLinetoVerticalAbs struct {
	svgpathseg.SVGPathSeg
}

// Y prop
func (*SVGPathSegLinetoVerticalAbs) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegLinetoVerticalAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
