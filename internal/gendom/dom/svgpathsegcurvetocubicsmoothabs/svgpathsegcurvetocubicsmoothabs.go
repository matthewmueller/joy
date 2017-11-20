package svgpathsegcurvetocubicsmoothabs

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegCurvetoCubicSmoothAbs struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegCurvetoCubicSmoothAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoCubicSmoothAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoCubicSmoothAbs) GetX2() (x2 float32) {
	js.Rewrite("$<.x2")
	return x2
}

func (*SVGPathSegCurvetoCubicSmoothAbs) SetX2(x2 float32) {
	js.Rewrite("$<.x2 = $1", x2)
}

func (*SVGPathSegCurvetoCubicSmoothAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoCubicSmoothAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

func (*SVGPathSegCurvetoCubicSmoothAbs) GetY2() (y2 float32) {
	js.Rewrite("$<.y2")
	return y2
}

func (*SVGPathSegCurvetoCubicSmoothAbs) SetY2(y2 float32) {
	js.Rewrite("$<.y2 = $1", y2)
}
