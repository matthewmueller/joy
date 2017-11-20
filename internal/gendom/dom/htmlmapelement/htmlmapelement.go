package htmlmapelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlareascollection"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLMapElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLMapElement) GetAreas() (areas *htmlareascollection.HTMLAreasCollection) {
	js.Rewrite("$<.areas")
	return areas
}

func (*HTMLMapElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLMapElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}
