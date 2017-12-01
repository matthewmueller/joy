package domerror

import "github.com/matthewmueller/golly/js"

// DOMError struct
// js:"DOMError,omit"
type DOMError struct {
}

// ToString fn
// js:"toString"
func (*DOMError) ToString() (s string) {
	js.Rewrite("$_.toString()")
	return s
}

// Name prop
// js:"name"
func (*DOMError) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}
