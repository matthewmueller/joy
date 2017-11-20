package svgpathsegcurvetocubicabs

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegCurvetoCubicAbs struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegCurvetoCubicAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoCubicAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoCubicAbs) GetX1() (x1 float32) {
	js.Rewrite("$<.x1")
	return x1
}

func (*SVGPathSegCurvetoCubicAbs) SetX1(x1 float32) {
	js.Rewrite("$<.x1 = $1", x1)
}

func (*SVGPathSegCurvetoCubicAbs) GetX2() (x2 float32) {
	js.Rewrite("$<.x2")
	return x2
}

func (*SVGPathSegCurvetoCubicAbs) SetX2(x2 float32) {
	js.Rewrite("$<.x2 = $1", x2)
}

func (*SVGPathSegCurvetoCubicAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoCubicAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

func (*SVGPathSegCurvetoCubicAbs) GetY1() (y1 float32) {
	js.Rewrite("$<.y1")
	return y1
}

func (*SVGPathSegCurvetoCubicAbs) SetY1(y1 float32) {
	js.Rewrite("$<.y1 = $1", y1)
}

func (*SVGPathSegCurvetoCubicAbs) GetY2() (y2 float32) {
	js.Rewrite("$<.y2")
	return y2
}

func (*SVGPathSegCurvetoCubicAbs) SetY2(y2 float32) {
	js.Rewrite("$<.y2 = $1", y2)
}
