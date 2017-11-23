package cssconditionrule

import "github.com/matthewmueller/golly/dom2/cssgroupingrule"

// js:"CSSConditionRule,omit"
type CSSConditionRule interface {
	cssgroupingrule.CSSGroupingRule

	// ConditionText
	ConditionText() (conditionText string)

	// ConditionText
	SetConditionText(conditionText string)
}
