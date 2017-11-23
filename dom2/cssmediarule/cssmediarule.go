package cssmediarule

import (
	"github.com/matthewmueller/golly/dom2/cssconditionrule"
	"github.com/matthewmueller/golly/js"
)

// js:"CSSMediaRule,omit"
type CSSMediaRule struct {
	cssconditionrule.CSSConditionRule
}

// Media
func (*CSSMediaRule) Media() (media *medialist.MediaList) {
	js.Rewrite("$<.Media")
	return media
}
