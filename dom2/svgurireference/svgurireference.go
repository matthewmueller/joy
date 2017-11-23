package svgurireference

import "github.com/matthewmueller/golly/dom2/svganimatedstring"

// js:"SVGURIReference,omit"
type SVGURIReference interface {

	// Href
	Href() (href *svganimatedstring.SVGAnimatedString)
}
