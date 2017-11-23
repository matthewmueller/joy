package domerror

import "github.com/matthewmueller/golly/js"

// DOMError struct
// js:"DOMError,omit"
type DOMError struct {
}

// ToString
func (*DOMError) ToString() (s string) {
	js.Rewrite("$<.ToString()")
	return s
}

// Name
func (*DOMError) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}
