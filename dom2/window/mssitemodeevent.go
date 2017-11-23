package window

import "github.com/matthewmueller/golly/js"

// MSSiteModeEvent struct
// js:"MSSiteModeEvent,omit"
type MSSiteModeEvent struct {
	Event
}

// ActionURL
func (*MSSiteModeEvent) ActionURL() (actionURL string) {
	js.Rewrite("$<.ActionURL")
	return actionURL
}

// ButtonID
func (*MSSiteModeEvent) ButtonID() (buttonID int) {
	js.Rewrite("$<.ButtonID")
	return buttonID
}
