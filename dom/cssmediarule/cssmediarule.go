package cssmediarule

import (
	"github.com/matthewmueller/golly/dom/cssconditionrule"
	"github.com/matthewmueller/golly/dom/medialist"
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
