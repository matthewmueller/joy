package speechsynthesisevent

import (
	"github.com/matthewmueller/golly/dom2/speechsynthesiseventinit"
	"github.com/matthewmueller/golly/dom2/window"
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

// CharIndex
func (*SpeechSynthesisEvent) CharIndex() (charIndex uint) {
	js.Rewrite("$<.CharIndex")
	return charIndex
}

// ElapsedTime
func (*SpeechSynthesisEvent) ElapsedTime() (elapsedTime float32) {
	js.Rewrite("$<.ElapsedTime")
	return elapsedTime
}

// Name
func (*SpeechSynthesisEvent) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Utterance
func (*SpeechSynthesisEvent) Utterance() (utterance *window.SpeechSynthesisUtterance) {
	js.Rewrite("$<.Utterance")
	return utterance
}
