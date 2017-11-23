package svgpathsegarcabs

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegArcAbs struct
// js:"SVGPathSegArcAbs,omit"
type SVGPathSegArcAbs struct {
	svgpathseg.SVGPathSeg
}

// Angle prop
func (*SVGPathSegArcAbs) Angle() (angle float32) {
	js.Rewrite("$<.angle")
	return angle
}

// Angle prop
func (*SVGPathSegArcAbs) SetAngle(angle float32) {
	js.Rewrite("$<.angle = $1", angle)
}

// LargeArcFlag prop
func (*SVGPathSegArcAbs) LargeArcFlag() (largeArcFlag bool) {
	js.Rewrite("$<.largeArcFlag")
	return largeArcFlag
}

// LargeArcFlag prop
func (*SVGPathSegArcAbs) SetLargeArcFlag(largeArcFlag bool) {
	js.Rewrite("$<.largeArcFlag = $1", largeArcFlag)
}

// R1 prop
func (*SVGPathSegArcAbs) R1() (r1 float32) {
	js.Rewrite("$<.r1")
	return r1
}

// R1 prop
func (*SVGPathSegArcAbs) SetR1(r1 float32) {
	js.Rewrite("$<.r1 = $1", r1)
}

// R2 prop
func (*SVGPathSegArcAbs) R2() (r2 float32) {
	js.Rewrite("$<.r2")
	return r2
}

// R2 prop
func (*SVGPathSegArcAbs) SetR2(r2 float32) {
	js.Rewrite("$<.r2 = $1", r2)
}

// SweepFlag prop
func (*SVGPathSegArcAbs) SweepFlag() (sweepFlag bool) {
	js.Rewrite("$<.sweepFlag")
	return sweepFlag
}

// SweepFlag prop
func (*SVGPathSegArcAbs) SetSweepFlag(sweepFlag bool) {
	js.Rewrite("$<.sweepFlag = $1", sweepFlag)
}

// X prop
func (*SVGPathSegArcAbs) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegArcAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// Y prop
func (*SVGPathSegArcAbs) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegArcAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
