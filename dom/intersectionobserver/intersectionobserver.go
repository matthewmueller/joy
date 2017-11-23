package intersectionobserver

import (
	"github.com/matthewmueller/golly/dom/intersectionobserverentry"
	"github.com/matthewmueller/golly/dom/intersectionobserverinit"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(callback func(entries []*intersectionobserverentry.IntersectionObserverEntry, observer *IntersectionObserver), options *intersectionobserverinit.IntersectionObserverInit) *IntersectionObserver {
	js.Rewrite("IntersectionObserver")
	return &IntersectionObserver{}
}

// IntersectionObserver struct
// js:"IntersectionObserver,omit"
type IntersectionObserver struct {
}

// Disconnect fn
func (*IntersectionObserver) Disconnect() {
	js.Rewrite("$<.disconnect()")
}

// Observe fn
func (*IntersectionObserver) Observe(target window.Element) {
	js.Rewrite("$<.observe($1)", target)
}

// TakeRecords fn
func (*IntersectionObserver) TakeRecords() (i []*intersectionobserverentry.IntersectionObserverEntry) {
	js.Rewrite("$<.takeRecords()")
	return i
}

// Unobserve fn
func (*IntersectionObserver) Unobserve(target window.Element) {
	js.Rewrite("$<.unobserve($1)", target)
}

// Root prop
func (*IntersectionObserver) Root() (root window.Element) {
	js.Rewrite("$<.root")
	return root
}

// RootMargin prop
func (*IntersectionObserver) RootMargin() (rootMargin string) {
	js.Rewrite("$<.rootMargin")
	return rootMargin
}

// Thresholds prop
func (*IntersectionObserver) Thresholds() (thresholds []float32) {
	js.Rewrite("$<.thresholds")
	return thresholds
}
