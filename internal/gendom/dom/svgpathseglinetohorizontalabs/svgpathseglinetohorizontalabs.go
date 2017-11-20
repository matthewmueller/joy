package svgpathseglinetohorizontalabs

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegLinetoHorizontalAbs struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegLinetoHorizontalAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegLinetoHorizontalAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}
