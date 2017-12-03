package cssfontfacerule

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.CSSRule = (*CSSFontFaceRule)(nil)

// CSSFontFaceRule struct
// js:"CSSFontFaceRule,omit"
type CSSFontFaceRule struct {
}

// Style prop
// js:"style"
func (*CSSFontFaceRule) Style() (style *window.CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

// CSSText prop
// js:"cssText"
func (*CSSFontFaceRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSFontFaceRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSFontFaceRule) ParentRule() (parentRule window.CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSFontFaceRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSFontFaceRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
