package navigationevent

import "github.com/matthewmueller/joy/dom/window"

// NavigationEvent interface
// js:"NavigationEvent"
type NavigationEvent interface {
	window.Event

	// URI prop
	// js:"uri"
	// jsrewrite:"$_.uri"
	URI() (uri string)
}
