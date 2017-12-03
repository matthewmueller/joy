package svgpathsegcurvetocubicsmoothabs

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegCurvetoCubicSmoothAbs)(nil)

// SVGPathSegCurvetoCubicSmoothAbs struct
// js:"SVGPathSegCurvetoCubicSmoothAbs,omit"
type SVGPathSegCurvetoCubicSmoothAbs struct {
}

// X prop
// js:"x"
func (*SVGPathSegCurvetoCubicSmoothAbs) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegCurvetoCubicSmoothAbs) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

// X2 prop
// js:"x2"
func (*SVGPathSegCurvetoCubicSmoothAbs) X2() (x2 float32) {
	macro.Rewrite("$_.x2")
	return x2
}

// SetX2 prop
// js:"x2"
func (*SVGPathSegCurvetoCubicSmoothAbs) SetX2(x2 float32) {
	macro.Rewrite("$_.x2 = $1", x2)
}

// Y prop
// js:"y"
func (*SVGPathSegCurvetoCubicSmoothAbs) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegCurvetoCubicSmoothAbs) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

// Y2 prop
// js:"y2"
func (*SVGPathSegCurvetoCubicSmoothAbs) Y2() (y2 float32) {
	macro.Rewrite("$_.y2")
	return y2
}

// SetY2 prop
// js:"y2"
func (*SVGPathSegCurvetoCubicSmoothAbs) SetY2(y2 float32) {
	macro.Rewrite("$_.y2 = $1", y2)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegCurvetoCubicSmoothAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegCurvetoCubicSmoothAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
