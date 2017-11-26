package cssconditionrule

import (
	"github.com/matthewmueller/golly/dom/cssgroupingrule"
	"github.com/matthewmueller/golly/js"
)

// js:"CSSConditionRule,omit"
type CSSConditionRule interface {
	cssgroupingrule.CSSGroupingRule

	// ConditionText prop
	// js:"conditionText,rewrite=conditiontext"
	ConditionText() (conditionText string)

	// ConditionText prop
	// js:"setconditionText,rewrite=setconditiontext"
	SetConditionText(conditionText string)
}

// conditiontext prop
func conditiontext() (conditionText string) {
	js.Rewrite("$<.conditionText")
	return conditionText
}

// setconditiontext prop
func setconditiontext(conditionText string) {
	js.Rewrite("$<.conditionText = conditionText")
}
