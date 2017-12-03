package domerror

import "github.com/matthewmueller/joy/macro"

// DOMError struct
// js:"DOMError,omit"
type DOMError struct {
}

// ToString fn
// js:"toString"
func (*DOMError) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

// Name prop
// js:"name"
func (*DOMError) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}
