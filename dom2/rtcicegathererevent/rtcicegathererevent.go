package rtcicegathererevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"RTCIceGathererEvent,omit"
type RTCIceGathererEvent struct {
	window.Event
}

// Candidate
func (*RTCIceGathererEvent) Candidate() (candidate interface{}) {
	js.Rewrite("$<.Candidate")
	return candidate
}
