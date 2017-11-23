package htmlformcontrolscollection

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLFormControlsCollection struct
// js:"HTMLFormControlsCollection,omit"
type HTMLFormControlsCollection struct {
	window.HTMLCollection
}

// NamedItem fn
func (*HTMLFormControlsCollection) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.namedItem($1)", name)
	return i
}
