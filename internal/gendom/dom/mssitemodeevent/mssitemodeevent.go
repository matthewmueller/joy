package mssitemodeevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type MSSiteModeEvent struct {
	*event.Event
}

func (*MSSiteModeEvent) GetActionURL() (actionURL string) {
	js.Rewrite("$<.actionURL")
	return actionURL
}

func (*MSSiteModeEvent) GetButtonID() (buttonID int) {
	js.Rewrite("$<.buttonID")
	return buttonID
}
