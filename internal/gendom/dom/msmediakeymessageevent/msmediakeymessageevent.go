package msmediakeymessageevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type MSMediaKeyMessageEvent struct {
	*event.Event
}

func (*MSMediaKeyMessageEvent) GetDestinationURL() (destinationURL string) {
	js.Rewrite("$<.destinationURL")
	return destinationURL
}

func (*MSMediaKeyMessageEvent) GetMessage() (message []uint8) {
	js.Rewrite("$<.message")
	return message
}
