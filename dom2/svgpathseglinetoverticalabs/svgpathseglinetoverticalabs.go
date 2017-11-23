package svgpathseglinetoverticalabs

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGPathSegLinetoVerticalAbs,omit"
type SVGPathSegLinetoVerticalAbs struct {
	svgpathseg.SVGPathSeg
}

// Y
func (*SVGPathSegLinetoVerticalAbs) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegLinetoVerticalAbs) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}
