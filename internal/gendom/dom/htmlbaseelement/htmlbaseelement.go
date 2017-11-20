package htmlbaseelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLBaseElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLBaseElement) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*HTMLBaseElement) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

func (*HTMLBaseElement) GetTarget() (target string) {
	js.Rewrite("$<.target")
	return target
}

func (*HTMLBaseElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}
