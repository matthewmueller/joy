package intersectionobserverentryinit

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/domrectinit"
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
)

type IntersectionObserverEntryInit struct {
	boundingClientRect *domrectinit.DOMRectInit
	intersectionRect   *domrectinit.DOMRectInit
	rootBounds         *domrectinit.DOMRectInit
	target             *element.Element
	time               int
}
