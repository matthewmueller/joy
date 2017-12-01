package cssfontfacerule

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.CSSRule = (*CSSFontFaceRule)(nil)

// CSSFontFaceRule struct
// js:"CSSFontFaceRule,omit"
type CSSFontFaceRule struct {
}

// Style prop
// js:"style"
func (*CSSFontFaceRule) Style() (style *window.CSSStyleDeclaration) {
	js.Rewrite("$_.style")
	return style
}

// CSSText prop
// js:"cssText"
func (*CSSFontFaceRule) CSSText() (cssText string) {
	js.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSFontFaceRule) SetCSSText(cssText string) {
	js.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSFontFaceRule) ParentRule() (parentRule window.CSSRule) {
	js.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSFontFaceRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	js.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSFontFaceRule) Type() (kind uint8) {
	js.Rewrite("$_.type")
	return kind
}
