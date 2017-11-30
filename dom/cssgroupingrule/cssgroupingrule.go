package cssgroupingrule

import "github.com/matthewmueller/golly/dom/window"

// CSSGroupingRule interface
// js:"CSSGroupingRule"
type CSSGroupingRule interface {
	window.CSSRule

	// DeleteRule
	// js:"deleteRule"
	DeleteRule(index uint)

	// InsertRule
	// js:"insertRule"
	InsertRule(rule string, index uint) (u uint)

	// CSSRules prop
	// js:"cssRules"
	CSSRules() (cssRules *window.CSSRuleList)
}
