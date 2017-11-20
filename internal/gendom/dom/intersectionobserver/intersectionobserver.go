package intersectionobserver

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/intersectionobserverentry"
	"github.com/matthewmueller/golly/js"
)

type IntersectionObserver struct {
}

func (*IntersectionObserver) Disconnect() {
	js.Rewrite("$<.disconnect()")
}

func (*IntersectionObserver) Observe(target *element.Element) {
	js.Rewrite("$<.observe($1)", target)
}

func (*IntersectionObserver) TakeRecords() (i []*intersectionobserverentry.IntersectionObserverEntry) {
	js.Rewrite("$<.takeRecords()")
	return i
}

func (*IntersectionObserver) Unobserve(target *element.Element) {
	js.Rewrite("$<.unobserve($1)", target)
}

func (*IntersectionObserver) GetRoot() (root *element.Element) {
	js.Rewrite("$<.root")
	return root
}

func (*IntersectionObserver) GetRootMargin() (rootMargin string) {
	js.Rewrite("$<.rootMargin")
	return rootMargin
}

func (*IntersectionObserver) GetThresholds() (thresholds []float32) {
	js.Rewrite("$<.thresholds")
	return thresholds
}
