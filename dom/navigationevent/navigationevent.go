package navigationevent

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// js:"NavigationEvent,omit"
type NavigationEvent interface {
	window.Event

	// URI prop
	// js:"uri,rewrite=uri"
	URI() (uri string)
}

// uri prop
func uri() (uri string) {
	js.Rewrite("$<.uri")
	return uri
}
