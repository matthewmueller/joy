package svgpathsegarcabs

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegArcAbs struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegArcAbs) GetAngle() (angle float32) {
	js.Rewrite("$<.angle")
	return angle
}

func (*SVGPathSegArcAbs) SetAngle(angle float32) {
	js.Rewrite("$<.angle = $1", angle)
}

func (*SVGPathSegArcAbs) GetLargeArcFlag() (largeArcFlag bool) {
	js.Rewrite("$<.largeArcFlag")
	return largeArcFlag
}

func (*SVGPathSegArcAbs) SetLargeArcFlag(largeArcFlag bool) {
	js.Rewrite("$<.largeArcFlag = $1", largeArcFlag)
}

func (*SVGPathSegArcAbs) GetR1() (r1 float32) {
	js.Rewrite("$<.r1")
	return r1
}

func (*SVGPathSegArcAbs) SetR1(r1 float32) {
	js.Rewrite("$<.r1 = $1", r1)
}

func (*SVGPathSegArcAbs) GetR2() (r2 float32) {
	js.Rewrite("$<.r2")
	return r2
}

func (*SVGPathSegArcAbs) SetR2(r2 float32) {
	js.Rewrite("$<.r2 = $1", r2)
}

func (*SVGPathSegArcAbs) GetSweepFlag() (sweepFlag bool) {
	js.Rewrite("$<.sweepFlag")
	return sweepFlag
}

func (*SVGPathSegArcAbs) SetSweepFlag(sweepFlag bool) {
	js.Rewrite("$<.sweepFlag = $1", sweepFlag)
}

func (*SVGPathSegArcAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegArcAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegArcAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegArcAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
