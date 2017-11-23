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
