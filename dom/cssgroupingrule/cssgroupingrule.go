package cssgroupingrule

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// js:"CSSGroupingRule,omit"
type CSSGroupingRule interface {
	window.CSSRule

	// DeleteRule
	// js:"deleteRule,rewrite=deleterule"
	DeleteRule(index uint)

	// InsertRule
	// js:"insertRule,rewrite=insertrule"
	InsertRule(rule string, index uint) (u uint)

	// CSSRules prop
	// js:"cssRules,rewrite=cssrules"
	CSSRules() (cssRules *window.CSSRuleList)
}

// deleterule fn
func deleterule(index uint) {
	js.Rewrite("$<.deleteRule($1)", index)
}

// insertrule fn
func insertrule(rule string, index uint) (u uint) {
	js.Rewrite("$<.insertRule($1, $2)", rule, index)
	return u
}

// cssrules prop
func cssrules() (cssRules *window.CSSRuleList) {
	js.Rewrite("$<.cssRules")
	return cssRules
}
