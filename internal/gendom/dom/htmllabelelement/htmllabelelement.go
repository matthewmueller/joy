package htmllabelelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLLabelElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLLabelElement) GetForm() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLLabelElement) GetHTMLFor() (htmlFor string) {
	js.Rewrite("$<.htmlFor")
	return htmlFor
}

func (*HTMLLabelElement) SetHTMLFor(htmlFor string) {
	js.Rewrite("$<.htmlFor = $1", htmlFor)
}
