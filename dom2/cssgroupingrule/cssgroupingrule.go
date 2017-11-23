package cssgroupingrule

// js:"CSSGroupingRule,omit"
type CSSGroupingRule interface {
	window.CSSRule

	// DeleteRule
	DeleteRule(index uint)

	// InsertRule
	InsertRule(rule string, index uint) (u uint)

	// CSSRules
	CSSRules() (cssRules *window.CSSRuleList)
}
