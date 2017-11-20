package svgpathseglinetohorizontalrel

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegLinetoHorizontalRel struct {
	*svgpathseg.SVGPathSeg
}

func (*SVGPathSegLinetoHorizontalRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegLinetoHorizontalRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}
