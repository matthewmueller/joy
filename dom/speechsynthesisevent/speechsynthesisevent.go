package speechsynthesisevent

import (
	"github.com/matthewmueller/golly/dom/speechsynthesiseventinit"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(kind string, eventinitdict *speechsynthesiseventinit.SpeechSynthesisEventInit) *SpeechSynthesisEvent {
	js.Rewrite("SpeechSynthesisEvent")
	return &SpeechSynthesisEvent{}
}

// SpeechSynthesisEvent struct
// js:"SpeechSynthesisEvent,omit"
type SpeechSynthesisEvent struct {
	window.Event
}

// CharIndex prop
func (*SpeechSynthesisEvent) CharIndex() (charIndex uint) {
	js.Rewrite("$<.charIndex")
	return charIndex
}

// ElapsedTime prop
func (*SpeechSynthesisEvent) ElapsedTime() (elapsedTime float32) {
	js.Rewrite("$<.elapsedTime")
	return elapsedTime
}

// Name prop
func (*SpeechSynthesisEvent) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Utterance prop
func (*SpeechSynthesisEvent) Utterance() (utterance *window.SpeechSynthesisUtterance) {
	js.Rewrite("$<.utterance")
	return utterance
}
