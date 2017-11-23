package svgscriptelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGScriptElement struct
// js:"SVGScriptElement,omit"
type SVGScriptElement struct {
	window.SVGElement
}

// Href
func (*SVGScriptElement) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Href")
	return href
}

// Type
func (*SVGScriptElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type
func (*SVGScriptElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}
