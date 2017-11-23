package mediakeymessageevent

import (
	"github.com/matthewmueller/golly/dom2/mediakeymessageeventinit"
	"github.com/matthewmueller/golly/dom2/mediakeymessagetype"
	"github.com/matthewmueller/golly/dom2/window"
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

// Message
func (*MediaKeyMessageEvent) Message() (message []byte) {
	js.Rewrite("$<.Message")
	return message
}

// MessageType
func (*MediaKeyMessageEvent) MessageType() (messageType *mediakeymessagetype.MediaKeyMessageType) {
	js.Rewrite("$<.MessageType")
	return messageType
}
