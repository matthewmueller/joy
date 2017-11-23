package idb

import (
	"github.com/matthewmueller/golly/dom2/hashchangeeventinit"
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
