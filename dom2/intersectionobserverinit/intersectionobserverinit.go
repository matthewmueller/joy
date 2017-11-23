package intersectionobserverinit

import "github.com/matthewmueller/golly/dom2/window"

type IntersectionObserverInit struct {
	root       *window.Element
	rootMargin *string
	threshold  *interface{}
}
