package domexception

import "github.com/matthewmueller/golly/js"

// js:"DOMException,omit"
type DOMException struct {
}

// ToString
func (*DOMException) ToString() (s string) {
	js.Rewrite("$<.ToString()")
	return s
}

// Code
func (*DOMException) Code() (code uint8) {
	js.Rewrite("$<.Code")
	return code
}

// Message
func (*DOMException) Message() (message string) {
	js.Rewrite("$<.Message")
	return message
}

// Name
func (*DOMException) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}
