package cssgroupingrule

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrule"
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrulelist"
	"github.com/matthewmueller/golly/js"
)

type CSSGroupingRule struct {
	*cssrule.CSSRule
}

func (*CSSGroupingRule) DeleteRule(index uint) {
	js.Rewrite("$<.deleteRule($1)", index)
}

func (*CSSGroupingRule) InsertRule(rule string, index uint) (u uint) {
	js.Rewrite("$<.insertRule($1, $2)", rule, index)
	return u
}

func (*CSSGroupingRule) GetCSSRules() (cssRules *cssrulelist.CSSRuleList) {
	js.Rewrite("$<.cssRules")
	return cssRules
}
