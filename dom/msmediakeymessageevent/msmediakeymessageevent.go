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

// DestinationURL prop
func (*MSMediaKeyMessageEvent) DestinationURL() (destinationURL string) {
	js.Rewrite("$<.destinationURL")
	return destinationURL
}

// Message prop
func (*MSMediaKeyMessageEvent) Message() (message []uint8) {
	js.Rewrite("$<.message")
	return message
}
