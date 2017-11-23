package svgpathsegarcrel

import (
	"github.com/matthewmueller/golly/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegArcRel struct
// js:"SVGPathSegArcRel,omit"
type SVGPathSegArcRel struct {
	svgpathseg.SVGPathSeg
}

// Angle prop
func (*SVGPathSegArcRel) Angle() (angle float32) {
	js.Rewrite("$<.angle")
	return angle
}

// Angle prop
func (*SVGPathSegArcRel) SetAngle(angle float32) {
	js.Rewrite("$<.angle = $1", angle)
}

// LargeArcFlag prop
func (*SVGPathSegArcRel) LargeArcFlag() (largeArcFlag bool) {
	js.Rewrite("$<.largeArcFlag")
	return largeArcFlag
}

// LargeArcFlag prop
func (*SVGPathSegArcRel) SetLargeArcFlag(largeArcFlag bool) {
	js.Rewrite("$<.largeArcFlag = $1", largeArcFlag)
}

// R1 prop
func (*SVGPathSegArcRel) R1() (r1 float32) {
	js.Rewrite("$<.r1")
	return r1
}

// R1 prop
func (*SVGPathSegArcRel) SetR1(r1 float32) {
	js.Rewrite("$<.r1 = $1", r1)
}

// R2 prop
func (*SVGPathSegArcRel) R2() (r2 float32) {
	js.Rewrite("$<.r2")
	return r2
}

// R2 prop
func (*SVGPathSegArcRel) SetR2(r2 float32) {
	js.Rewrite("$<.r2 = $1", r2)
}

// SweepFlag prop
func (*SVGPathSegArcRel) SweepFlag() (sweepFlag bool) {
	js.Rewrite("$<.sweepFlag")
	return sweepFlag
}

// SweepFlag prop
func (*SVGPathSegArcRel) SetSweepFlag(sweepFlag bool) {
	js.Rewrite("$<.sweepFlag = $1", sweepFlag)
}

// X prop
func (*SVGPathSegArcRel) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGPathSegArcRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// Y prop
func (*SVGPathSegArcRel) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGPathSegArcRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
