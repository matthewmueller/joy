package positionerror

import "github.com/matthewmueller/golly/js"

// PositionError struct
// js:"PositionError,omit"
type PositionError struct {
}

// ToString fn
func (*PositionError) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

// Code prop
func (*PositionError) Code() (code uint8) {
	js.Rewrite("$<.code")
	return code
}

// Message prop
func (*PositionError) Message() (message string) {
	js.Rewrite("$<.message")
	return message
}
