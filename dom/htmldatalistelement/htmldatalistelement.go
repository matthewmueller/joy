package htmldatalistelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLDataListElement struct
// js:"HTMLDataListElement,omit"
type HTMLDataListElement struct {
	window.HTMLElement
}

// Options prop
func (*HTMLDataListElement) Options() (options window.HTMLCollection) {
	js.Rewrite("$<.options")
	return options
}
