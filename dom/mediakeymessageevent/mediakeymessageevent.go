package mediakeymessageevent

import (
	"github.com/matthewmueller/golly/dom/mediakeymessageeventinit"
	"github.com/matthewmueller/golly/dom/mediakeymessagetype"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(kind string, eventinitdict *mediakeymessageeventinit.MediaKeyMessageEventInit) *MediaKeyMessageEvent {
	js.Rewrite("MediaKeyMessageEvent")
	return &MediaKeyMessageEvent{}
}

// MediaKeyMessageEvent struct
// js:"MediaKeyMessageEvent,omit"
type MediaKeyMessageEvent struct {
	window.Event
}

// Message prop
func (*MediaKeyMessageEvent) Message() (message []byte) {
	js.Rewrite("$<.message")
	return message
}

// MessageType prop
func (*MediaKeyMessageEvent) MessageType() (messageType *mediakeymessagetype.MediaKeyMessageType) {
	js.Rewrite("$<.messageType")
	return messageType
}
