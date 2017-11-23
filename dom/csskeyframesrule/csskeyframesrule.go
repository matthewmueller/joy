package csskeyframesrule

import (
	"github.com/matthewmueller/golly/dom/csskeyframerule"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// CSSKeyframesRule struct
// js:"CSSKeyframesRule,omit"
type CSSKeyframesRule struct {
	window.CSSRule
}

// AppendRule fn
func (*CSSKeyframesRule) AppendRule(rule string) {
	js.Rewrite("$<.appendRule($1)", rule)
}

// DeleteRule fn
func (*CSSKeyframesRule) DeleteRule(rule string) {
	js.Rewrite("$<.deleteRule($1)", rule)
}

// FindRule fn
func (*CSSKeyframesRule) FindRule(rule string) (c *csskeyframerule.CSSKeyframeRule) {
	js.Rewrite("$<.findRule($1)", rule)
	return c
}

// CSSRules prop
func (*CSSKeyframesRule) CSSRules() (cssRules *window.CSSRuleList) {
	js.Rewrite("$<.cssRules")
	return cssRules
}

// Name prop
func (*CSSKeyframesRule) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop
func (*CSSKeyframesRule) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}
