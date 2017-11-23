package windowbase64

import "github.com/matthewmueller/golly/js"

// WindowBase64 struct
// js:"WindowBase64,omit"
type WindowBase64 struct {
}

// Atob
func (*WindowBase64) Atob(encodedString string) (s string) {
	js.Rewrite("$<.Atob($1)", encodedString)
	return s
}

// Btoa
func (*WindowBase64) Btoa(rawString string) (s string) {
	js.Rewrite("$<.Btoa($1)", rawString)
	return s
}
