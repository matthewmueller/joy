package offlineaudiocompletionevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audiobuffer"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type OfflineAudioCompletionEvent struct {
	*event.Event
}

func (*OfflineAudioCompletionEvent) GetRenderedBuffer() (renderedBuffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.renderedBuffer")
	return renderedBuffer
}
