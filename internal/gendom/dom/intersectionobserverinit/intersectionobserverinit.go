package intersectionobserverinit

import "github.com/matthewmueller/golly/internal/gendom/dom/element"

type IntersectionObserverInit struct {
	root       *element.Element
	rootMargin *string
	threshold  *interface{}
}
