package intersectionobserverinit

import "github.com/matthewmueller/joy/dom/window"

type IntersectionObserverInit struct {
	root       *window.Element
	rootMargin *string
	threshold  *interface{}
}
