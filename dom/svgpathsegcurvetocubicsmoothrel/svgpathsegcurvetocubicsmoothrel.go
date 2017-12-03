package svgpathsegcurvetocubicsmoothrel

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegCurvetoCubicSmoothRel)(nil)

// SVGPathSegCurvetoCubicSmoothRel struct
// js:"SVGPathSegCurvetoCubicSmoothRel,omit"
type SVGPathSegCurvetoCubicSmoothRel struct {
}

// X prop
// js:"x"
func (*SVGPathSegCurvetoCubicSmoothRel) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegCurvetoCubicSmoothRel) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

// X2 prop
// js:"x2"
func (*SVGPathSegCurvetoCubicSmoothRel) X2() (x2 float32) {
	macro.Rewrite("$_.x2")
	return x2
}

// SetX2 prop
// js:"x2"
func (*SVGPathSegCurvetoCubicSmoothRel) SetX2(x2 float32) {
	macro.Rewrite("$_.x2 = $1", x2)
}

// Y prop
// js:"y"
func (*SVGPathSegCurvetoCubicSmoothRel) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegCurvetoCubicSmoothRel) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

// Y2 prop
// js:"y2"
func (*SVGPathSegCurvetoCubicSmoothRel) Y2() (y2 float32) {
	macro.Rewrite("$_.y2")
	return y2
}

// SetY2 prop
// js:"y2"
func (*SVGPathSegCurvetoCubicSmoothRel) SetY2(y2 float32) {
	macro.Rewrite("$_.y2 = $1", y2)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegCurvetoCubicSmoothRel) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegCurvetoCubicSmoothRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
