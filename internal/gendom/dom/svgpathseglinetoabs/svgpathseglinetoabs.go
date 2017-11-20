package svgpathseglinetoabs

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegLinetoAbs struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegLinetoAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegLinetoAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegLinetoAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegLinetoAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
