package mediakeymessageevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

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
