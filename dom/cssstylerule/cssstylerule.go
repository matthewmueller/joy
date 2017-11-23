package cssstylerule

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// CSSStyleRule struct
// js:"CSSStyleRule,omit"
type CSSStyleRule struct {
	window.CSSRule
}

// ReadOnly prop
func (*CSSStyleRule) ReadOnly() (readOnly bool) {
	js.Rewrite("$<.readOnly")
	return readOnly
}

// SelectorText prop
func (*CSSStyleRule) SelectorText() (selectorText string) {
	js.Rewrite("$<.selectorText")
	return selectorText
}

// SelectorText prop
func (*CSSStyleRule) SetSelectorText(selectorText string) {
	js.Rewrite("$<.selectorText = $1", selectorText)
}

// Style prop
func (*CSSStyleRule) Style() (style *window.CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}
