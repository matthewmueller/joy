package cssconditionrule

import "github.com/matthewmueller/joy/dom/cssgroupingrule"

// CSSConditionRule interface
// js:"CSSConditionRule"
type CSSConditionRule interface {
	cssgroupingrule.CSSGroupingRule

	// ConditionText prop
	// js:"conditionText"
	// jsrewrite:"$_.conditionText"
	ConditionText() (conditionText string)

	// SetConditionText prop
	// js:"conditionText"
	// jsrewrite:"$_.conditionText = $1"
	SetConditionText(conditionText string)
}
