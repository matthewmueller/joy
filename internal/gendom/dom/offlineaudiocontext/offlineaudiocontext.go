package offlineaudiocontext

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audiobuffer"
	"github.com/matthewmueller/golly/internal/gendom/dom/audiocontext"
	"github.com/matthewmueller/golly/internal/gendom/dom/offlineaudiocompletionevent"
	"github.com/matthewmueller/golly/js"
)

type OfflineAudioContext struct {
	*audiocontext.AudioContext
}

func (*OfflineAudioContext) StartRendering() (a *audiobuffer.AudioBuffer) {
	js.Rewrite("await $<.startRendering()")
	return a
}

func (*OfflineAudioContext) Suspend(suspendTime float32) {
	js.Rewrite("await $<.suspend($1)", suspendTime)
}

func (*OfflineAudioContext) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*OfflineAudioContext) GetOncomplete() (complete *offlineaudiocompletionevent.OfflineAudioCompletionEvent) {
	js.Rewrite("$<.oncomplete")
	return complete
}

func (*OfflineAudioContext) SetOncomplete(complete *offlineaudiocompletionevent.OfflineAudioCompletionEvent) {
	js.Rewrite("$<.oncomplete = $1", complete)
}
