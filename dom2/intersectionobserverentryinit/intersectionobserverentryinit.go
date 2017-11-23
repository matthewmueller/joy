package intersectionobserverentryinit

import (
	"github.com/matthewmueller/golly/dom2/domrectinit"
	"github.com/matthewmueller/golly/dom2/window"
)

type IntersectionObserverEntryInit struct {
	boundingClientRect *domrectinit.DOMRectInit
	intersectionRect   *domrectinit.DOMRectInit
	rootBounds         *domrectinit.DOMRectInit
	target             window.Element
	time               int
}
