package csskeyframesrule

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/csskeyframerule"
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrule"
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrulelist"
	"github.com/matthewmueller/golly/js"
)

type CSSKeyframesRule struct {
	*cssrule.CSSRule
}

func (*CSSKeyframesRule) AppendRule(rule string) {
	js.Rewrite("$<.appendRule($1)", rule)
}

func (*CSSKeyframesRule) DeleteRule(rule string) {
	js.Rewrite("$<.deleteRule($1)", rule)
}

func (*CSSKeyframesRule) FindRule(rule string) (c *csskeyframerule.CSSKeyframeRule) {
	js.Rewrite("$<.findRule($1)", rule)
	return c
}

func (*CSSKeyframesRule) GetCSSRules() (cssRules *cssrulelist.CSSRuleList) {
	js.Rewrite("$<.cssRules")
	return cssRules
}

func (*CSSKeyframesRule) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*CSSKeyframesRule) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}
