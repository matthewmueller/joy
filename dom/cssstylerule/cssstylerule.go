package cssstylerule

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.CSSRule = (*CSSStyleRule)(nil)

// CSSStyleRule struct
// js:"CSSStyleRule,omit"
type CSSStyleRule struct {
}

// ReadOnly prop
// js:"readOnly"
func (*CSSStyleRule) ReadOnly() (readOnly bool) {
	js.Rewrite("$_.readOnly")
	return readOnly
}

// SelectorText prop
// js:"selectorText"
func (*CSSStyleRule) SelectorText() (selectorText string) {
	js.Rewrite("$_.selectorText")
	return selectorText
}

// SetSelectorText prop
// js:"selectorText"
func (*CSSStyleRule) SetSelectorText(selectorText string) {
	js.Rewrite("$_.selectorText = $1", selectorText)
}

// Style prop
// js:"style"
func (*CSSStyleRule) Style() (style *window.CSSStyleDeclaration) {
	js.Rewrite("$_.style")
	return style
}

// CSSText prop
// js:"cssText"
func (*CSSStyleRule) CSSText() (cssText string) {
	js.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSStyleRule) SetCSSText(cssText string) {
	js.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSStyleRule) ParentRule() (parentRule window.CSSRule) {
	js.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSStyleRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	js.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSStyleRule) Type() (kind uint8) {
	js.Rewrite("$_.type")
	return kind
}
