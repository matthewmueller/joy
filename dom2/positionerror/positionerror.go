package positionerror

import "github.com/matthewmueller/golly/js"

// js:"PositionError,omit"
type PositionError struct {
}

// ToString
func (*PositionError) ToString() (s string) {
	js.Rewrite("$<.ToString()")
	return s
}

// Code
func (*PositionError) Code() (code uint8) {
	js.Rewrite("$<.Code")
	return code
}

// Message
func (*PositionError) Message() (message string) {
	js.Rewrite("$<.Message")
	return message
}
