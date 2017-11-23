package msmediakeymessageevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// MSMediaKeyMessageEvent struct
// js:"MSMediaKeyMessageEvent,omit"
type MSMediaKeyMessageEvent struct {
	window.Event
}

// DestinationURL
func (*MSMediaKeyMessageEvent) DestinationURL() (destinationURL string) {
	js.Rewrite("$<.DestinationURL")
	return destinationURL
}

// Message
func (*MSMediaKeyMessageEvent) Message() (message []uint8) {
	js.Rewrite("$<.Message")
	return message
}
