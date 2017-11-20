package svgpathsegarcrel

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegArcRel struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegArcRel) GetAngle() (angle float32) {
	js.Rewrite("$<.angle")
	return angle
}

func (*SVGPathSegArcRel) SetAngle(angle float32) {
	js.Rewrite("$<.angle = $1", angle)
}

func (*SVGPathSegArcRel) GetLargeArcFlag() (largeArcFlag bool) {
	js.Rewrite("$<.largeArcFlag")
	return largeArcFlag
}

func (*SVGPathSegArcRel) SetLargeArcFlag(largeArcFlag bool) {
	js.Rewrite("$<.largeArcFlag = $1", largeArcFlag)
}

func (*SVGPathSegArcRel) GetR1() (r1 float32) {
	js.Rewrite("$<.r1")
	return r1
}

func (*SVGPathSegArcRel) SetR1(r1 float32) {
	js.Rewrite("$<.r1 = $1", r1)
}

func (*SVGPathSegArcRel) GetR2() (r2 float32) {
	js.Rewrite("$<.r2")
	return r2
}

func (*SVGPathSegArcRel) SetR2(r2 float32) {
	js.Rewrite("$<.r2 = $1", r2)
}

func (*SVGPathSegArcRel) GetSweepFlag() (sweepFlag bool) {
	js.Rewrite("$<.sweepFlag")
	return sweepFlag
}

func (*SVGPathSegArcRel) SetSweepFlag(sweepFlag bool) {
	js.Rewrite("$<.sweepFlag = $1", sweepFlag)
}

func (*SVGPathSegArcRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegArcRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegArcRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegArcRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
