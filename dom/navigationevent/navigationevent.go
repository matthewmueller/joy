package navigationevent

import "github.com/matthewmueller/golly/dom/window"

// js:"NavigationEvent,omit"
type NavigationEvent interface {
	window.Event

	// URI prop
	// js:"uri"
	URI() (uri string)
}
