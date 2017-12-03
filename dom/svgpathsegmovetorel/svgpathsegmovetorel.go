package svgpathsegmovetorel

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegMovetoRel)(nil)

// SVGPathSegMovetoRel struct
// js:"SVGPathSegMovetoRel,omit"
type SVGPathSegMovetoRel struct {
}

// X prop
// js:"x"
func (*SVGPathSegMovetoRel) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegMovetoRel) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

// Y prop
// js:"y"
func (*SVGPathSegMovetoRel) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegMovetoRel) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegMovetoRel) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegMovetoRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
