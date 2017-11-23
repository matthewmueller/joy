package intersectionobserverentry

import (
	"github.com/matthewmueller/golly/dom/clientrect"
	"github.com/matthewmueller/golly/dom/intersectionobserverentryinit"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(intersectionobserverentryinit *intersectionobserverentryinit.IntersectionObserverEntryInit) *IntersectionObserverEntry {
	js.Rewrite("IntersectionObserverEntry")
	return &IntersectionObserverEntry{}
}

// IntersectionObserverEntry struct
// js:"IntersectionObserverEntry,omit"
type IntersectionObserverEntry struct {
}

// BoundingClientRect prop
func (*IntersectionObserverEntry) BoundingClientRect() (boundingClientRect *clientrect.ClientRect) {
	js.Rewrite("$<.boundingClientRect")
	return boundingClientRect
}

// IntersectionRatio prop
func (*IntersectionObserverEntry) IntersectionRatio() (intersectionRatio float32) {
	js.Rewrite("$<.intersectionRatio")
	return intersectionRatio
}

// IntersectionRect prop
func (*IntersectionObserverEntry) IntersectionRect() (intersectionRect *clientrect.ClientRect) {
	js.Rewrite("$<.intersectionRect")
	return intersectionRect
}

// RootBounds prop
func (*IntersectionObserverEntry) RootBounds() (rootBounds *clientrect.ClientRect) {
	js.Rewrite("$<.rootBounds")
	return rootBounds
}

// Target prop
func (*IntersectionObserverEntry) Target() (target window.Element) {
	js.Rewrite("$<.target")
	return target
}

// Time prop
func (*IntersectionObserverEntry) Time() (time int) {
	js.Rewrite("$<.time")
	return time
}
