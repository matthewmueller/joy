package focusnavigationevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/navigationreason"
	"github.com/matthewmueller/golly/js"
)

type FocusNavigationEvent struct {
	*event.Event
}

func (*FocusNavigationEvent) RequestFocus() {
	js.Rewrite("$<.requestFocus()")
}

func (*FocusNavigationEvent) GetNavigationReason() (navigationReason *navigationreason.NavigationReason) {
	js.Rewrite("$<.navigationReason")
	return navigationReason
}

func (*FocusNavigationEvent) GetOriginHeight() (originHeight float32) {
	js.Rewrite("$<.originHeight")
	return originHeight
}

func (*FocusNavigationEvent) GetOriginLeft() (originLeft float32) {
	js.Rewrite("$<.originLeft")
	return originLeft
}

func (*FocusNavigationEvent) GetOriginTop() (originTop float32) {
	js.Rewrite("$<.originTop")
	return originTop
}

func (*FocusNavigationEvent) GetOriginWidth() (originWidth float32) {
	js.Rewrite("$<.originWidth")
	return originWidth
}
