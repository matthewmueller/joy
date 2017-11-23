package htmldatalistelement

import "github.com/matthewmueller/golly/js"

// js:"HTMLDataListElement,omit"
type HTMLDataListElement struct {
	window.HTMLElement
}

// Options
func (*HTMLDataListElement) Options() (options window.HTMLCollection) {
	js.Rewrite("$<.Options")
	return options
}
