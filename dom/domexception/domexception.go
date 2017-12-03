package domexception

import "github.com/matthewmueller/joy/js"

// DOMException struct
// js:"DOMException,omit"
type DOMException struct {
}

// ToString fn
// js:"toString"
func (*DOMException) ToString() (s string) {
	js.Rewrite("$_.toString()")
	return s
}

// Code prop
// js:"code"
func (*DOMException) Code() (code uint8) {
	js.Rewrite("$_.code")
	return code
}

// Message prop
// js:"message"
func (*DOMException) Message() (message string) {
	js.Rewrite("$_.message")
	return message
}

// Name prop
// js:"name"
func (*DOMException) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}
