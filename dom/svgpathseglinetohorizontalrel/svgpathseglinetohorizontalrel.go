package svgpathseglinetohorizontalrel

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegLinetoHorizontalRel)(nil)

// SVGPathSegLinetoHorizontalRel struct
// js:"SVGPathSegLinetoHorizontalRel,omit"
type SVGPathSegLinetoHorizontalRel struct {
}

// X prop
// js:"x"
func (*SVGPathSegLinetoHorizontalRel) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegLinetoHorizontalRel) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegLinetoHorizontalRel) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegLinetoHorizontalRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
