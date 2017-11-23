package window

import (
	"github.com/matthewmueller/golly/dom/hashchangeeventinit"
	"github.com/matthewmueller/golly/js"
)

// NewHashChangeEvent fn
func NewHashChangeEvent(typearg string, eventinitdict *hashchangeeventinit.HashChangeEventInit) *HashChangeEvent {
	js.Rewrite("HashChangeEvent")
	return &HashChangeEvent{}
}

// HashChangeEvent struct
// js:"HashChangeEvent,omit"
type HashChangeEvent struct {
	Event
}

// NewURL prop
func (*HashChangeEvent) NewURL() (newURL string) {
	js.Rewrite("$<.newURL")
	return newURL
}

// OldURL prop
func (*HashChangeEvent) OldURL() (oldURL string) {
	js.Rewrite("$<.oldURL")
	return oldURL
}
