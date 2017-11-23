package cssfontfacerule

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"CSSFontFaceRule,omit"
type CSSFontFaceRule struct {
	window.CSSRule
}

// Style
func (*CSSFontFaceRule) Style() (style *window.CSSStyleDeclaration) {
	js.Rewrite("$<.Style")
	return style
}
