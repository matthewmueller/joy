package mediaencryptedevent

import (
	"github.com/matthewmueller/golly/dom2/mediaencryptedeventinit"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(kind string, eventinitdict *mediaencryptedeventinit.MediaEncryptedEventInit) *MediaEncryptedEvent {
	js.Rewrite("MediaEncryptedEvent")
	return &MediaEncryptedEvent{}
}

// MediaEncryptedEvent struct
// js:"MediaEncryptedEvent,omit"
type MediaEncryptedEvent struct {
	window.Event
}

// InitData prop
func (*MediaEncryptedEvent) InitData() (initData []byte) {
	js.Rewrite("$<.initData")
	return initData
}

// InitDataType prop
func (*MediaEncryptedEvent) InitDataType() (initDataType string) {
	js.Rewrite("$<.initDataType")
	return initDataType
}
