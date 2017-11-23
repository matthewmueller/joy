package navigationevent

import "github.com/matthewmueller/golly/dom2/window"

// js:"NavigationEvent,omit"
type NavigationEvent interface {
	window.Event

	// URI prop
	// js:"uri"
	URI() (uri string)
}
