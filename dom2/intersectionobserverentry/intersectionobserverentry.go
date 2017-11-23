package intersectionobserverentry

import (
	"github.com/matthewmueller/golly/dom2/clientrect"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"IntersectionObserverEntry,omit"
type IntersectionObserverEntry struct {
}

// BoundingClientRect
func (*IntersectionObserverEntry) BoundingClientRect() (boundingClientRect *clientrect.ClientRect) {
	js.Rewrite("$<.BoundingClientRect")
	return boundingClientRect
}

// IntersectionRatio
func (*IntersectionObserverEntry) IntersectionRatio() (intersectionRatio float32) {
	js.Rewrite("$<.IntersectionRatio")
	return intersectionRatio
}

// IntersectionRect
func (*IntersectionObserverEntry) IntersectionRect() (intersectionRect *clientrect.ClientRect) {
	js.Rewrite("$<.IntersectionRect")
	return intersectionRect
}

// RootBounds
func (*IntersectionObserverEntry) RootBounds() (rootBounds *clientrect.ClientRect) {
	js.Rewrite("$<.RootBounds")
	return rootBounds
}

// Target
func (*IntersectionObserverEntry) Target() (target window.Element) {
	js.Rewrite("$<.Target")
	return target
}

// Time
func (*IntersectionObserverEntry) Time() (time int) {
	js.Rewrite("$<.Time")
	return time
}
