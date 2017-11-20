package uievent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type UIEvent struct {
	*event.Event
}

func (*UIEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int) {
	js.Rewrite("$<.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*UIEvent) GetDetail() (detail int) {
	js.Rewrite("$<.detail")
	return detail
}

func (*UIEvent) GetView() (view *window.Window) {
	js.Rewrite("$<.view")
	return view
}
