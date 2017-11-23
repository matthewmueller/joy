package offlineaudiocompletionevent

import (
	"github.com/matthewmueller/golly/dom/audiobuffer"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// OfflineAudioCompletionEvent struct
// js:"OfflineAudioCompletionEvent,omit"
type OfflineAudioCompletionEvent struct {
	window.Event
}

// RenderedBuffer prop
func (*OfflineAudioCompletionEvent) RenderedBuffer() (renderedBuffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.renderedBuffer")
	return renderedBuffer
}
