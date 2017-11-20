package cssmediarule

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssconditionrule"
	"github.com/matthewmueller/golly/internal/gendom/dom/medialist"
	"github.com/matthewmueller/golly/js"
)

type CSSMediaRule struct {
	*cssconditionrule.CSSConditionRule
}

func (*CSSMediaRule) GetMedia() (media *medialist.MediaList) {
	js.Rewrite("$<.media")
	return media
}
