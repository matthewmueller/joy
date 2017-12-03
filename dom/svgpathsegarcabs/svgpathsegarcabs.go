package svgpathsegarcabs

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegArcAbs)(nil)

// SVGPathSegArcAbs struct
// js:"SVGPathSegArcAbs,omit"
type SVGPathSegArcAbs struct {
}

// Angle prop
// js:"angle"
func (*SVGPathSegArcAbs) Angle() (angle float32) {
	macro.Rewrite("$_.angle")
	return angle
}

// SetAngle prop
// js:"angle"
func (*SVGPathSegArcAbs) SetAngle(angle float32) {
	macro.Rewrite("$_.angle = $1", angle)
}

// LargeArcFlag prop
// js:"largeArcFlag"
func (*SVGPathSegArcAbs) LargeArcFlag() (largeArcFlag bool) {
	macro.Rewrite("$_.largeArcFlag")
	return largeArcFlag
}

// SetLargeArcFlag prop
// js:"largeArcFlag"
func (*SVGPathSegArcAbs) SetLargeArcFlag(largeArcFlag bool) {
	macro.Rewrite("$_.largeArcFlag = $1", largeArcFlag)
}

// R1 prop
// js:"r1"
func (*SVGPathSegArcAbs) R1() (r1 float32) {
	macro.Rewrite("$_.r1")
	return r1
}

// SetR1 prop
// js:"r1"
func (*SVGPathSegArcAbs) SetR1(r1 float32) {
	macro.Rewrite("$_.r1 = $1", r1)
}

// R2 prop
// js:"r2"
func (*SVGPathSegArcAbs) R2() (r2 float32) {
	macro.Rewrite("$_.r2")
	return r2
}

// SetR2 prop
// js:"r2"
func (*SVGPathSegArcAbs) SetR2(r2 float32) {
	macro.Rewrite("$_.r2 = $1", r2)
}

// SweepFlag prop
// js:"sweepFlag"
func (*SVGPathSegArcAbs) SweepFlag() (sweepFlag bool) {
	macro.Rewrite("$_.sweepFlag")
	return sweepFlag
}

// SetSweepFlag prop
// js:"sweepFlag"
func (*SVGPathSegArcAbs) SetSweepFlag(sweepFlag bool) {
	macro.Rewrite("$_.sweepFlag = $1", sweepFlag)
}

// X prop
// js:"x"
func (*SVGPathSegArcAbs) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegArcAbs) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

// Y prop
// js:"y"
func (*SVGPathSegArcAbs) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegArcAbs) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegArcAbs) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegArcAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
