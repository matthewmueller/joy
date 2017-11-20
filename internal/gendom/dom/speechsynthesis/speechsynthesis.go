package speechsynthesis

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/speechsynthesisutterance"
	"github.com/matthewmueller/golly/internal/gendom/dom/speechsynthesisvoice"
	"github.com/matthewmueller/golly/js"
)

type SpeechSynthesis struct {
	*eventtarget.EventTarget
}

func (*SpeechSynthesis) Cancel() {
	js.Rewrite("$<.cancel()")
}

func (*SpeechSynthesis) GetVoices() (s []*speechsynthesisvoice.SpeechSynthesisVoice) {
	js.Rewrite("$<.getVoices()")
	return s
}

func (*SpeechSynthesis) Pause() {
	js.Rewrite("$<.pause()")
}

func (*SpeechSynthesis) Resume() {
	js.Rewrite("$<.resume()")
}

func (*SpeechSynthesis) Speak(utterance *speechsynthesisutterance.SpeechSynthesisUtterance) {
	js.Rewrite("$<.speak($1)", utterance)
}

func (*SpeechSynthesis) GetOnvoiceschanged() (voiceschanged *event.Event) {
	js.Rewrite("$<.onvoiceschanged")
	return voiceschanged
}

func (*SpeechSynthesis) SetOnvoiceschanged(voiceschanged *event.Event) {
	js.Rewrite("$<.onvoiceschanged = $1", voiceschanged)
}

func (*SpeechSynthesis) GetPaused() (paused bool) {
	js.Rewrite("$<.paused")
	return paused
}

func (*SpeechSynthesis) GetPending() (pending bool) {
	js.Rewrite("$<.pending")
	return pending
}

func (*SpeechSynthesis) GetSpeaking() (speaking bool) {
	js.Rewrite("$<.speaking")
	return speaking
}
