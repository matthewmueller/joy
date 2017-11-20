package svgpathsegmovetorel

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegMovetoRel struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegMovetoRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegMovetoRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegMovetoRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegMovetoRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
