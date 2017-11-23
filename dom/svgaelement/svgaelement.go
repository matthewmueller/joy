package svgaelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGAElement struct
// js:"SVGAElement,omit"
type SVGAElement struct {
	window.SVGGraphicsElement
}

// Href prop
func (*SVGAElement) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.href")
	return href
}

// Target prop
func (*SVGAElement) Target() (target *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.target")
	return target
}
