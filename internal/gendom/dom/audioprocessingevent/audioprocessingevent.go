package audioprocessingevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audiobuffer"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type AudioProcessingEvent struct {
	*event.Event
}

func (*AudioProcessingEvent) GetInputBuffer() (inputBuffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.inputBuffer")
	return inputBuffer
}

func (*AudioProcessingEvent) GetOutputBuffer() (outputBuffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.outputBuffer")
	return outputBuffer
}

func (*AudioProcessingEvent) GetPlaybackTime() (playbackTime float32) {
	js.Rewrite("$<.playbackTime")
	return playbackTime
}
