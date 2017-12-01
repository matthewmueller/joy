package csskeyframerule

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.CSSRule = (*CSSKeyframeRule)(nil)

// CSSKeyframeRule struct
// js:"CSSKeyframeRule,omit"
type CSSKeyframeRule struct {
}

// KeyText prop
// js:"keyText"
func (*CSSKeyframeRule) KeyText() (keyText string) {
	js.Rewrite("$_.keyText")
	return keyText
}

// SetKeyText prop
// js:"keyText"
func (*CSSKeyframeRule) SetKeyText(keyText string) {
	js.Rewrite("$_.keyText = $1", keyText)
}

// Style prop
// js:"style"
func (*CSSKeyframeRule) Style() (style *window.CSSStyleDeclaration) {
	js.Rewrite("$_.style")
	return style
}

// CSSText prop
// js:"cssText"
func (*CSSKeyframeRule) CSSText() (cssText string) {
	js.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSKeyframeRule) SetCSSText(cssText string) {
	js.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSKeyframeRule) ParentRule() (parentRule window.CSSRule) {
	js.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSKeyframeRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	js.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSKeyframeRule) Type() (kind uint8) {
	js.Rewrite("$_.type")
	return kind
}
