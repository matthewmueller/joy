package window

import "github.com/matthewmueller/golly/js"

// js:"Touch,omit"
type Touch struct {
}

// ClientX
func (*Touch) ClientX() (clientX int) {
	js.Rewrite("$<.ClientX")
	return clientX
}

// ClientY
func (*Touch) ClientY() (clientY int) {
	js.Rewrite("$<.ClientY")
	return clientY
}

// Identifier
func (*Touch) Identifier() (identifier int) {
	js.Rewrite("$<.Identifier")
	return identifier
}

// PageX
func (*Touch) PageX() (pageX int) {
	js.Rewrite("$<.PageX")
	return pageX
}

// PageY
func (*Touch) PageY() (pageY int) {
	js.Rewrite("$<.PageY")
	return pageY
}

// ScreenX
func (*Touch) ScreenX() (screenX int) {
	js.Rewrite("$<.ScreenX")
	return screenX
}

// ScreenY
func (*Touch) ScreenY() (screenY int) {
	js.Rewrite("$<.ScreenY")
	return screenY
}

// Target
func (*Touch) Target() (target EventTarget) {
	js.Rewrite("$<.Target")
	return target
}
