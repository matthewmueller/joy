package svgpathsegcurvetocubicrel

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegCurvetoCubicRel)(nil)

// SVGPathSegCurvetoCubicRel struct
// js:"SVGPathSegCurvetoCubicRel,omit"
type SVGPathSegCurvetoCubicRel struct {
}

// X prop
// js:"x"
func (*SVGPathSegCurvetoCubicRel) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegCurvetoCubicRel) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

// X1 prop
// js:"x1"
func (*SVGPathSegCurvetoCubicRel) X1() (x1 float32) {
	macro.Rewrite("$_.x1")
	return x1
}

// SetX1 prop
// js:"x1"
func (*SVGPathSegCurvetoCubicRel) SetX1(x1 float32) {
	macro.Rewrite("$_.x1 = $1", x1)
}

// X2 prop
// js:"x2"
func (*SVGPathSegCurvetoCubicRel) X2() (x2 float32) {
	macro.Rewrite("$_.x2")
	return x2
}

// SetX2 prop
// js:"x2"
func (*SVGPathSegCurvetoCubicRel) SetX2(x2 float32) {
	macro.Rewrite("$_.x2 = $1", x2)
}

// Y prop
// js:"y"
func (*SVGPathSegCurvetoCubicRel) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegCurvetoCubicRel) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

// Y1 prop
// js:"y1"
func (*SVGPathSegCurvetoCubicRel) Y1() (y1 float32) {
	macro.Rewrite("$_.y1")
	return y1
}

// SetY1 prop
// js:"y1"
func (*SVGPathSegCurvetoCubicRel) SetY1(y1 float32) {
	macro.Rewrite("$_.y1 = $1", y1)
}

// Y2 prop
// js:"y2"
func (*SVGPathSegCurvetoCubicRel) Y2() (y2 float32) {
	macro.Rewrite("$_.y2")
	return y2
}

// SetY2 prop
// js:"y2"
func (*SVGPathSegCurvetoCubicRel) SetY2(y2 float32) {
	macro.Rewrite("$_.y2 = $1", y2)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegCurvetoCubicRel) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegCurvetoCubicRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
