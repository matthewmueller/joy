package positionerror

import "github.com/matthewmueller/golly/js"

// PositionError struct
// js:"PositionError,omit"
type PositionError struct {
}

// ToString fn
// js:"toString"
func (*PositionError) ToString() (s string) {
	js.Rewrite("$_.toString()")
	return s
}

// Code prop
// js:"code"
func (*PositionError) Code() (code uint8) {
	js.Rewrite("$_.code")
	return code
}

// Message prop
// js:"message"
func (*PositionError) Message() (message string) {
	js.Rewrite("$_.message")
	return message
}
