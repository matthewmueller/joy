package window

import "github.com/matthewmueller/joy/macro"

// Touch struct
// js:"Touch,omit"
type Touch struct {
}

// ClientX prop
// js:"clientX"
func (*Touch) ClientX() (clientX int) {
	macro.Rewrite("$_.clientX")
	return clientX
}

// ClientY prop
// js:"clientY"
func (*Touch) ClientY() (clientY int) {
	macro.Rewrite("$_.clientY")
	return clientY
}

// Identifier prop
// js:"identifier"
func (*Touch) Identifier() (identifier int) {
	macro.Rewrite("$_.identifier")
	return identifier
}

// PageX prop
// js:"pageX"
func (*Touch) PageX() (pageX int) {
	macro.Rewrite("$_.pageX")
	return pageX
}

// PageY prop
// js:"pageY"
func (*Touch) PageY() (pageY int) {
	macro.Rewrite("$_.pageY")
	return pageY
}

// ScreenX prop
// js:"screenX"
func (*Touch) ScreenX() (screenX int) {
	macro.Rewrite("$_.screenX")
	return screenX
}

// ScreenY prop
// js:"screenY"
func (*Touch) ScreenY() (screenY int) {
	macro.Rewrite("$_.screenY")
	return screenY
}

// Target prop
// js:"target"
func (*Touch) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}
