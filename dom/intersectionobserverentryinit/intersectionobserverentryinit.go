package intersectionobserverentryinit

import (
	"github.com/matthewmueller/golly/dom/domrectinit"
	"github.com/matthewmueller/golly/dom/window"
)

type IntersectionObserverEntryInit struct {
	boundingClientRect *domrectinit.DOMRectInit
	intersectionRect   *domrectinit.DOMRectInit
	rootBounds         *domrectinit.DOMRectInit
	target             window.Element
	time               int
}
