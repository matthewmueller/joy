package htmlmapelement

import (
	"github.com/matthewmueller/golly/dom2/htmlareascollection"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLMapElement struct
// js:"HTMLMapElement,omit"
type HTMLMapElement struct {
	window.HTMLElement
}

// Areas prop Retrieves a collection of the area objects defined for the given map object.
func (*HTMLMapElement) Areas() (areas *htmlareascollection.HTMLAreasCollection) {
	js.Rewrite("$<.areas")
	return areas
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLMapElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLMapElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}
