package windowbase64

import "github.com/matthewmueller/golly/js"

// WindowBase64 struct
// js:"WindowBase64,omit"
type WindowBase64 struct {
}

// Atob fn
func (*WindowBase64) Atob(encodedString string) (s string) {
	js.Rewrite("$<.atob($1)", encodedString)
	return s
}

// Btoa fn
func (*WindowBase64) Btoa(rawString string) (s string) {
	js.Rewrite("$<.btoa($1)", rawString)
	return s
}
