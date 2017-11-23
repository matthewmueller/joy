package rtcicegathererevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// RTCIceGathererEvent struct
// js:"RTCIceGathererEvent,omit"
type RTCIceGathererEvent struct {
	window.Event
}

// Candidate prop
func (*RTCIceGathererEvent) Candidate() (candidate interface{}) {
	js.Rewrite("$<.candidate")
	return candidate
}
