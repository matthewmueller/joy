package offlineaudiocompletionevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"OfflineAudioCompletionEvent,omit"
type OfflineAudioCompletionEvent struct {
	window.Event
}

// RenderedBuffer
func (*OfflineAudioCompletionEvent) RenderedBuffer() (renderedBuffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.RenderedBuffer")
	return renderedBuffer
}
