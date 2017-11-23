package audioprocessingevent

import (
	"github.com/matthewmueller/golly/dom2/audiobuffer"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// AudioProcessingEvent struct
// js:"AudioProcessingEvent,omit"
type AudioProcessingEvent struct {
	window.Event
}

// InputBuffer prop
func (*AudioProcessingEvent) InputBuffer() (inputBuffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.inputBuffer")
	return inputBuffer
}

// OutputBuffer prop
func (*AudioProcessingEvent) OutputBuffer() (outputBuffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.outputBuffer")
	return outputBuffer
}

// PlaybackTime prop
func (*AudioProcessingEvent) PlaybackTime() (playbackTime float32) {
	js.Rewrite("$<.playbackTime")
	return playbackTime
}
