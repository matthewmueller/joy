package svgaelement

import "github.com/matthewmueller/golly/js"

// js:"SVGAElement,omit"
type SVGAElement struct {
	window.SVGGraphicsElement
	svgurireference.SVGURIReference
}

// Target
func (*SVGAElement) Target() (target *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Target")
	return target
}
