package headers

import (
	"github.com/matthewmueller/golly/dom/mediakeystatus"
	"github.com/matthewmueller/golly/js"
)

// Headers struct
// js:"Headers,omit"
type Headers struct {
}

// Append fn
// js:"append"
func (*Headers) Append(name string, value string) {
	js.Rewrite("$_.append($1, $2)", name, value)
}

// Delete fn
// js:"delete"
func (*Headers) Delete(name string) {
	js.Rewrite("$_.delete($1)", name)
}

// ForEach fn
// js:"forEach"
func (*Headers) ForEach(callback func(keyId []byte, status *mediakeystatus.MediaKeyStatus)) {
	js.Rewrite("$_.forEach($1)", callback)
}

// Get fn
// js:"get"
func (*Headers) Get(name string) (s string) {
	js.Rewrite("$_.get($1)", name)
	return s
}

// Has fn
// js:"has"
func (*Headers) Has(name string) (b bool) {
	js.Rewrite("$_.has($1)", name)
	return b
}

// Set fn
// js:"set"
func (*Headers) Set(name string, value string) {
	js.Rewrite("$_.set($1, $2)", name, value)
}
