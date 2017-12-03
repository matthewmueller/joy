package svgpathseglinetorel

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegLinetoRel)(nil)

// SVGPathSegLinetoRel struct
// js:"SVGPathSegLinetoRel,omit"
type SVGPathSegLinetoRel struct {
}

// X prop
// js:"x"
func (*SVGPathSegLinetoRel) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegLinetoRel) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

// Y prop
// js:"y"
func (*SVGPathSegLinetoRel) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegLinetoRel) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegLinetoRel) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegLinetoRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
