package svgpathsegarcrel

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGPathSegArcRel,omit"
type SVGPathSegArcRel struct {
	svgpathseg.SVGPathSeg
}

// Angle
func (*SVGPathSegArcRel) Angle() (angle float32) {
	js.Rewrite("$<.Angle")
	return angle
}

// Angle
func (*SVGPathSegArcRel) SetAngle(angle float32) {
	js.Rewrite("$<.Angle = $1", angle)
}

// LargeArcFlag
func (*SVGPathSegArcRel) LargeArcFlag() (largeArcFlag bool) {
	js.Rewrite("$<.LargeArcFlag")
	return largeArcFlag
}

// LargeArcFlag
func (*SVGPathSegArcRel) SetLargeArcFlag(largeArcFlag bool) {
	js.Rewrite("$<.LargeArcFlag = $1", largeArcFlag)
}

// R1
func (*SVGPathSegArcRel) R1() (r1 float32) {
	js.Rewrite("$<.R1")
	return r1
}

// R1
func (*SVGPathSegArcRel) SetR1(r1 float32) {
	js.Rewrite("$<.R1 = $1", r1)
}

// R2
func (*SVGPathSegArcRel) R2() (r2 float32) {
	js.Rewrite("$<.R2")
	return r2
}

// R2
func (*SVGPathSegArcRel) SetR2(r2 float32) {
	js.Rewrite("$<.R2 = $1", r2)
}

// SweepFlag
func (*SVGPathSegArcRel) SweepFlag() (sweepFlag bool) {
	js.Rewrite("$<.SweepFlag")
	return sweepFlag
}

// SweepFlag
func (*SVGPathSegArcRel) SetSweepFlag(sweepFlag bool) {
	js.Rewrite("$<.SweepFlag = $1", sweepFlag)
}

// X
func (*SVGPathSegArcRel) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegArcRel) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// Y
func (*SVGPathSegArcRel) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegArcRel) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}
