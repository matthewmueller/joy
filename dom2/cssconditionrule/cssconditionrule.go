package cssconditionrule

// js:"CSSConditionRule,omit"
type CSSConditionRule interface {
	cssgroupingrule.CSSGroupingRule

	// ConditionText
	ConditionText() (conditionText string)

	// ConditionText
	SetConditionText(conditionText string)
}
