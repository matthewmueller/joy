package cssconditionrule

import "github.com/matthewmueller/golly/dom/cssgroupingrule"

// CSSConditionRule interface
// js:"CSSConditionRule"
type CSSConditionRule interface {
	cssgroupingrule.CSSGroupingRule

	// ConditionText prop
	// js:"conditionText"
	ConditionText() (conditionText string)

	// SetConditionText prop
	// js:"conditionText"
	SetConditionText(conditionText string)
}
