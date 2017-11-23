package svgurireference

import (
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/js"
)

// SVGURIReference struct
// js:"SVGURIReference,omit"
type SVGURIReference struct {
}

// Href
func (*SVGURIReference) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Href")
	return href
}
