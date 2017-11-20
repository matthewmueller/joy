package svgpathseglinetoverticalrel

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegLinetoVerticalRel struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegLinetoVerticalRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegLinetoVerticalRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
