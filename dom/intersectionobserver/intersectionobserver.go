package intersectionobserver

import (
	"github.com/matthewmueller/golly/dom/intersectionobserverentry"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// IntersectionObserver struct
// js:"IntersectionObserver,omit"
type IntersectionObserver struct {
}

// Disconnect fn
// js:"disconnect"
func (*IntersectionObserver) Disconnect() {
	js.Rewrite("$_.disconnect()")
}

// Observe fn
// js:"observe"
func (*IntersectionObserver) Observe(target window.Element) {
	js.Rewrite("$_.observe($1)", target)
}

// TakeRecords fn
// js:"takeRecords"
func (*IntersectionObserver) TakeRecords() (i []*intersectionobserverentry.IntersectionObserverEntry) {
	js.Rewrite("$_.takeRecords()")
	return i
}

// Unobserve fn
// js:"unobserve"
func (*IntersectionObserver) Unobserve(target window.Element) {
	js.Rewrite("$_.unobserve($1)", target)
}

// Root prop
// js:"root"
func (*IntersectionObserver) Root() (root window.Element) {
	js.Rewrite("$_.root")
	return root
}

// RootMargin prop
// js:"rootMargin"
func (*IntersectionObserver) RootMargin() (rootMargin string) {
	js.Rewrite("$_.rootMargin")
	return rootMargin
}

// Thresholds prop
// js:"thresholds"
func (*IntersectionObserver) Thresholds() (thresholds []float32) {
	js.Rewrite("$_.thresholds")
	return thresholds
}
