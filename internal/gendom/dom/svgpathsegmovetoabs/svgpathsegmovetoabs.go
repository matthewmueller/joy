package svgpathsegmovetoabs

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegMovetoAbs struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegMovetoAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegMovetoAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegMovetoAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegMovetoAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
