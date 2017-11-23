package focusnavigationevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"FocusNavigationEvent,omit"
type FocusNavigationEvent struct {
	window.Event
}

// RequestFocus
func (*FocusNavigationEvent) RequestFocus() {
	js.Rewrite("$<.RequestFocus()")
}

// NavigationReason
func (*FocusNavigationEvent) NavigationReason() (navigationReason *navigationreason.NavigationReason) {
	js.Rewrite("$<.NavigationReason")
	return navigationReason
}

// OriginHeight
func (*FocusNavigationEvent) OriginHeight() (originHeight float32) {
	js.Rewrite("$<.OriginHeight")
	return originHeight
}

// OriginLeft
func (*FocusNavigationEvent) OriginLeft() (originLeft float32) {
	js.Rewrite("$<.OriginLeft")
	return originLeft
}

// OriginTop
func (*FocusNavigationEvent) OriginTop() (originTop float32) {
	js.Rewrite("$<.OriginTop")
	return originTop
}

// OriginWidth
func (*FocusNavigationEvent) OriginWidth() (originWidth float32) {
	js.Rewrite("$<.OriginWidth")
	return originWidth
}
