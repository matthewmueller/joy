package csssupportsrule

import (
	"github.com/matthewmueller/golly/dom/cssconditionrule"
	"github.com/matthewmueller/golly/dom/cssgroupingrule"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
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
	js.Rewrite("$_.deleteRule($1)", index)
}

// InsertRule fn
// js:"insertRule"
func (*CSSSupportsRule) InsertRule(rule string, index uint) (u uint) {
	js.Rewrite("$_.insertRule($1, $2)", rule, index)
	return u
}

// ConditionText prop
// js:"conditionText"
func (*CSSSupportsRule) ConditionText() (conditionText string) {
	js.Rewrite("$_.conditionText")
	return conditionText
}

// SetConditionText prop
// js:"conditionText"
func (*CSSSupportsRule) SetConditionText(conditionText string) {
	js.Rewrite("$_.conditionText = $1", conditionText)
}

// CSSRules prop
// js:"cssRules"
func (*CSSSupportsRule) CSSRules() (cssRules *window.CSSRuleList) {
	js.Rewrite("$_.cssRules")
	return cssRules
}

// CSSText prop
// js:"cssText"
func (*CSSSupportsRule) CSSText() (cssText string) {
	js.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSSupportsRule) SetCSSText(cssText string) {
	js.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSSupportsRule) ParentRule() (parentRule window.CSSRule) {
	js.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSSupportsRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	js.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSSupportsRule) Type() (kind uint8) {
	js.Rewrite("$_.type")
	return kind
}
