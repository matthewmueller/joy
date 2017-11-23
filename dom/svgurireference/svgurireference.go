package svgurireference

import (
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/js"
)

// SVGURIReference struct
// js:"SVGURIReference,omit"
type SVGURIReference struct {
}

// Href prop
func (*SVGURIReference) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.href")
	return href
}
