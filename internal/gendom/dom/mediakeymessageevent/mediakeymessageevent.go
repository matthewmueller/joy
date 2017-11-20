package mediakeymessageevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediakeymessagetype"
	"github.com/matthewmueller/golly/js"
)

type MediaKeyMessageEvent struct {
	*event.Event
}

func (*MediaKeyMessageEvent) GetMessage() (message []byte) {
	js.Rewrite("$<.message")
	return message
}

func (*MediaKeyMessageEvent) GetMessageType() (messageType *mediakeymessagetype.MediaKeyMessageType) {
	js.Rewrite("$<.messageType")
	return messageType
}
