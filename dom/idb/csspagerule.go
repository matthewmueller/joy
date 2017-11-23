package idb

import "github.com/matthewmueller/golly/js"

// CSSPageRule struct
// js:"CSSPageRule,omit"
type CSSPageRule struct {
	CSSRule
}

// PseudoClass
func (*CSSPageRule) PseudoClass() (pseudoClass string) {
	js.Rewrite("$<.PseudoClass")
	return pseudoClass
}

// Selector
func (*CSSPageRule) Selector() (selector string) {
	js.Rewrite("$<.Selector")
	return selector
}

// SelectorText
func (*CSSPageRule) SelectorText() (selectorText string) {
	js.Rewrite("$<.SelectorText")
	return selectorText
}

// SelectorText
func (*CSSPageRule) SetSelectorText(selectorText string) {
	js.Rewrite("$<.SelectorText = $1", selectorText)
}

// Style
func (*CSSPageRule) Style() (style *CSSStyleDeclaration) {
	js.Rewrite("$<.Style")
	return style
}
