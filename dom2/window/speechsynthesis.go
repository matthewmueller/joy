package window

import (
	"github.com/matthewmueller/golly/dom2/speechsynthesisvoice"
	"github.com/matthewmueller/golly/js"
)

// js:"SpeechSynthesis,omit"
type SpeechSynthesis struct {
	EventTarget
}

// Cancel
func (*SpeechSynthesis) Cancel() {
	js.Rewrite("$<.Cancel()")
}

// GetVoices
func (*SpeechSynthesis) GetVoices() (s []*speechsynthesisvoice.SpeechSynthesisVoice) {
	js.Rewrite("$<.GetVoices()")
	return s
}

// Pause
func (*SpeechSynthesis) Pause() {
	js.Rewrite("$<.Pause()")
}

// Resume
func (*SpeechSynthesis) Resume() {
	js.Rewrite("$<.Resume()")
}

// Speak
func (*SpeechSynthesis) Speak(utterance *SpeechSynthesisUtterance) {
	js.Rewrite("$<.Speak($1)", utterance)
}

// Onvoiceschanged
func (*SpeechSynthesis) Onvoiceschanged() (onvoiceschanged func(Event)) {
	js.Rewrite("$<.Onvoiceschanged")
	return onvoiceschanged
}

// Onvoiceschanged
func (*SpeechSynthesis) SetOnvoiceschanged(onvoiceschanged func(Event)) {
	js.Rewrite("$<.Onvoiceschanged = $1", onvoiceschanged)
}

// Paused
func (*SpeechSynthesis) Paused() (paused bool) {
	js.Rewrite("$<.Paused")
	return paused
}

// Pending
func (*SpeechSynthesis) Pending() (pending bool) {
	js.Rewrite("$<.Pending")
	return pending
}

// Speaking
func (*SpeechSynthesis) Speaking() (speaking bool) {
	js.Rewrite("$<.Speaking")
	return speaking
}
