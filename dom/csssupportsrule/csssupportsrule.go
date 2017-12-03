package csssupportsrule

import (
	"github.com/matthewmueller/joy/dom/cssconditionrule"
	"github.com/matthewmueller/joy/dom/cssgroupingrule"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ cssconditionrule.CSSConditionRule = (*CSSSupportsRule)(nil)
var _ cssgroupingrule.CSSGroupingRule = (*CSSSupportsRule)(nil)
var _ window.CSSRule = (*CSSSupportsRule)(nil)

// CSSSupportsRule struct
// js:"CSSSupportsRule,omit"
type CSSSupportsRule struct {
}

// DeleteRule fn
// js:"deleteRule"
func (*CSSSupportsRule) DeleteRule(index uint) {
	macro.Rewrite("$_.deleteRule($1)", index)
}

// InsertRule fn
// js:"insertRule"
func (*CSSSupportsRule) InsertRule(rule string, index uint) (u uint) {
	macro.Rewrite("$_.insertRule($1, $2)", rule, index)
	return u
}

// ConditionText prop
// js:"conditionText"
func (*CSSSupportsRule) ConditionText() (conditionText string) {
	macro.Rewrite("$_.conditionText")
	return conditionText
}

// SetConditionText prop
// js:"conditionText"
func (*CSSSupportsRule) SetConditionText(conditionText string) {
	macro.Rewrite("$_.conditionText = $1", conditionText)
}

// CSSRules prop
// js:"cssRules"
func (*CSSSupportsRule) CSSRules() (cssRules *window.CSSRuleList) {
	macro.Rewrite("$_.cssRules")
	return cssRules
}

// CSSText prop
// js:"cssText"
func (*CSSSupportsRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSSupportsRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSSupportsRule) ParentRule() (parentRule window.CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSSupportsRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSSupportsRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
