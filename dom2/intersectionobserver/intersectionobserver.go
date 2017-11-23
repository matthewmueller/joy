package intersectionobserver

import (
	"github.com/matthewmueller/golly/dom2/intersectionobserverentry"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"IntersectionObserver,omit"
type IntersectionObserver struct {
}

// Disconnect
func (*IntersectionObserver) Disconnect() {
	js.Rewrite("$<.Disconnect()")
}

// Observe
func (*IntersectionObserver) Observe(target window.Element) {
	js.Rewrite("$<.Observe($1)", target)
}

// TakeRecords
func (*IntersectionObserver) TakeRecords() (i []*intersectionobserverentry.IntersectionObserverEntry) {
	js.Rewrite("$<.TakeRecords()")
	return i
}

// Unobserve
func (*IntersectionObserver) Unobserve(target window.Element) {
	js.Rewrite("$<.Unobserve($1)", target)
}

// Root
func (*IntersectionObserver) Root() (root window.Element) {
	js.Rewrite("$<.Root")
	return root
}

// RootMargin
func (*IntersectionObserver) RootMargin() (rootMargin string) {
	js.Rewrite("$<.RootMargin")
	return rootMargin
}

// Thresholds
func (*IntersectionObserver) Thresholds() (thresholds []float32) {
	js.Rewrite("$<.Thresholds")
	return thresholds
}
