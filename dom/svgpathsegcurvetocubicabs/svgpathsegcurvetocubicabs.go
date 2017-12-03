package svgpathsegcurvetocubicabs

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegCurvetoCubicAbs)(nil)

// SVGPathSegCurvetoCubicAbs struct
// js:"SVGPathSegCurvetoCubicAbs,omit"
type SVGPathSegCurvetoCubicAbs struct {
}

// X prop
// js:"x"
func (*SVGPathSegCurvetoCubicAbs) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegCurvetoCubicAbs) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

// X1 prop
// js:"x1"
func (*SVGPathSegCurvetoCubicAbs) X1() (x1 float32) {
	macro.Rewrite("$_.x1")
	return x1
}

// SetX1 prop
// js:"x1"
func (*SVGPathSegCurvetoCubicAbs) SetX1(x1 float32) {
	macro.Rewrite("$_.x1 = $1", x1)
}

// X2 prop
// js:"x2"
func (*SVGPathSegCurvetoCubicAbs) X2() (x2 float32) {
	macro.Rewrite("$_.x2")
	return x2
}

// SetX2 prop
// js:"x2"
func (*SVGPathSegCurvetoCubicAbs) SetX2(x2 float32) {
	macro.Rewrite("$_.x2 = $1", x2)
}

// Y prop
// js:"y"
func (*SVGPathSegCurvetoCubicAbs) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegCurvetoCubicAbs) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

// Y1 prop
// js:"y1"
func (*SVGPathSegCurvetoCubicAbs) Y1() (y1 float32) {
	macro.Rewrite("$_.y1")
	return y1
}

// SetY1 prop
// js:"y1"
func (*SVGPathSegCurvetoCubicAbs) SetY1(y1 float32) {
	macro.Rewrite("$_.y1 = $1", y1)
}

// Y2 prop
// js:"y2"
func (*SVGPathSegCurvetoCubicAbs) Y2() (y2 float32) {
	macro.Rewrite("$_.y2")
	return y2
}

// SetY2 prop
// js:"y2"
func (*SVGPathSegCurvetoCubicAbs) SetY2(y2 float32) {
	macro.Rewrite("$_.y2 = $1", y2)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegCurvetoCubicAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegCurvetoCubicAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
