package csskeyframesrule

import "github.com/matthewmueller/golly/js"

// js:"CSSKeyframesRule,omit"
type CSSKeyframesRule struct {
	window.CSSRule
}

// AppendRule
func (*CSSKeyframesRule) AppendRule(rule string) {
	js.Rewrite("$<.AppendRule($1)", rule)
}

// DeleteRule
func (*CSSKeyframesRule) DeleteRule(rule string) {
	js.Rewrite("$<.DeleteRule($1)", rule)
}

// FindRule
func (*CSSKeyframesRule) FindRule(rule string) (c *csskeyframerule.CSSKeyframeRule) {
	js.Rewrite("$<.FindRule($1)", rule)
	return c
}

// CSSRules
func (*CSSKeyframesRule) CSSRules() (cssRules *window.CSSRuleList) {
	js.Rewrite("$<.CSSRules")
	return cssRules
}

// Name
func (*CSSKeyframesRule) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name
func (*CSSKeyframesRule) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}
