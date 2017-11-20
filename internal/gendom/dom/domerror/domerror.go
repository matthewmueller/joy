package domerror

import "github.com/matthewmueller/golly/js"

type DOMError struct {
}

func (*DOMError) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*DOMError) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}
