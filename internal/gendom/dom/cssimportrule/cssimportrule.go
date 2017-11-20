package cssimportrule

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrule"
	"github.com/matthewmueller/golly/internal/gendom/dom/cssstylesheet"
	"github.com/matthewmueller/golly/internal/gendom/dom/medialist"
	"github.com/matthewmueller/golly/js"
)

type CSSImportRule struct {
	*cssrule.CSSRule
}

func (*CSSImportRule) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*CSSImportRule) GetMedia() (media *medialist.MediaList) {
	js.Rewrite("$<.media")
	return media
}

func (*CSSImportRule) GetStyleSheet() (styleSheet *cssstylesheet.CSSStyleSheet) {
	js.Rewrite("$<.styleSheet")
	return styleSheet
}
