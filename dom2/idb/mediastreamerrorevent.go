package idb

import (
	"github.com/matthewmueller/golly/dom2/mediastreamerror"
	"github.com/matthewmueller/golly/dom2/mediastreamerroreventinit"
	"github.com/matthewmueller/golly/js"
)

// NewMediaStreamErrorEvent fn
func NewMediaStreamErrorEvent(typearg string, eventinitdict *mediastreamerroreventinit.MediaStreamErrorEventInit) *MediaStreamErrorEvent {
	js.Rewrite("MediaStreamErrorEvent")
	return &MediaStreamErrorEvent{}
}

// MediaStreamErrorEvent struct
// js:"MediaStreamErrorEvent,omit"
type MediaStreamErrorEvent struct {
	Event
}

// Error
func (*MediaStreamErrorEvent) Error() (err *mediastreamerror.MediaStreamError) {
	js.Rewrite("$<.Error")
	return err
}
