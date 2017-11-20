package intersectionobserverentry

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/clientrect"
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/js"
)

type IntersectionObserverEntry struct {
}

func (*IntersectionObserverEntry) GetBoundingClientRect() (boundingClientRect *clientrect.ClientRect) {
	js.Rewrite("$<.boundingClientRect")
	return boundingClientRect
}

func (*IntersectionObserverEntry) GetIntersectionRatio() (intersectionRatio float32) {
	js.Rewrite("$<.intersectionRatio")
	return intersectionRatio
}

func (*IntersectionObserverEntry) GetIntersectionRect() (intersectionRect *clientrect.ClientRect) {
	js.Rewrite("$<.intersectionRect")
	return intersectionRect
}

func (*IntersectionObserverEntry) GetRootBounds() (rootBounds *clientrect.ClientRect) {
	js.Rewrite("$<.rootBounds")
	return rootBounds
}

func (*IntersectionObserverEntry) GetTarget() (target *element.Element) {
	js.Rewrite("$<.target")
	return target
}

func (*IntersectionObserverEntry) GetTime() (time int) {
	js.Rewrite("$<.time")
	return time
}
