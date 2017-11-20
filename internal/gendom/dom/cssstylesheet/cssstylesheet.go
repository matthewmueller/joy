package cssstylesheet

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrule"
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrulelist"
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/stylesheet"
	"github.com/matthewmueller/golly/internal/gendom/dom/stylesheetlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/stylesheetpagelist"
	"github.com/matthewmueller/golly/js"
)

type CSSStyleSheet struct {
	*stylesheet.StyleSheet
}

func (*CSSStyleSheet) AddImport(bstrURL string, lIndex int) (i int) {
	js.Rewrite("$<.addImport($1, $2)", bstrURL, lIndex)
	return i
}

func (*CSSStyleSheet) AddPageRule(bstrSelector string, bstrStyle string, lIndex int) (i int) {
	js.Rewrite("$<.addPageRule($1, $2, $3)", bstrSelector, bstrStyle, lIndex)
	return i
}

func (*CSSStyleSheet) AddRule(bstrSelector string, bstrStyle string, lIndex int) (i int) {
	js.Rewrite("$<.addRule($1, $2, $3)", bstrSelector, bstrStyle, lIndex)
	return i
}

func (*CSSStyleSheet) DeleteRule(index uint) {
	js.Rewrite("$<.deleteRule($1)", index)
}

func (*CSSStyleSheet) InsertRule(rule string, index uint) (u uint) {
	js.Rewrite("$<.insertRule($1, $2)", rule, index)
	return u
}

func (*CSSStyleSheet) RemoveImport(lIndex int) {
	js.Rewrite("$<.removeImport($1)", lIndex)
}

func (*CSSStyleSheet) RemoveRule(lIndex int) {
	js.Rewrite("$<.removeRule($1)", lIndex)
}

func (*CSSStyleSheet) GetCSSRules() (cssRules *cssrulelist.CSSRuleList) {
	js.Rewrite("$<.cssRules")
	return cssRules
}

func (*CSSStyleSheet) GetCSSText() (cssText string) {
	js.Rewrite("$<.cssText")
	return cssText
}

func (*CSSStyleSheet) SetCSSText(cssText string) {
	js.Rewrite("$<.cssText = $1", cssText)
}

func (*CSSStyleSheet) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*CSSStyleSheet) GetImports() (imports *stylesheetlist.StyleSheetList) {
	js.Rewrite("$<.imports")
	return imports
}

func (*CSSStyleSheet) GetIsAlternate() (isAlternate bool) {
	js.Rewrite("$<.isAlternate")
	return isAlternate
}

func (*CSSStyleSheet) GetIsPrefAlternate() (isPrefAlternate bool) {
	js.Rewrite("$<.isPrefAlternate")
	return isPrefAlternate
}

func (*CSSStyleSheet) GetOwnerRule() (ownerRule *cssrule.CSSRule) {
	js.Rewrite("$<.ownerRule")
	return ownerRule
}

func (*CSSStyleSheet) GetOwningElement() (owningElement *element.Element) {
	js.Rewrite("$<.owningElement")
	return owningElement
}

func (*CSSStyleSheet) GetPages() (pages *stylesheetpagelist.StyleSheetPageList) {
	js.Rewrite("$<.pages")
	return pages
}

func (*CSSStyleSheet) GetReadOnly() (readOnly bool) {
	js.Rewrite("$<.readOnly")
	return readOnly
}

func (*CSSStyleSheet) GetRules() (rules *cssrulelist.CSSRuleList) {
	js.Rewrite("$<.rules")
	return rules
}
