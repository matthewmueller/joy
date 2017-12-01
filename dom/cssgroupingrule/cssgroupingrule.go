package cssgroupingrule

import "github.com/matthewmueller/golly/dom/window"

// CSSGroupingRule interface
// js:"CSSGroupingRule"
type CSSGroupingRule interface {
	window.CSSRule

	// DeleteRule
	// js:"deleteRule"
	// jsrewrite:"$_.deleteRule($1)"
	DeleteRule(index uint)

	// InsertRule
	// js:"insertRule"
	// jsrewrite:"$_.insertRule($1, $2)"
	InsertRule(rule string, index uint) (u uint)

	// CSSRules prop
	// js:"cssRules"
	// jsrewrite:"$_.cssRules"
	CSSRules() (cssRules *window.CSSRuleList)
}
