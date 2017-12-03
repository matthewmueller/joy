package domexception

import "github.com/matthewmueller/joy/macro"

// DOMException struct
// js:"DOMException,omit"
type DOMException struct {
}

// ToString fn
// js:"toString"
func (*DOMException) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

// Code prop
// js:"code"
func (*DOMException) Code() (code uint8) {
	macro.Rewrite("$_.code")
	return code
}

// Message prop
// js:"message"
func (*DOMException) Message() (message string) {
	macro.Rewrite("$_.message")
	return message
}

// Name prop
// js:"name"
func (*DOMException) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}
