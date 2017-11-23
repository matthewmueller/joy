package csskeyframerule

import "github.com/matthewmueller/golly/js"

// js:"CSSKeyframeRule,omit"
type CSSKeyframeRule struct {
	window.CSSRule
}

// KeyText
func (*CSSKeyframeRule) KeyText() (keyText string) {
	js.Rewrite("$<.KeyText")
	return keyText
}

// KeyText
func (*CSSKeyframeRule) SetKeyText(keyText string) {
	js.Rewrite("$<.KeyText = $1", keyText)
}

// Style
func (*CSSKeyframeRule) Style() (style *window.CSSStyleDeclaration) {
	js.Rewrite("$<.Style")
	return style
}
