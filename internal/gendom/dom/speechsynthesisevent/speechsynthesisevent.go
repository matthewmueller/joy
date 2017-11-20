package speechsynthesisevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/speechsynthesisutterance"
	"github.com/matthewmueller/golly/js"
)

type SpeechSynthesisEvent struct {
	*event.Event
}

func (*SpeechSynthesisEvent) GetCharIndex() (charIndex uint) {
	js.Rewrite("$<.charIndex")
	return charIndex
}

func (*SpeechSynthesisEvent) GetElapsedTime() (elapsedTime float32) {
	js.Rewrite("$<.elapsedTime")
	return elapsedTime
}

func (*SpeechSynthesisEvent) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*SpeechSynthesisEvent) GetUtterance() (utterance *speechsynthesisutterance.SpeechSynthesisUtterance) {
	js.Rewrite("$<.utterance")
	return utterance
}
