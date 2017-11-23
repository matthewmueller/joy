package rtcdtmftonechangeevent

import (
	"github.com/matthewmueller/golly/dom2/rtcdtmftonechangeeventinit"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(typearg string, eventinitdict *rtcdtmftonechangeeventinit.RTCDTMFToneChangeEventInit) *RTCDTMFToneChangeEvent {
	js.Rewrite("RTCDTMFToneChangeEvent")
	return &RTCDTMFToneChangeEvent{}
}

// RTCDTMFToneChangeEvent struct
// js:"RTCDTMFToneChangeEvent,omit"
type RTCDTMFToneChangeEvent struct {
	window.Event
}

// Tone
func (*RTCDTMFToneChangeEvent) Tone() (tone string) {
	js.Rewrite("$<.Tone")
	return tone
}
