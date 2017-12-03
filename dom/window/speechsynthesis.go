package window

import (
	"github.com/matthewmueller/joy/dom/speechsynthesisvoice"
	"github.com/matthewmueller/joy/js"
)

var _ EventTarget = (*SpeechSynthesis)(nil)

// SpeechSynthesis struct
// js:"SpeechSynthesis,omit"
type SpeechSynthesis struct {
}

// Cancel fn
// js:"cancel"
func (*SpeechSynthesis) Cancel() {
	js.Rewrite("$_.cancel()")
}

// GetVoices fn
// js:"getVoices"
func (*SpeechSynthesis) GetVoices() (s []*speechsynthesisvoice.SpeechSynthesisVoice) {
	js.Rewrite("$_.getVoices()")
	return s
}

// Pause fn
// js:"pause"
func (*SpeechSynthesis) Pause() {
	js.Rewrite("$_.pause()")
}

// Resume fn
// js:"resume"
func (*SpeechSynthesis) Resume() {
	js.Rewrite("$_.resume()")
}

// Speak fn
// js:"speak"
func (*SpeechSynthesis) Speak(utterance *SpeechSynthesisUtterance) {
	js.Rewrite("$_.speak($1)", utterance)
}

// AddEventListener fn
// js:"addEventListener"
func (*SpeechSynthesis) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SpeechSynthesis) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SpeechSynthesis) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Onvoiceschanged prop
// js:"onvoiceschanged"
func (*SpeechSynthesis) Onvoiceschanged() (onvoiceschanged func(Event)) {
	js.Rewrite("$_.onvoiceschanged")
	return onvoiceschanged
}

// SetOnvoiceschanged prop
// js:"onvoiceschanged"
func (*SpeechSynthesis) SetOnvoiceschanged(onvoiceschanged func(Event)) {
	js.Rewrite("$_.onvoiceschanged = $1", onvoiceschanged)
}

// Paused prop
// js:"paused"
func (*SpeechSynthesis) Paused() (paused bool) {
	js.Rewrite("$_.paused")
	return paused
}

// Pending prop
// js:"pending"
func (*SpeechSynthesis) Pending() (pending bool) {
	js.Rewrite("$_.pending")
	return pending
}

// Speaking prop
// js:"speaking"
func (*SpeechSynthesis) Speaking() (speaking bool) {
	js.Rewrite("$_.speaking")
	return speaking
}
