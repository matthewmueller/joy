package window

import (
	"github.com/matthewmueller/golly/dom2/speechsynthesisvoice"
	"github.com/matthewmueller/golly/js"
)

// SpeechSynthesis struct
// js:"SpeechSynthesis,omit"
type SpeechSynthesis struct {
	EventTarget
}

// Cancel fn
func (*SpeechSynthesis) Cancel() {
	js.Rewrite("$<.cancel()")
}

// GetVoices fn
func (*SpeechSynthesis) GetVoices() (s []*speechsynthesisvoice.SpeechSynthesisVoice) {
	js.Rewrite("$<.getVoices()")
	return s
}

// Pause fn
func (*SpeechSynthesis) Pause() {
	js.Rewrite("$<.pause()")
}

// Resume fn
func (*SpeechSynthesis) Resume() {
	js.Rewrite("$<.resume()")
}

// Speak fn
func (*SpeechSynthesis) Speak(utterance *SpeechSynthesisUtterance) {
	js.Rewrite("$<.speak($1)", utterance)
}

// Onvoiceschanged prop
func (*SpeechSynthesis) Onvoiceschanged() (onvoiceschanged func(Event)) {
	js.Rewrite("$<.onvoiceschanged")
	return onvoiceschanged
}

// Onvoiceschanged prop
func (*SpeechSynthesis) SetOnvoiceschanged(onvoiceschanged func(Event)) {
	js.Rewrite("$<.onvoiceschanged = $1", onvoiceschanged)
}

// Paused prop
func (*SpeechSynthesis) Paused() (paused bool) {
	js.Rewrite("$<.paused")
	return paused
}

// Pending prop
func (*SpeechSynthesis) Pending() (pending bool) {
	js.Rewrite("$<.pending")
	return pending
}

// Speaking prop
func (*SpeechSynthesis) Speaking() (speaking bool) {
	js.Rewrite("$<.speaking")
	return speaking
}
