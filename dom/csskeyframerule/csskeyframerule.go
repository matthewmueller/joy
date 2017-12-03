package csskeyframerule

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.CSSRule = (*CSSKeyframeRule)(nil)

// CSSKeyframeRule struct
// js:"CSSKeyframeRule,omit"
type CSSKeyframeRule struct {
}

// KeyText prop
// js:"keyText"
func (*CSSKeyframeRule) KeyText() (keyText string) {
	macro.Rewrite("$_.keyText")
	return keyText
}

// SetKeyText prop
// js:"keyText"
func (*CSSKeyframeRule) SetKeyText(keyText string) {
	macro.Rewrite("$_.keyText = $1", keyText)
}

// Style prop
// js:"style"
func (*CSSKeyframeRule) Style() (style *window.CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

// CSSText prop
// js:"cssText"
func (*CSSKeyframeRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSKeyframeRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSKeyframeRule) ParentRule() (parentRule window.CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSKeyframeRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSKeyframeRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
