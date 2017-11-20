package svgpathseglinetoverticalabs

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegLinetoVerticalAbs struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegLinetoVerticalAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegLinetoVerticalAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
