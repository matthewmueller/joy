package svgscriptelement

import (
	"github.com/matthewmueller/golly/dom2/svgurireference"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGScriptElement,omit"
type SVGScriptElement struct {
	window.SVGElement
	svgurireference.SVGURIReference
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
