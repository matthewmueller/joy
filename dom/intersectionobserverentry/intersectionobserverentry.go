package intersectionobserverentry

import (
	"github.com/matthewmueller/joy/dom/clientrect"
	"github.com/matthewmueller/joy/dom/intersectionobserverentryinit"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
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
// js:"boundingClientRect"
func (*IntersectionObserverEntry) BoundingClientRect() (boundingClientRect *clientrect.ClientRect) {
	js.Rewrite("$_.boundingClientRect")
	return boundingClientRect
}

// IntersectionRatio prop
// js:"intersectionRatio"
func (*IntersectionObserverEntry) IntersectionRatio() (intersectionRatio float32) {
	js.Rewrite("$_.intersectionRatio")
	return intersectionRatio
}

// IntersectionRect prop
// js:"intersectionRect"
func (*IntersectionObserverEntry) IntersectionRect() (intersectionRect *clientrect.ClientRect) {
	js.Rewrite("$_.intersectionRect")
	return intersectionRect
}

// RootBounds prop
// js:"rootBounds"
func (*IntersectionObserverEntry) RootBounds() (rootBounds *clientrect.ClientRect) {
	js.Rewrite("$_.rootBounds")
	return rootBounds
}

// Target prop
// js:"target"
func (*IntersectionObserverEntry) Target() (target window.Element) {
	js.Rewrite("$_.target")
	return target
}

// Time prop
// js:"time"
func (*IntersectionObserverEntry) Time() (time int) {
	js.Rewrite("$_.time")
	return time
}
