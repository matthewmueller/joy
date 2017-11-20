package positionerror

import "github.com/matthewmueller/golly/js"

type PositionError struct {
}

func (*PositionError) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*PositionError) GetCode() (code uint8) {
	js.Rewrite("$<.code")
	return code
}

func (*PositionError) GetMessage() (message string) {
	js.Rewrite("$<.message")
	return message
}
