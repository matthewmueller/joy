package window

import (
	"github.com/matthewmueller/joy/dom/medialist"
	"github.com/matthewmueller/joy/js"
)

var _ StyleSheet = (*CSSStyleSheet)(nil)

// CSSStyleSheet struct
// js:"CSSStyleSheet,omit"
type CSSStyleSheet struct {
}

// AddImport fn
// js:"addImport"
func (*CSSStyleSheet) AddImport(bstrURL string, lIndex *int) (i int) {
	js.Rewrite("$_.addImport($1, $2)", bstrURL, lIndex)
	return i
}

// AddPageRule fn
// js:"addPageRule"
func (*CSSStyleSheet) AddPageRule(bstrSelector string, bstrStyle string, lIndex *int) (i int) {
	js.Rewrite("$_.addPageRule($1, $2, $3)", bstrSelector, bstrStyle, lIndex)
	return i
}

// AddRule fn
// js:"addRule"
func (*CSSStyleSheet) AddRule(bstrSelector string, bstrStyle *string, lIndex *int) (i int) {
	js.Rewrite("$_.addRule($1, $2, $3)", bstrSelector, bstrStyle, lIndex)
	return i
}

// DeleteRule fn
// js:"deleteRule"
func (*CSSStyleSheet) DeleteRule(index *uint) {
	js.Rewrite("$_.deleteRule($1)", index)
}

// InsertRule fn
// js:"insertRule"
func (*CSSStyleSheet) InsertRule(rule string, index *uint) (u uint) {
	js.Rewrite("$_.insertRule($1, $2)", rule, index)
	return u
}

// RemoveImport fn
// js:"removeImport"
func (*CSSStyleSheet) RemoveImport(lIndex int) {
	js.Rewrite("$_.removeImport($1)", lIndex)
}

// RemoveRule fn
// js:"removeRule"
func (*CSSStyleSheet) RemoveRule(lIndex int) {
	js.Rewrite("$_.removeRule($1)", lIndex)
}

// CSSRules prop
// js:"cssRules"
func (*CSSStyleSheet) CSSRules() (cssRules *CSSRuleList) {
	js.Rewrite("$_.cssRules")
	return cssRules
}

// CSSText prop
// js:"cssText"
func (*CSSStyleSheet) CSSText() (cssText string) {
	js.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSStyleSheet) SetCSSText(cssText string) {
	js.Rewrite("$_.cssText = $1", cssText)
}

// ID prop
// js:"id"
func (*CSSStyleSheet) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// Imports prop
// js:"imports"
func (*CSSStyleSheet) Imports() (imports *StyleSheetList) {
	js.Rewrite("$_.imports")
	return imports
}

// IsAlternate prop
// js:"isAlternate"
func (*CSSStyleSheet) IsAlternate() (isAlternate bool) {
	js.Rewrite("$_.isAlternate")
	return isAlternate
}

// IsPrefAlternate prop
// js:"isPrefAlternate"
func (*CSSStyleSheet) IsPrefAlternate() (isPrefAlternate bool) {
	js.Rewrite("$_.isPrefAlternate")
	return isPrefAlternate
}

// OwnerRule prop
// js:"ownerRule"
func (*CSSStyleSheet) OwnerRule() (ownerRule CSSRule) {
	js.Rewrite("$_.ownerRule")
	return ownerRule
}

// OwningElement prop
// js:"owningElement"
func (*CSSStyleSheet) OwningElement() (owningElement Element) {
	js.Rewrite("$_.owningElement")
	return owningElement
}

// Pages prop
// js:"pages"
func (*CSSStyleSheet) Pages() (pages *StyleSheetPageList) {
	js.Rewrite("$_.pages")
	return pages
}

// ReadOnly prop
// js:"readOnly"
func (*CSSStyleSheet) ReadOnly() (readOnly bool) {
	js.Rewrite("$_.readOnly")
	return readOnly
}

// Rules prop
// js:"rules"
func (*CSSStyleSheet) Rules() (rules *CSSRuleList) {
	js.Rewrite("$_.rules")
	return rules
}

// Disabled prop
// js:"disabled"
func (*CSSStyleSheet) Disabled() (disabled bool) {
	js.Rewrite("$_.disabled")
	return disabled
}

// SetDisabled prop
// js:"disabled"
func (*CSSStyleSheet) SetDisabled(disabled bool) {
	js.Rewrite("$_.disabled = $1", disabled)
}

// Href prop
// js:"href"
func (*CSSStyleSheet) Href() (href string) {
	js.Rewrite("$_.href")
	return href
}

// Media prop
// js:"media"
func (*CSSStyleSheet) Media() (media *medialist.MediaList) {
	js.Rewrite("$_.media")
	return media
}

// OwnerNode prop
// js:"ownerNode"
func (*CSSStyleSheet) OwnerNode() (ownerNode Node) {
	js.Rewrite("$_.ownerNode")
	return ownerNode
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSStyleSheet) ParentStyleSheet() (parentStyleSheet StyleSheet) {
	js.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Title prop
// js:"title"
func (*CSSStyleSheet) Title() (title string) {
	js.Rewrite("$_.title")
	return title
}

// Type prop
// js:"type"
func (*CSSStyleSheet) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
