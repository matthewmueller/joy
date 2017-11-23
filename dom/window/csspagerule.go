package window

import "github.com/matthewmueller/golly/js"

// CSSPageRule struct
// js:"CSSPageRule,omit"
type CSSPageRule struct {
	CSSRule
}

// PseudoClass prop
func (*CSSPageRule) PseudoClass() (pseudoClass string) {
	js.Rewrite("$<.pseudoClass")
	return pseudoClass
}

// Selector prop
func (*CSSPageRule) Selector() (selector string) {
	js.Rewrite("$<.selector")
	return selector
}

// SelectorText prop
func (*CSSPageRule) SelectorText() (selectorText string) {
	js.Rewrite("$<.selectorText")
	return selectorText
}

// SelectorText prop
func (*CSSPageRule) SetSelectorText(selectorText string) {
	js.Rewrite("$<.selectorText = $1", selectorText)
}

// Style prop
func (*CSSPageRule) Style() (style *CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}
