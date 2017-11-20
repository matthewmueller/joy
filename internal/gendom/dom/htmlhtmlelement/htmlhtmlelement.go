package htmlhtmlelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLHtmlElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLHtmlElement) GetVersion() (version string) {
	js.Rewrite("$<.version")
	return version
}

func (*HTMLHtmlElement) SetVersion(version string) {
	js.Rewrite("$<.version = $1", version)
}
