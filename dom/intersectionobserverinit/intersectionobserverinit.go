package intersectionobserverinit

import "github.com/matthewmueller/golly/dom/window"

type IntersectionObserverInit struct {
	root       *window.Element
	rootMargin *string
	threshold  *interface{}
}
