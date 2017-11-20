package windowbase64

import "github.com/matthewmueller/golly/js"

type WindowBase64 struct {
}

func (*WindowBase64) Atob(encodedString string) (s string) {
	js.Rewrite("$<.atob($1)", encodedString)
	return s
}

func (*WindowBase64) Btoa(rawString string) (s string) {
	js.Rewrite("$<.btoa($1)", rawString)
	return s
}
