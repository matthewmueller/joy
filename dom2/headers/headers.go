package headers

import (
	"github.com/matthewmueller/golly/dom2/mediakeystatus"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(init *interface{}) *Headers {
	js.Rewrite("Headers")
	return &Headers{}
}

// Headers struct
// js:"Headers,omit"
type Headers struct {
}

// Append
func (*Headers) Append(name string, value string) {
	js.Rewrite("$<.Append($1, $2)", name, value)
}

// Delete
func (*Headers) Delete(name string) {
	js.Rewrite("$<.Delete($1)", name)
}

// ForEach
func (*Headers) ForEach(callback func(keyId []byte, status *mediakeystatus.MediaKeyStatus)) {
	js.Rewrite("$<.ForEach($1)", callback)
}

// Get
func (*Headers) Get(name string) (s string) {
	js.Rewrite("$<.Get($1)", name)
	return s
}

// Has
func (*Headers) Has(name string) (b bool) {
	js.Rewrite("$<.Has($1)", name)
	return b
}

// Set
func (*Headers) Set(name string, value string) {
	js.Rewrite("$<.Set($1, $2)", name, value)
}
