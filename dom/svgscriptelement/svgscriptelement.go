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

// Href prop
func (*SVGScriptElement) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.href")
	return href
}

// Type prop
func (*SVGScriptElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop
func (*SVGScriptElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
