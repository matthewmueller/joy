package htmlheadelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLHeadElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLHeadElement) GetProfile() (profile string) {
	js.Rewrite("$<.profile")
	return profile
}

func (*HTMLHeadElement) SetProfile(profile string) {
	js.Rewrite("$<.profile = $1", profile)
}
