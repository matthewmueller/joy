package intersectionobserverentryinit

import (
	"github.com/matthewmueller/joy/dom/domrectinit"
	"github.com/matthewmueller/joy/dom/window"
)

type IntersectionObserverEntryInit struct {
	boundingClientRect *domrectinit.DOMRectInit
	intersectionRect   *domrectinit.DOMRectInit
	rootBounds         *domrectinit.DOMRectInit
	target             window.Element
	time               int
}
