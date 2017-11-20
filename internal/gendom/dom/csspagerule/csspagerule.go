package csspagerule

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrule"
	"github.com/matthewmueller/golly/internal/gendom/dom/cssstyledeclaration"
	"github.com/matthewmueller/golly/js"
)

type CSSPageRule struct {
	*cssrule.CSSRule
}

func (*CSSPageRule) GetPseudoClass() (pseudoClass string) {
	js.Rewrite("$<.pseudoClass")
	return pseudoClass
}

func (*CSSPageRule) GetSelector() (selector string) {
	js.Rewrite("$<.selector")
	return selector
}

func (*CSSPageRule) GetSelectorText() (selectorText string) {
	js.Rewrite("$<.selectorText")
	return selectorText
}

func (*CSSPageRule) SetSelectorText(selectorText string) {
	js.Rewrite("$<.selectorText = $1", selectorText)
}

func (*CSSPageRule) GetStyle() (style *cssstyledeclaration.CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}
