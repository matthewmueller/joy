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

// Angle
func (*SVGPathSegArcAbs) Angle() (angle float32) {
	js.Rewrite("$<.Angle")
	return angle
}

// Angle
func (*SVGPathSegArcAbs) SetAngle(angle float32) {
	js.Rewrite("$<.Angle = $1", angle)
}

// LargeArcFlag
func (*SVGPathSegArcAbs) LargeArcFlag() (largeArcFlag bool) {
	js.Rewrite("$<.LargeArcFlag")
	return largeArcFlag
}

// LargeArcFlag
func (*SVGPathSegArcAbs) SetLargeArcFlag(largeArcFlag bool) {
	js.Rewrite("$<.LargeArcFlag = $1", largeArcFlag)
}

// R1
func (*SVGPathSegArcAbs) R1() (r1 float32) {
	js.Rewrite("$<.R1")
	return r1
}

// R1
func (*SVGPathSegArcAbs) SetR1(r1 float32) {
	js.Rewrite("$<.R1 = $1", r1)
}

// R2
func (*SVGPathSegArcAbs) R2() (r2 float32) {
	js.Rewrite("$<.R2")
	return r2
}

// R2
func (*SVGPathSegArcAbs) SetR2(r2 float32) {
	js.Rewrite("$<.R2 = $1", r2)
}

// SweepFlag
func (*SVGPathSegArcAbs) SweepFlag() (sweepFlag bool) {
	js.Rewrite("$<.SweepFlag")
	return sweepFlag
}

// SweepFlag
func (*SVGPathSegArcAbs) SetSweepFlag(sweepFlag bool) {
	js.Rewrite("$<.SweepFlag = $1", sweepFlag)
}

// X
func (*SVGPathSegArcAbs) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegArcAbs) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// Y
func (*SVGPathSegArcAbs) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegArcAbs) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}
