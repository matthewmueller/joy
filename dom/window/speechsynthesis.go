package window

import (
	"github.com/matthewmueller/joy/dom/speechsynthesisvoice"
	"github.com/matthewmueller/joy/macro"
)

var _ EventTarget = (*SpeechSynthesis)(nil)

// SpeechSynthesis struct
// js:"SpeechSynthesis,omit"
type SpeechSynthesis struct {
}

// Cancel fn
// js:"cancel"
func (*SpeechSynthesis) Cancel() {
	macro.Rewrite("$_.cancel()")
}

// GetVoices fn
// js:"getVoices"
func (*SpeechSynthesis) GetVoices() (s []*speechsynthesisvoice.SpeechSynthesisVoice) {
	macro.Rewrite("$_.getVoices()")
	return s
}

// Pause fn
// js:"pause"
func (*SpeechSynthesis) Pause() {
	macro.Rewrite("$_.pause()")
}

// Resume fn
// js:"resume"
func (*SpeechSynthesis) Resume() {
	macro.Rewrite("$_.resume()")
}

// Speak fn
// js:"speak"
func (*SpeechSynthesis) Speak(utterance *SpeechSynthesisUtterance) {
	macro.Rewrite("$_.speak($1)", utterance)
}

// AddEventListener fn
// js:"addEventListener"
func (*SpeechSynthesis) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SpeechSynthesis) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SpeechSynthesis) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Onvoiceschanged prop
// js:"onvoiceschanged"
func (*SpeechSynthesis) Onvoiceschanged() (onvoiceschanged func(Event)) {
	macro.Rewrite("$_.onvoiceschanged")
	return onvoiceschanged
}

// SetOnvoiceschanged prop
// js:"onvoiceschanged"
func (*SpeechSynthesis) SetOnvoiceschanged(onvoiceschanged func(Event)) {
	macro.Rewrite("$_.onvoiceschanged = $1", onvoiceschanged)
}

// Paused prop
// js:"paused"
func (*SpeechSynthesis) Paused() (paused bool) {
	macro.Rewrite("$_.paused")
	return paused
}

// Pending prop
// js:"pending"
func (*SpeechSynthesis) Pending() (pending bool) {
	macro.Rewrite("$_.pending")
	return pending
}

// Speaking prop
// js:"speaking"
func (*SpeechSynthesis) Speaking() (speaking bool) {
	macro.Rewrite("$_.speaking")
	return speaking
}
