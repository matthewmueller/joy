package rtcdtmftonechangeevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"RTCDTMFToneChangeEvent,omit"
type RTCDTMFToneChangeEvent struct {
	window.Event
}

// Tone
func (*RTCDTMFToneChangeEvent) Tone() (tone string) {
	js.Rewrite("$<.Tone")
	return tone
}
