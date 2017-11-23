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

// InitData
func (*MediaEncryptedEvent) InitData() (initData []byte) {
	js.Rewrite("$<.InitData")
	return initData
}

// InitDataType
func (*MediaEncryptedEvent) InitDataType() (initDataType string) {
	js.Rewrite("$<.InitDataType")
	return initDataType
}
