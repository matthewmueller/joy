package window

import "github.com/matthewmueller/golly/js"

// CSSStyleSheet struct
// js:"CSSStyleSheet,omit"
type CSSStyleSheet struct {
	StyleSheet
}

// AddImport fn
func (*CSSStyleSheet) AddImport(bstrURL string, lIndex *int) (i int) {
	js.Rewrite("$<.addImport($1, $2)", bstrURL, lIndex)
	return i
}

// AddPageRule fn
func (*CSSStyleSheet) AddPageRule(bstrSelector string, bstrStyle string, lIndex *int) (i int) {
	js.Rewrite("$<.addPageRule($1, $2, $3)", bstrSelector, bstrStyle, lIndex)
	return i
}

// AddRule fn
func (*CSSStyleSheet) AddRule(bstrSelector string, bstrStyle *string, lIndex *int) (i int) {
	js.Rewrite("$<.addRule($1, $2, $3)", bstrSelector, bstrStyle, lIndex)
	return i
}

// DeleteRule fn
func (*CSSStyleSheet) DeleteRule(index *uint) {
	js.Rewrite("$<.deleteRule($1)", index)
}

// InsertRule fn
func (*CSSStyleSheet) InsertRule(rule string, index *uint) (u uint) {
	js.Rewrite("$<.insertRule($1, $2)", rule, index)
	return u
}

// RemoveImport fn
func (*CSSStyleSheet) RemoveImport(lIndex int) {
	js.Rewrite("$<.removeImport($1)", lIndex)
}

// RemoveRule fn
func (*CSSStyleSheet) RemoveRule(lIndex int) {
	js.Rewrite("$<.removeRule($1)", lIndex)
}

// CSSRules prop
func (*CSSStyleSheet) CSSRules() (cssRules *CSSRuleList) {
	js.Rewrite("$<.cssRules")
	return cssRules
}

// CSSText prop
func (*CSSStyleSheet) CSSText() (cssText string) {
	js.Rewrite("$<.cssText")
	return cssText
}

// CSSText prop
func (*CSSStyleSheet) SetCSSText(cssText string) {
	js.Rewrite("$<.cssText = $1", cssText)
}

// ID prop
func (*CSSStyleSheet) ID() (id string) {
	js.Rewrite("$<.id")
	return id
}

// Imports prop
func (*CSSStyleSheet) Imports() (imports *StyleSheetList) {
	js.Rewrite("$<.imports")
	return imports
}

// IsAlternate prop
func (*CSSStyleSheet) IsAlternate() (isAlternate bool) {
	js.Rewrite("$<.isAlternate")
	return isAlternate
}

// IsPrefAlternate prop
func (*CSSStyleSheet) IsPrefAlternate() (isPrefAlternate bool) {
	js.Rewrite("$<.isPrefAlternate")
	return isPrefAlternate
}

// OwnerRule prop
func (*CSSStyleSheet) OwnerRule() (ownerRule CSSRule) {
	js.Rewrite("$<.ownerRule")
	return ownerRule
}

// OwningElement prop
func (*CSSStyleSheet) OwningElement() (owningElement Element) {
	js.Rewrite("$<.owningElement")
	return owningElement
}

// Pages prop
func (*CSSStyleSheet) Pages() (pages *StyleSheetPageList) {
	js.Rewrite("$<.pages")
	return pages
}

// ReadOnly prop
func (*CSSStyleSheet) ReadOnly() (readOnly bool) {
	js.Rewrite("$<.readOnly")
	return readOnly
}

// Rules prop
func (*CSSStyleSheet) Rules() (rules *CSSRuleList) {
	js.Rewrite("$<.rules")
	return rules
}
