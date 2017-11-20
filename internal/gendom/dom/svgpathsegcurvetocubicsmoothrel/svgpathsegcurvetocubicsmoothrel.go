package svgpathsegcurvetocubicsmoothrel

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegCurvetoCubicSmoothRel struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegCurvetoCubicSmoothRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoCubicSmoothRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoCubicSmoothRel) GetX2() (x2 float32) {
	js.Rewrite("$<.x2")
	return x2
}

func (*SVGPathSegCurvetoCubicSmoothRel) SetX2(x2 float32) {
	js.Rewrite("$<.x2 = $1", x2)
}

func (*SVGPathSegCurvetoCubicSmoothRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoCubicSmoothRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

func (*SVGPathSegCurvetoCubicSmoothRel) GetY2() (y2 float32) {
	js.Rewrite("$<.y2")
	return y2
}

func (*SVGPathSegCurvetoCubicSmoothRel) SetY2(y2 float32) {
	js.Rewrite("$<.y2 = $1", y2)
}
