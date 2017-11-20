package svgpathsegcurvetocubicrel

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegCurvetoCubicRel struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegCurvetoCubicRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoCubicRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoCubicRel) GetX1() (x1 float32) {
	js.Rewrite("$<.x1")
	return x1
}

func (*SVGPathSegCurvetoCubicRel) SetX1(x1 float32) {
	js.Rewrite("$<.x1 = $1", x1)
}

func (*SVGPathSegCurvetoCubicRel) GetX2() (x2 float32) {
	js.Rewrite("$<.x2")
	return x2
}

func (*SVGPathSegCurvetoCubicRel) SetX2(x2 float32) {
	js.Rewrite("$<.x2 = $1", x2)
}

func (*SVGPathSegCurvetoCubicRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoCubicRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

func (*SVGPathSegCurvetoCubicRel) GetY1() (y1 float32) {
	js.Rewrite("$<.y1")
	return y1
}

func (*SVGPathSegCurvetoCubicRel) SetY1(y1 float32) {
	js.Rewrite("$<.y1 = $1", y1)
}

func (*SVGPathSegCurvetoCubicRel) GetY2() (y2 float32) {
	js.Rewrite("$<.y2")
	return y2
}

func (*SVGPathSegCurvetoCubicRel) SetY2(y2 float32) {
	js.Rewrite("$<.y2 = $1", y2)
}
