package svgurireference

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedstring"
	"github.com/matthewmueller/golly/js"
)

type SVGURIReference struct {
}

func (*SVGURIReference) GetHref() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.href")
	return href
}
