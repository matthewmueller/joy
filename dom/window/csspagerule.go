package window

import "github.com/matthewmueller/joy/js"

var _ CSSRule = (*CSSPageRule)(nil)

// CSSPageRule struct
// js:"CSSPageRule,omit"
type CSSPageRule struct {
}

// PseudoClass prop
// js:"pseudoClass"
func (*CSSPageRule) PseudoClass() (pseudoClass string) {
	js.Rewrite("$_.pseudoClass")
	return pseudoClass
}

// Selector prop
// js:"selector"
func (*CSSPageRule) Selector() (selector string) {
	js.Rewrite("$_.selector")
	return selector
}

// SelectorText prop
// js:"selectorText"
func (*CSSPageRule) SelectorText() (selectorText string) {
	js.Rewrite("$_.selectorText")
	return selectorText
}

// SetSelectorText prop
// js:"selectorText"
func (*CSSPageRule) SetSelectorText(selectorText string) {
	js.Rewrite("$_.selectorText = $1", selectorText)
}

// Style prop
// js:"style"
func (*CSSPageRule) Style() (style *CSSStyleDeclaration) {
	js.Rewrite("$_.style")
	return style
}

// CSSText prop
// js:"cssText"
func (*CSSPageRule) CSSText() (cssText string) {
	js.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSPageRule) SetCSSText(cssText string) {
	js.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSPageRule) ParentRule() (parentRule CSSRule) {
	js.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSPageRule) ParentStyleSheet() (parentStyleSheet *CSSStyleSheet) {
	js.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSPageRule) Type() (kind uint8) {
	js.Rewrite("$_.type")
	return kind
}
