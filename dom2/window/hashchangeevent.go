package window

import "github.com/matthewmueller/golly/js"

// js:"HashChangeEvent,omit"
type HashChangeEvent struct {
	Event
}

// NewURL
func (*HashChangeEvent) NewURL() (newURL string) {
	js.Rewrite("$<.NewURL")
	return newURL
}

// OldURL
func (*HashChangeEvent) OldURL() (oldURL string) {
	js.Rewrite("$<.OldURL")
	return oldURL
}
