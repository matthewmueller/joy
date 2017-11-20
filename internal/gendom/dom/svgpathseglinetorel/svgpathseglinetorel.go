package svgpathseglinetorel

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegLinetoRel struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegLinetoRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegLinetoRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegLinetoRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegLinetoRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
