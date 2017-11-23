package window

import "github.com/matthewmueller/golly/js"

// Touch struct
// js:"Touch,omit"
type Touch struct {
}

// ClientX prop
func (*Touch) ClientX() (clientX int) {
	js.Rewrite("$<.clientX")
	return clientX
}

// ClientY prop
func (*Touch) ClientY() (clientY int) {
	js.Rewrite("$<.clientY")
	return clientY
}

// Identifier prop
func (*Touch) Identifier() (identifier int) {
	js.Rewrite("$<.identifier")
	return identifier
}

// PageX prop
func (*Touch) PageX() (pageX int) {
	js.Rewrite("$<.pageX")
	return pageX
}

// PageY prop
func (*Touch) PageY() (pageY int) {
	js.Rewrite("$<.pageY")
	return pageY
}

// ScreenX prop
func (*Touch) ScreenX() (screenX int) {
	js.Rewrite("$<.screenX")
	return screenX
}

// ScreenY prop
func (*Touch) ScreenY() (screenY int) {
	js.Rewrite("$<.screenY")
	return screenY
}

// Target prop
func (*Touch) Target() (target EventTarget) {
	js.Rewrite("$<.target")
	return target
}
