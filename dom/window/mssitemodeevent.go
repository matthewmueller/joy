package window

import "github.com/matthewmueller/golly/js"

// MSSiteModeEvent struct
// js:"MSSiteModeEvent,omit"
type MSSiteModeEvent struct {
	Event
}

// ActionURL prop
func (*MSSiteModeEvent) ActionURL() (actionURL string) {
	js.Rewrite("$<.actionURL")
	return actionURL
}

// ButtonID prop
func (*MSSiteModeEvent) ButtonID() (buttonID int) {
	js.Rewrite("$<.buttonID")
	return buttonID
}
