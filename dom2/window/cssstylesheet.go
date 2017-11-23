package window

import "github.com/matthewmueller/golly/js"

// js:"CSSStyleSheet,omit"
type CSSStyleSheet struct {
	StyleSheet
}

// AddImport
func (*CSSStyleSheet) AddImport(bstrURL string, lIndex *int) (i int) {
	js.Rewrite("$<.AddImport($1, $2)", bstrURL, lIndex)
	return i
}

// AddPageRule
func (*CSSStyleSheet) AddPageRule(bstrSelector string, bstrStyle string, lIndex *int) (i int) {
	js.Rewrite("$<.AddPageRule($1, $2, $3)", bstrSelector, bstrStyle, lIndex)
	return i
}

// AddRule
func (*CSSStyleSheet) AddRule(bstrSelector string, bstrStyle *string, lIndex *int) (i int) {
	js.Rewrite("$<.AddRule($1, $2, $3)", bstrSelector, bstrStyle, lIndex)
	return i
}

// DeleteRule
func (*CSSStyleSheet) DeleteRule(index *uint) {
	js.Rewrite("$<.DeleteRule($1)", index)
}

// InsertRule
func (*CSSStyleSheet) InsertRule(rule string, index *uint) (u uint) {
	js.Rewrite("$<.InsertRule($1, $2)", rule, index)
	return u
}

// RemoveImport
func (*CSSStyleSheet) RemoveImport(lIndex int) {
	js.Rewrite("$<.RemoveImport($1)", lIndex)
}

// RemoveRule
func (*CSSStyleSheet) RemoveRule(lIndex int) {
	js.Rewrite("$<.RemoveRule($1)", lIndex)
}

// CSSRules
func (*CSSStyleSheet) CSSRules() (cssRules *CSSRuleList) {
	js.Rewrite("$<.CSSRules")
	return cssRules
}

// CSSText
func (*CSSStyleSheet) CSSText() (cssText string) {
	js.Rewrite("$<.CSSText")
	return cssText
}

// CSSText
func (*CSSStyleSheet) SetCSSText(cssText string) {
	js.Rewrite("$<.CSSText = $1", cssText)
}

// ID
func (*CSSStyleSheet) ID() (id string) {
	js.Rewrite("$<.ID")
	return id
}

// Imports
func (*CSSStyleSheet) Imports() (imports *StyleSheetList) {
	js.Rewrite("$<.Imports")
	return imports
}

// IsAlternate
func (*CSSStyleSheet) IsAlternate() (isAlternate bool) {
	js.Rewrite("$<.IsAlternate")
	return isAlternate
}

// IsPrefAlternate
func (*CSSStyleSheet) IsPrefAlternate() (isPrefAlternate bool) {
	js.Rewrite("$<.IsPrefAlternate")
	return isPrefAlternate
}

// OwnerRule
func (*CSSStyleSheet) OwnerRule() (ownerRule CSSRule) {
	js.Rewrite("$<.OwnerRule")
	return ownerRule
}

// OwningElement
func (*CSSStyleSheet) OwningElement() (owningElement Element) {
	js.Rewrite("$<.OwningElement")
	return owningElement
}

// Pages
func (*CSSStyleSheet) Pages() (pages *StyleSheetPageList) {
	js.Rewrite("$<.Pages")
	return pages
}

// ReadOnly
func (*CSSStyleSheet) ReadOnly() (readOnly bool) {
	js.Rewrite("$<.ReadOnly")
	return readOnly
}

// Rules
func (*CSSStyleSheet) Rules() (rules *CSSRuleList) {
	js.Rewrite("$<.Rules")
	return rules
}
