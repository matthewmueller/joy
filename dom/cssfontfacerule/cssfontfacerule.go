package cssfontfacerule

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// CSSFontFaceRule struct
// js:"CSSFontFaceRule,omit"
type CSSFontFaceRule struct {
	window.CSSRule
}

// Style prop
func (*CSSFontFaceRule) Style() (style *window.CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}
