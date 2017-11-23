package offlineaudiocontext

import (
	"github.com/matthewmueller/golly/dom2/audiobuffer"
	"github.com/matthewmueller/golly/dom2/audionode"
	"github.com/matthewmueller/golly/dom2/offlineaudiocompletionevent"
	"github.com/matthewmueller/golly/js"
)

// js:"OfflineAudioContext,omit"
type OfflineAudioContext struct {
	audionode.AudioContext
}

// StartRendering
func (*OfflineAudioContext) StartRendering() (a *audiobuffer.AudioBuffer) {
	js.Rewrite("await $<.StartRendering()")
	return a
}

// Suspend
func (*OfflineAudioContext) Suspend(suspendTime float32) {
	js.Rewrite("await $<.Suspend($1)", suspendTime)
}

// Length
func (*OfflineAudioContext) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}

// Oncomplete
func (*OfflineAudioContext) Oncomplete() (oncomplete func(*offlineaudiocompletionevent.OfflineAudioCompletionEvent)) {
	js.Rewrite("$<.Oncomplete")
	return oncomplete
}

// Oncomplete
func (*OfflineAudioContext) SetOncomplete(oncomplete func(*offlineaudiocompletionevent.OfflineAudioCompletionEvent)) {
	js.Rewrite("$<.Oncomplete = $1", oncomplete)
}
