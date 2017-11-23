package svgaelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/svgurireference"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

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
