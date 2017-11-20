package htmlformcontrolscollection

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlcollection"
	"github.com/matthewmueller/golly/js"
)

type HTMLFormControlsCollection struct {
	*htmlcollection.HTMLCollection
}

func (*HTMLFormControlsCollection) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.namedItem($1)", name)
	return i
}
