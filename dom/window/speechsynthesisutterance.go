package window

import (
	"github.com/matthewmueller/golly/dom2/speechsynthesisvoice"
	"github.com/matthewmueller/golly/js"
)

// NewSpeechSynthesisUtterance fn
func NewSpeechSynthesisUtterance(text *string) *SpeechSynthesisUtterance {
	js.Rewrite("SpeechSynthesisUtterance")
	return &SpeechSynthesisUtterance{}
}

// SpeechSynthesisUtterance struct
// js:"SpeechSynthesisUtterance,omit"
type SpeechSynthesisUtterance struct {
	EventTarget
}

// Lang prop
func (*SpeechSynthesisUtterance) Lang() (lang string) {
	js.Rewrite("$<.lang")
	return lang
}

// Lang prop
func (*SpeechSynthesisUtterance) SetLang(lang string) {
	js.Rewrite("$<.lang = $1", lang)
}

// Onboundary prop
func (*SpeechSynthesisUtterance) Onboundary() (onboundary func(Event)) {
	js.Rewrite("$<.onboundary")
	return onboundary
}

// Onboundary prop
func (*SpeechSynthesisUtterance) SetOnboundary(onboundary func(Event)) {
	js.Rewrite("$<.onboundary = $1", onboundary)
}

// Onend prop
func (*SpeechSynthesisUtterance) Onend() (onend func(Event)) {
	js.Rewrite("$<.onend")
	return onend
}

// Onend prop
func (*SpeechSynthesisUtterance) SetOnend(onend func(Event)) {
	js.Rewrite("$<.onend = $1", onend)
}

// Onerror prop
func (*SpeechSynthesisUtterance) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*SpeechSynthesisUtterance) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onmark prop
func (*SpeechSynthesisUtterance) Onmark() (onmark func(Event)) {
	js.Rewrite("$<.onmark")
	return onmark
}

// Onmark prop
func (*SpeechSynthesisUtterance) SetOnmark(onmark func(Event)) {
	js.Rewrite("$<.onmark = $1", onmark)
}

// Onpause prop
func (*SpeechSynthesisUtterance) Onpause() (onpause func(Event)) {
	js.Rewrite("$<.onpause")
	return onpause
}

// Onpause prop
func (*SpeechSynthesisUtterance) SetOnpause(onpause func(Event)) {
	js.Rewrite("$<.onpause = $1", onpause)
}

// Onresume prop
func (*SpeechSynthesisUtterance) Onresume() (onresume func(Event)) {
	js.Rewrite("$<.onresume")
	return onresume
}

// Onresume prop
func (*SpeechSynthesisUtterance) SetOnresume(onresume func(Event)) {
	js.Rewrite("$<.onresume = $1", onresume)
}

// Onstart prop
func (*SpeechSynthesisUtterance) Onstart() (onstart func(Event)) {
	js.Rewrite("$<.onstart")
	return onstart
}

// Onstart prop
func (*SpeechSynthesisUtterance) SetOnstart(onstart func(Event)) {
	js.Rewrite("$<.onstart = $1", onstart)
}

// Pitch prop
func (*SpeechSynthesisUtterance) Pitch() (pitch float32) {
	js.Rewrite("$<.pitch")
	return pitch
}

// Pitch prop
func (*SpeechSynthesisUtterance) SetPitch(pitch float32) {
	js.Rewrite("$<.pitch = $1", pitch)
}

// Rate prop
func (*SpeechSynthesisUtterance) Rate() (rate float32) {
	js.Rewrite("$<.rate")
	return rate
}

// Rate prop
func (*SpeechSynthesisUtterance) SetRate(rate float32) {
	js.Rewrite("$<.rate = $1", rate)
}

// Text prop
func (*SpeechSynthesisUtterance) Text() (text string) {
	js.Rewrite("$<.text")
	return text
}

// Text prop
func (*SpeechSynthesisUtterance) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

// Voice prop
func (*SpeechSynthesisUtterance) Voice() (voice *speechsynthesisvoice.SpeechSynthesisVoice) {
	js.Rewrite("$<.voice")
	return voice
}

// Voice prop
func (*SpeechSynthesisUtterance) SetVoice(voice *speechsynthesisvoice.SpeechSynthesisVoice) {
	js.Rewrite("$<.voice = $1", voice)
}

// Volume prop
func (*SpeechSynthesisUtterance) Volume() (volume float32) {
	js.Rewrite("$<.volume")
	return volume
}

// Volume prop
func (*SpeechSynthesisUtterance) SetVolume(volume float32) {
	js.Rewrite("$<.volume = $1", volume)
}
