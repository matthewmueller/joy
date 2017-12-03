package svgpathsegarcrel

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/js"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegArcRel)(nil)

// SVGPathSegArcRel struct
// js:"SVGPathSegArcRel,omit"
type SVGPathSegArcRel struct {
}

// Angle prop
// js:"angle"
func (*SVGPathSegArcRel) Angle() (angle float32) {
	js.Rewrite("$_.angle")
	return angle
}

// SetAngle prop
// js:"angle"
func (*SVGPathSegArcRel) SetAngle(angle float32) {
	js.Rewrite("$_.angle = $1", angle)
}

// LargeArcFlag prop
// js:"largeArcFlag"
func (*SVGPathSegArcRel) LargeArcFlag() (largeArcFlag bool) {
	js.Rewrite("$_.largeArcFlag")
	return largeArcFlag
}

// SetLargeArcFlag prop
// js:"largeArcFlag"
func (*SVGPathSegArcRel) SetLargeArcFlag(largeArcFlag bool) {
	js.Rewrite("$_.largeArcFlag = $1", largeArcFlag)
}

// R1 prop
// js:"r1"
func (*SVGPathSegArcRel) R1() (r1 float32) {
	js.Rewrite("$_.r1")
	return r1
}

// SetR1 prop
// js:"r1"
func (*SVGPathSegArcRel) SetR1(r1 float32) {
	js.Rewrite("$_.r1 = $1", r1)
}

// R2 prop
// js:"r2"
func (*SVGPathSegArcRel) R2() (r2 float32) {
	js.Rewrite("$_.r2")
	return r2
}

// SetR2 prop
// js:"r2"
func (*SVGPathSegArcRel) SetR2(r2 float32) {
	js.Rewrite("$_.r2 = $1", r2)
}

// SweepFlag prop
// js:"sweepFlag"
func (*SVGPathSegArcRel) SweepFlag() (sweepFlag bool) {
	js.Rewrite("$_.sweepFlag")
	return sweepFlag
}

// SetSweepFlag prop
// js:"sweepFlag"
func (*SVGPathSegArcRel) SetSweepFlag(sweepFlag bool) {
	js.Rewrite("$_.sweepFlag = $1", sweepFlag)
}

// X prop
// js:"x"
func (*SVGPathSegArcRel) X() (x float32) {
	js.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegArcRel) SetX(x float32) {
	js.Rewrite("$_.x = $1", x)
}

// Y prop
// js:"y"
func (*SVGPathSegArcRel) Y() (y float32) {
	js.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegArcRel) SetY(y float32) {
	js.Rewrite("$_.y = $1", y)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegArcRel) PathSegType() (pathSegType uint8) {
	js.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegArcRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	js.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
