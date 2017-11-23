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

// Append fn
func (*Headers) Append(name string, value string) {
	js.Rewrite("$<.append($1, $2)", name, value)
}

// Delete fn
func (*Headers) Delete(name string) {
	js.Rewrite("$<.delete($1)", name)
}

// ForEach fn
func (*Headers) ForEach(callback func(keyId []byte, status *mediakeystatus.MediaKeyStatus)) {
	js.Rewrite("$<.forEach($1)", callback)
}

// Get fn
func (*Headers) Get(name string) (s string) {
	js.Rewrite("$<.get($1)", name)
	return s
}

// Has fn
func (*Headers) Has(name string) (b bool) {
	js.Rewrite("$<.has($1)", name)
	return b
}

// Set fn
func (*Headers) Set(name string, value string) {
	js.Rewrite("$<.set($1, $2)", name, value)
}
