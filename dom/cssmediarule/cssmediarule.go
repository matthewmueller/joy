package cssmediarule

import (
	"github.com/matthewmueller/joy/dom/cssconditionrule"
	"github.com/matthewmueller/joy/dom/cssgroupingrule"
	"github.com/matthewmueller/joy/dom/medialist"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ cssconditionrule.CSSConditionRule = (*CSSMediaRule)(nil)
var _ cssgroupingrule.CSSGroupingRule = (*CSSMediaRule)(nil)
var _ window.CSSRule = (*CSSMediaRule)(nil)

// CSSMediaRule struct
// js:"CSSMediaRule,omit"
type CSSMediaRule struct {
}

// DeleteRule fn
// js:"deleteRule"
func (*CSSMediaRule) DeleteRule(index uint) {
	macro.Rewrite("$_.deleteRule($1)", index)
}

// InsertRule fn
// js:"insertRule"
func (*CSSMediaRule) InsertRule(rule string, index uint) (u uint) {
	macro.Rewrite("$_.insertRule($1, $2)", rule, index)
	return u
}

// Media prop
// js:"media"
func (*CSSMediaRule) Media() (media *medialist.MediaList) {
	macro.Rewrite("$_.media")
	return media
}

// ConditionText prop
// js:"conditionText"
func (*CSSMediaRule) ConditionText() (conditionText string) {
	macro.Rewrite("$_.conditionText")
	return conditionText
}

// SetConditionText prop
// js:"conditionText"
func (*CSSMediaRule) SetConditionText(conditionText string) {
	macro.Rewrite("$_.conditionText = $1", conditionText)
}

// CSSRules prop
// js:"cssRules"
func (*CSSMediaRule) CSSRules() (cssRules *window.CSSRuleList) {
	macro.Rewrite("$_.cssRules")
	return cssRules
}

// CSSText prop
// js:"cssText"
func (*CSSMediaRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSMediaRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSMediaRule) ParentRule() (parentRule window.CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSMediaRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSMediaRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
