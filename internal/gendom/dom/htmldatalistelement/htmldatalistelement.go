package htmldatalistelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlcollection"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLDataListElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLDataListElement) GetOptions() (options *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.options")
	return options
}
