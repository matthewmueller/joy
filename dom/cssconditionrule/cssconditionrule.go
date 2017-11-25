package cssconditionrule

import "github.com/matthewmueller/golly/dom/cssgroupingrule"

// js:"CSSConditionRule,omit"
type CSSConditionRule interface {
	cssgroupingrule.CSSGroupingRule

	// ConditionText prop
	// js:"conditionText"
	ConditionText() (conditionText string)

	// ConditionText prop
	// js:"setconditionText"
	SetConditionText(conditionText string)
}
