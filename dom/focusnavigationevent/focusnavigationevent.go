package focusnavigationevent

import (
	"github.com/matthewmueller/golly/dom/focusnavigationeventinit"
	"github.com/matthewmueller/golly/dom/navigationreason"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(kind string, eventinitdict *focusnavigationeventinit.FocusNavigationEventInit) *FocusNavigationEvent {
	js.Rewrite("FocusNavigationEvent")
	return &FocusNavigationEvent{}
}

// FocusNavigationEvent struct
// js:"FocusNavigationEvent,omit"
type FocusNavigationEvent struct {
	window.Event
}

// RequestFocus fn
func (*FocusNavigationEvent) RequestFocus() {
	js.Rewrite("$<.requestFocus()")
}

// NavigationReason prop
func (*FocusNavigationEvent) NavigationReason() (navigationReason *navigationreason.NavigationReason) {
	js.Rewrite("$<.navigationReason")
	return navigationReason
}

// OriginHeight prop
func (*FocusNavigationEvent) OriginHeight() (originHeight float32) {
	js.Rewrite("$<.originHeight")
	return originHeight
}

// OriginLeft prop
func (*FocusNavigationEvent) OriginLeft() (originLeft float32) {
	js.Rewrite("$<.originLeft")
	return originLeft
}

// OriginTop prop
func (*FocusNavigationEvent) OriginTop() (originTop float32) {
	js.Rewrite("$<.originTop")
	return originTop
}

// OriginWidth prop
func (*FocusNavigationEvent) OriginWidth() (originWidth float32) {
	js.Rewrite("$<.originWidth")
	return originWidth
}
