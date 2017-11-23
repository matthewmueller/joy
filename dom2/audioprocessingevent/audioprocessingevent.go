package audioprocessingevent

import (
	"github.com/matthewmueller/golly/dom2/audiobuffer"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"AudioProcessingEvent,omit"
type AudioProcessingEvent struct {
	window.Event
}

// InputBuffer
func (*AudioProcessingEvent) InputBuffer() (inputBuffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.InputBuffer")
	return inputBuffer
}

// OutputBuffer
func (*AudioProcessingEvent) OutputBuffer() (outputBuffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.OutputBuffer")
	return outputBuffer
}

// PlaybackTime
func (*AudioProcessingEvent) PlaybackTime() (playbackTime float32) {
	js.Rewrite("$<.PlaybackTime")
	return playbackTime
}
