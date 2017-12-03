package window

import "github.com/matthewmueller/joy/macro"

var _ CSSRule = (*CSSPageRule)(nil)

// CSSPageRule struct
// js:"CSSPageRule,omit"
type CSSPageRule struct {
}

// PseudoClass prop
// js:"pseudoClass"
func (*CSSPageRule) PseudoClass() (pseudoClass string) {
	macro.Rewrite("$_.pseudoClass")
	return pseudoClass
}

// Selector prop
// js:"selector"
func (*CSSPageRule) Selector() (selector string) {
	macro.Rewrite("$_.selector")
	return selector
}

// SelectorText prop
// js:"selectorText"
func (*CSSPageRule) SelectorText() (selectorText string) {
	macro.Rewrite("$_.selectorText")
	return selectorText
}

// SetSelectorText prop
// js:"selectorText"
func (*CSSPageRule) SetSelectorText(selectorText string) {
	macro.Rewrite("$_.selectorText = $1", selectorText)
}

// Style prop
// js:"style"
func (*CSSPageRule) Style() (style *CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

// CSSText prop
// js:"cssText"
func (*CSSPageRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSPageRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSPageRule) ParentRule() (parentRule CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSPageRule) ParentStyleSheet() (parentStyleSheet *CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSPageRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
