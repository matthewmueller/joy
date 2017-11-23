package domerror

import "github.com/matthewmueller/golly/js"

// DOMError struct
// js:"DOMError,omit"
type DOMError struct {
}

// ToString fn
func (*DOMError) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

// Name prop
func (*DOMError) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}
