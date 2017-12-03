package svgpathseglinetoabs

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegLinetoAbs)(nil)

// SVGPathSegLinetoAbs struct
// js:"SVGPathSegLinetoAbs,omit"
type SVGPathSegLinetoAbs struct {
}

// X prop
// js:"x"
func (*SVGPathSegLinetoAbs) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegLinetoAbs) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

// Y prop
// js:"y"
func (*SVGPathSegLinetoAbs) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegLinetoAbs) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegLinetoAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegLinetoAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
