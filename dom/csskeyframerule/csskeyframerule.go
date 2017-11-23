package csskeyframerule

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// CSSKeyframeRule struct
// js:"CSSKeyframeRule,omit"
type CSSKeyframeRule struct {
	window.CSSRule
}

// KeyText prop
func (*CSSKeyframeRule) KeyText() (keyText string) {
	js.Rewrite("$<.keyText")
	return keyText
}

// KeyText prop
func (*CSSKeyframeRule) SetKeyText(keyText string) {
	js.Rewrite("$<.keyText = $1", keyText)
}

// Style prop
func (*CSSKeyframeRule) Style() (style *window.CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}
