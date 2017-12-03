package svgpathseglinetohorizontalabs

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/js"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegLinetoHorizontalAbs)(nil)

// SVGPathSegLinetoHorizontalAbs struct
// js:"SVGPathSegLinetoHorizontalAbs,omit"
type SVGPathSegLinetoHorizontalAbs struct {
}

// X prop
// js:"x"
func (*SVGPathSegLinetoHorizontalAbs) X() (x float32) {
	js.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegLinetoHorizontalAbs) SetX(x float32) {
	js.Rewrite("$_.x = $1", x)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegLinetoHorizontalAbs) PathSegType() (pathSegType uint8) {
	js.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegLinetoHorizontalAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	js.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
