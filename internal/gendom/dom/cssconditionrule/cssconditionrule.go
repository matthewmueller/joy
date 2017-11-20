package cssconditionrule

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssgroupingrule"
	"github.com/matthewmueller/golly/js"
)

type CSSConditionRule struct {
	*cssgroupingrule.CSSGroupingRule
}

func (*CSSConditionRule) GetConditionText() (conditionText string) {
	js.Rewrite("$<.conditionText")
	return conditionText
}

func (*CSSConditionRule) SetConditionText(conditionText string) {
	js.Rewrite("$<.conditionText = $1", conditionText)
}
