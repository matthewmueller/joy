package cssstylerule

import "github.com/matthewmueller/golly/js"

// js:"CSSStyleRule,omit"
type CSSStyleRule struct {
	window.CSSRule
}

// ReadOnly
func (*CSSStyleRule) ReadOnly() (readOnly bool) {
	js.Rewrite("$<.ReadOnly")
	return readOnly
}

// SelectorText
func (*CSSStyleRule) SelectorText() (selectorText string) {
	js.Rewrite("$<.SelectorText")
	return selectorText
}

// SelectorText
func (*CSSStyleRule) SetSelectorText(selectorText string) {
	js.Rewrite("$<.SelectorText = $1", selectorText)
}

// Style
func (*CSSStyleRule) Style() (style *window.CSSStyleDeclaration) {
	js.Rewrite("$<.Style")
	return style
}
