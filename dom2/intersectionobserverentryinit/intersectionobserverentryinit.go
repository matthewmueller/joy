package intersectionobserverentryinit

import "github.com/matthewmueller/golly/dom2/domrectinit"

type IntersectionObserverEntryInit struct {
	boundingClientRect *domrectinit.DOMRectInit
	intersectionRect   *domrectinit.DOMRectInit
	rootBounds         *domrectinit.DOMRectInit
	target             window.Element
	time               int
}
