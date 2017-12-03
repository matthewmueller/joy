package positionerror

import "github.com/matthewmueller/joy/macro"

// PositionError struct
// js:"PositionError,omit"
type PositionError struct {
}

// ToString fn
// js:"toString"
func (*PositionError) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

// Code prop
// js:"code"
func (*PositionError) Code() (code uint8) {
	macro.Rewrite("$_.code")
	return code
}

// Message prop
// js:"message"
func (*PositionError) Message() (message string) {
	macro.Rewrite("$_.message")
	return message
}
