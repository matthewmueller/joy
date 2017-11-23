package offlineaudiocontext

import (
	"github.com/matthewmueller/golly/dom/audiobuffer"
	"github.com/matthewmueller/golly/dom/audionode"
	"github.com/matthewmueller/golly/dom/offlineaudiocompletionevent"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(numberofchannels uint, length uint, samplerate float32) *OfflineAudioContext {
	js.Rewrite("OfflineAudioContext")
	return &OfflineAudioContext{}
}

// OfflineAudioContext struct
// js:"OfflineAudioContext,omit"
type OfflineAudioContext struct {
	audionode.AudioContext
}

// StartRendering fn
func (*OfflineAudioContext) StartRendering() (a *audiobuffer.AudioBuffer) {
	js.Rewrite("await $<.startRendering()")
	return a
}

// Suspend fn
func (*OfflineAudioContext) Suspend(suspendTime float32) {
	js.Rewrite("await $<.suspend($1)", suspendTime)
}

// Length prop
func (*OfflineAudioContext) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}

// Oncomplete prop
func (*OfflineAudioContext) Oncomplete() (oncomplete func(*offlineaudiocompletionevent.OfflineAudioCompletionEvent)) {
	js.Rewrite("$<.oncomplete")
	return oncomplete
}

// Oncomplete prop
func (*OfflineAudioContext) SetOncomplete(oncomplete func(*offlineaudiocompletionevent.OfflineAudioCompletionEvent)) {
	js.Rewrite("$<.oncomplete = $1", oncomplete)
}
