package htmlmapelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLMapElement,omit"
type HTMLMapElement struct {
	window.HTMLElement
}

// Areas Retrieves a collection of the area objects defined for the given map object.
func (*HTMLMapElement) Areas() (areas *htmlareascollection.HTMLAreasCollection) {
	js.Rewrite("$<.Areas")
	return areas
}

// Name Sets or retrieves the name of the object.
func (*HTMLMapElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the name of the object.
func (*HTMLMapElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}
