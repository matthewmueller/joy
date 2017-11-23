package domexception

import "github.com/matthewmueller/golly/js"

// DOMException struct
// js:"DOMException,omit"
type DOMException struct {
}

// ToString fn
func (*DOMException) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

// Code prop
func (*DOMException) Code() (code uint8) {
	js.Rewrite("$<.code")
	return code
}

// Message prop
func (*DOMException) Message() (message string) {
	js.Rewrite("$<.message")
	return message
}

// Name prop
func (*DOMException) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}
