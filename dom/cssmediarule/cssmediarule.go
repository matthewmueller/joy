package cssmediarule

import (
	"github.com/matthewmueller/golly/dom2/cssconditionrule"
	"github.com/matthewmueller/golly/dom2/medialist"
	"github.com/matthewmueller/golly/js"
)

// CSSMediaRule struct
// js:"CSSMediaRule,omit"
type CSSMediaRule struct {
	cssconditionrule.CSSConditionRule
}

// Media prop
func (*CSSMediaRule) Media() (media *medialist.MediaList) {
	js.Rewrite("$<.media")
	return media
}
