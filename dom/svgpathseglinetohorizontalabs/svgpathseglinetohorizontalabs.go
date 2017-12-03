package svgpathseglinetohorizontalabs

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegLinetoHorizontalAbs)(nil)

// SVGPathSegLinetoHorizontalAbs struct
// js:"SVGPathSegLinetoHorizontalAbs,omit"
type SVGPathSegLinetoHorizontalAbs struct {
}

// X prop
// js:"x"
func (*SVGPathSegLinetoHorizontalAbs) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegLinetoHorizontalAbs) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegLinetoHorizontalAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegLinetoHorizontalAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
