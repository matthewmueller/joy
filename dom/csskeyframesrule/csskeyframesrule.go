package csskeyframesrule

import (
	"github.com/matthewmueller/golly/dom/csskeyframerule"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.CSSRule = (*CSSKeyframesRule)(nil)

// CSSKeyframesRule struct
// js:"CSSKeyframesRule,omit"
type CSSKeyframesRule struct {
}

// AppendRule fn
// js:"appendRule"
func (*CSSKeyframesRule) AppendRule(rule string) {
	js.Rewrite("$_.appendRule($1)", rule)
}

// DeleteRule fn
// js:"deleteRule"
func (*CSSKeyframesRule) DeleteRule(rule string) {
	js.Rewrite("$_.deleteRule($1)", rule)
}

// FindRule fn
// js:"findRule"
func (*CSSKeyframesRule) FindRule(rule string) (c *csskeyframerule.CSSKeyframeRule) {
	js.Rewrite("$_.findRule($1)", rule)
	return c
}

// CSSRules prop
// js:"cssRules"
func (*CSSKeyframesRule) CSSRules() (cssRules *window.CSSRuleList) {
	js.Rewrite("$_.cssRules")
	return cssRules
}

// Name prop
// js:"name"
func (*CSSKeyframesRule) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// SetName prop
// js:"name"
func (*CSSKeyframesRule) SetName(name string) {
	js.Rewrite("$_.name = $1", name)
}

// CSSText prop
// js:"cssText"
func (*CSSKeyframesRule) CSSText() (cssText string) {
	js.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSKeyframesRule) SetCSSText(cssText string) {
	js.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSKeyframesRule) ParentRule() (parentRule window.CSSRule) {
	js.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSKeyframesRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	js.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSKeyframesRule) Type() (kind uint8) {
	js.Rewrite("$_.type")
	return kind
}
