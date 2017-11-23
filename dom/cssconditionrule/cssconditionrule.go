package cssconditionrule

import "github.com/matthewmueller/golly/dom2/cssgroupingrule"

// js:"CSSConditionRule,omit"
type CSSConditionRule interface {
	cssgroupingrule.CSSGroupingRule

	// ConditionText prop
	// js:"conditionText"
	ConditionText() (conditionText string)

	// ConditionText prop
	SetConditionText(conditionText string)
}
