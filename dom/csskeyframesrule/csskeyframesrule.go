package csskeyframesrule

import (
	"github.com/matthewmueller/joy/dom/csskeyframerule"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.CSSRule = (*CSSKeyframesRule)(nil)

// CSSKeyframesRule struct
// js:"CSSKeyframesRule,omit"
type CSSKeyframesRule struct {
}

// AppendRule fn
// js:"appendRule"
func (*CSSKeyframesRule) AppendRule(rule string) {
	macro.Rewrite("$_.appendRule($1)", rule)
}

// DeleteRule fn
// js:"deleteRule"
func (*CSSKeyframesRule) DeleteRule(rule string) {
	macro.Rewrite("$_.deleteRule($1)", rule)
}

// FindRule fn
// js:"findRule"
func (*CSSKeyframesRule) FindRule(rule string) (c *csskeyframerule.CSSKeyframeRule) {
	macro.Rewrite("$_.findRule($1)", rule)
	return c
}

// CSSRules prop
// js:"cssRules"
func (*CSSKeyframesRule) CSSRules() (cssRules *window.CSSRuleList) {
	macro.Rewrite("$_.cssRules")
	return cssRules
}

// Name prop
// js:"name"
func (*CSSKeyframesRule) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// SetName prop
// js:"name"
func (*CSSKeyframesRule) SetName(name string) {
	macro.Rewrite("$_.name = $1", name)
}

// CSSText prop
// js:"cssText"
func (*CSSKeyframesRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSKeyframesRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSKeyframesRule) ParentRule() (parentRule window.CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSKeyframesRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSKeyframesRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
